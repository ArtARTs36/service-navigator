package monitor

import (
	"context"
	"fmt"
	"github.com/artarts36/service-navigator/internal/service/entity"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Monitor struct {
	docker *client.Client
	filler Filler
}

func NewMonitor(docker *client.Client, urlResolver Filler) *Monitor {
	return &Monitor{docker: docker, filler: urlResolver}
}

func (m *Monitor) Show(ctx context.Context) ([]*entity.Service, error) {
	_, err := client.NewClientWithOpts()

	if err != nil {
		return []*entity.Service{}, errors.Errorf("Failed to create docker client: %s", err)
	}

	network, err := m.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "network",
			Value: "infra",
		}),
	})

	if err != nil {
		return []*entity.Service{}, err
	}

	statuses := make([]*entity.Service, 0, len(network))

	for _, srv := range network {
		cont, err := m.docker.ContainerInspect(ctx, srv.ID)

		if err != nil {
			fmt.Println(err)

			continue
		}

		service := &entity.Service{
			Name:   srv.Names[0],
			Status: cont.State.Status,
		}

		m.filler.Fill(service, &entity.Container{
			Short: &srv,
			Full:  &cont,
		})

		statuses = append(statuses, service)
	}

	return statuses, nil
}
