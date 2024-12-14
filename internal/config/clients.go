package config

import (
	"fmt"

	dockerClient "github.com/docker/docker/client"

	githubClient "github.com/google/go-github/v67/github"
)

func (c *Container) initClients(cfg *Config) error {
	dockClient, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create docker client: %w", err)
	}

	c.Clients.Docker = dockClient

	ghClient := githubClient.NewClient(nil)
	if cfg.Credentials.GithubToken != "" {
		ghClient = ghClient.WithAuthToken(cfg.Credentials.GithubToken)
	}
	c.Clients.Github = ghClient

	return nil
}
