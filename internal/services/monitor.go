package services

import (
	"context"
	"fmt"
	"github.com/artarts36/service-navigator/internal/services/weburl"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Monitor struct {
	docker      *client.Client
	urlResolver weburl.UrlResolver
}

func NewMonitor(docker *client.Client, urlResolver weburl.UrlResolver) *Monitor {
	return &Monitor{docker: docker, urlResolver: urlResolver}
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

		statuses = append(statuses, &Service{
			Name:   srv.Names[0],
			WebUrl: m.urlResolver.Resolve(&cont),
			Status: cont.State.Status,
		})
	}

	return statuses, nil
}
