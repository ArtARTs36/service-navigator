package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/service/entity"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
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

func (m *Monitor) Show(ctx context.Context) (map[string]*domain.ServiceStatus, error) {
	_, err := client.NewClientWithOpts()

	if err != nil {
		return map[string]*domain.ServiceStatus{}, errors.Errorf("Failed to create docker client: %s", err)
	}

	containers, err := m.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "network",
			Value: m.networkName,
		}),
	})

	if err != nil {
		return map[string]*domain.ServiceStatus{}, err
	}

	return m.collectServices(ctx, containers)
}

func (m *Monitor) KillContainer(ctx context.Context, containerID string) error {
	return m.docker.ContainerKill(ctx, containerID, "")
}

func (m *Monitor) collectServices(
	ctx context.Context,
	containers []types.Container,
) (map[string]*domain.ServiceStatus, error) {
	services := make(map[string]*domain.ServiceStatus, 0)

	wg := sync.WaitGroup{}

	for _, container := range containers {
		wg.Add(1)

		container := container
		go func() {
			service, err := m.collectServiceStatus(ctx, container)

			if err == nil {
				services[service.ContainerID] = service
			} else {
				log.Printf("Failed to collect service: %s", err)
			}

			wg.Done()
		}()
	}

	wg.Wait()

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
		Status:      cont.State.Status,
		ContainerID: cont.ID,
		Self:        strings.HasPrefix(cont.ID, m.currentContainerID),
	}

	m.filler.Fill(status, &entity.Container{
		Short: container,
		Full:  cont,
		Stats: &stat,
	})

	return status, nil
}

func (m *Monitor) collectStats(ctx context.Context, containerID string) (entity.Stats, error) {
	response, err := m.docker.ContainerStatsOneShot(ctx, containerID)
	if err != nil {
		return entity.Stats{}, fmt.Errorf("failed to get response: %v", err)
	}

	responseContent, err := io.ReadAll(response.Body)
	if err != nil {
		return entity.Stats{}, fmt.Errorf("failed to read response: %v", err)
	}

	var stat entity.Stats

	err = json.Unmarshal(responseContent, &stat)

	if err != nil {
		return entity.Stats{}, fmt.Errorf("failed to parse response: %v", err)
	}

	return stat, nil
}
