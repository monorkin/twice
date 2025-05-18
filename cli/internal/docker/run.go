package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/monorkin/twice/cli/internal/config"
)

func RunProduct(product *config.ProductConfig) error {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	config := &container.Config{
		Image: product.Image(),
		Env: []string{
			"APP_DOMAIN=" + product.Domain,
			"TLS_EMAIL=" + product.EmailAddress,
			fmt.Sprintf("TLS_ENABLED=%t", product.HTTPS),
		},
	}

	hostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped",
		},
	}

	resp, err := dockerClient.ContainerCreate(
		ctx,
		config,
		hostConfig,
		nil,
		nil,
		product.ContainerName(),
	)
	if err != nil {
		return fmt.Errorf("failed to create app container: %w", err)
	}

	if err := dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start app: %w", err)
	}

	return nil
}
