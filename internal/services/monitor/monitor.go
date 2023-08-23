package monitor

import (
	"context"
	"fmt"
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

func (m *Monitor) Show(ctx context.Context) ([]*Service, error) {
	_, err := client.NewClientWithOpts()

	if err != nil {
		return []*Service{}, errors.Errorf("Failed to create docker client: %s", err)
	}

	network, err := m.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "network",
			Value: "infra",
		}),
	})

	if err != nil {
		return []*Service{}, err
	}

	statuses := make([]*Service, 0, len(network))

	for _, srv := range network {
		cont, err := m.docker.ContainerInspect(ctx, srv.ID)

		if err != nil {
			fmt.Println(err)

			continue
		}

		service := &Service{
			Name:   srv.Names[0],
			Status: cont.State.Status,
		}

		m.filler.Fill(service, &cont)

		statuses = append(statuses, service)
	}

	return statuses, nil
}
