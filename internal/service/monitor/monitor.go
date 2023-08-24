package monitor

import (
	"context"
	"log"

	"github.com/artarts36/service-navigator/internal/service/entity"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Monitor struct {
	docker      *client.Client
	filler      Filler
	networkName string
}

func NewMonitor(docker *client.Client, urlResolver Filler, networkName string) *Monitor {
	return &Monitor{docker: docker, filler: urlResolver, networkName: networkName}
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

	services := make([]*entity.Service, 0, len(containers))

	for _, container := range containers {
		cont, inspectErr := m.docker.ContainerInspect(ctx, container.ID)

		if inspectErr != nil {
			log.Println(inspectErr)

			continue
		}

		service := &entity.Service{
			Name:   cont.Name,
			Status: cont.State.Status,
		}

		m.filler.Fill(service, &entity.Container{
			Short: container,
			Full:  cont,
		})

		services = append(services, service)
	}

	return services, nil
}
