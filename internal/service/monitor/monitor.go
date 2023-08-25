package monitor

import (
	"context"
	"log"
	"strings"
	"sync"

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

func (m *Monitor) Show(ctx context.Context) ([]*entity.Service, error) {
	_, err := client.NewClientWithOpts()

	if err != nil {
		return []*entity.Service{}, errors.Errorf("Failed to create docker client: %s", err)
	}

	containers, err := m.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "network",
			Value: "infra",
		}),
	})

	if err != nil {
		return []*entity.Service{}, err
	}

	return m.collectServices(ctx, containers)
}

func (m *Monitor) KillContainer(ctx context.Context, containerID string) error {
	return m.docker.ContainerKill(ctx, containerID, "")
}

func (m *Monitor) collectServices(ctx context.Context, containers []types.Container) ([]*entity.Service, error) {
	services := make([]*entity.Service, 0, len(containers))

	wg := sync.WaitGroup{}

	for _, container := range containers {
		wg.Add(1)

		container := container
		go func() {
			service, err := m.collectService(ctx, container)

			if err == nil {
				services = append(services, service)
			} else {
				log.Printf("Failed to collect service: %s", err)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return services, nil
}

func (m *Monitor) collectService(ctx context.Context, container types.Container) (*entity.Service, error) {
	cont, inspectErr := m.docker.ContainerInspect(ctx, container.ID)

	if inspectErr != nil {
		return nil, inspectErr
	}

	service := &entity.Service{
		Name:        cont.Name,
		Status:      cont.State.Status,
		ContainerID: cont.ID,
		Self:        strings.HasPrefix(cont.ID, m.currentContainerID),
	}

	m.filler.Fill(service, &entity.Container{
		Short: container,
		Full:  cont,
	})

	return service, nil
}
