package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/artarts36/service-navigator/internal/shared"
	"io"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Monitor struct {
	docker             *client.Client
	filler             Filler
	networkName        string
	currentContainerID string
}

func NewMonitor(docker *client.Client, urlResolver Filler, networkName string, currentContainerID string) *Monitor {
	return &Monitor{docker: docker, filler: urlResolver, networkName: networkName, currentContainerID: currentContainerID}
}

type Opts struct {
	Concurrent int
}

func (m *Monitor) Show(ctx context.Context, opts Opts) ([]*domain.ServiceStatus, error) {
	log.Printf("[Service][Monitor] Fetching containers")

	containers, err := m.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "network",
			Value: m.networkName,
		}),
	})

	log.Printf("[Service][Monitor] Fetched %d containers", len(containers))

	if err != nil {
		return []*domain.ServiceStatus{}, err
	}

	return m.collectServices(ctx, containers, opts)
}

func (m *Monitor) KillContainer(ctx context.Context, containerID string) error {
	return m.docker.ContainerKill(ctx, containerID, "")
}

func (m *Monitor) collectServices(
	ctx context.Context,
	containers []types.Container,
	opts Opts,
) ([]*domain.ServiceStatus, error) {
	services := make([]*domain.ServiceStatus, 0, len(containers))

	chunks := make([][]types.Container, 0)
	if opts.Concurrent <= 0 {
		chunks = append(chunks, containers)
	} else {
		chunks = shared.ChunkSlice(containers, opts.Concurrent)
	}

	for _, chunk := range chunks {
		wg := sync.WaitGroup{}

		for _, container := range chunk {
			wg.Add(1)

			container := container
			go func() {
				defer wg.Done()

				service, err := m.collectServiceStatus(ctx, container)

				if err == nil {
					services = append(services, service)
				} else {
					log.Printf("Failed to collect service: %s", err)
				}
			}()
		}

		wg.Wait()
	}

	return services, nil
}

func (m *Monitor) collectServiceStatus(ctx context.Context, container types.Container) (*domain.ServiceStatus, error) {
	cont, inspectErr := m.docker.ContainerInspect(ctx, container.ID)

	if inspectErr != nil {
		return nil, inspectErr
	}

	stat, _ := m.collectStats(ctx, container.ID)

	status := &domain.ServiceStatus{
		Name:        cont.Name,
		Status:      container.Status,
		ContainerID: cont.ID,
		Self:        strings.HasPrefix(cont.ID, m.currentContainerID),
	}

	environment := map[string]string{}
	for _, envVar := range cont.Config.Env {
		varBag := strings.Split(envVar, "=")

		varName := varBag[0]
		varValue := varBag[1]

		environment[varName] = varValue
	}

	m.filler.Fill(status, &datastruct.Container{
		Short:       container,
		Full:        cont,
		Stats:       &stat,
		Environment: environment,
	})

	return status, nil
}

func (m *Monitor) collectStats(ctx context.Context, containerID string) (datastruct.Stats, error) {
	response, err := m.docker.ContainerStatsOneShot(ctx, containerID)
	if err != nil {
		return datastruct.Stats{}, fmt.Errorf("failed to get response: %v", err)
	}

	responseContent, err := io.ReadAll(response.Body)
	if err != nil {
		return datastruct.Stats{}, fmt.Errorf("failed to read response: %v", err)
	}

	var stat datastruct.Stats

	err = json.Unmarshal(responseContent, &stat)

	if err != nil {
		return datastruct.Stats{}, fmt.Errorf("failed to parse response: %v", err)
	}

	return stat, nil
}
