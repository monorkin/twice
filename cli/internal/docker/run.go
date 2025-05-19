package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/monorkin/twice/cli/internal/config"
)

func RunProduct(product *config.Product) error {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	config := &container.Config{
		Image: product.Image(),
		Env: []string{
			"SSL_EMAIL=" + product.EmailAddress,
			"SSL_DOMAIN=" + product.Domain,
			fmt.Sprintf("DISABLE_SSL=%t", !product.HTTPS),
			"VAPID_PRIVATE_KEY=" + product.VAPIDPrivateKey,
			"VAPID_PUBLIC_KEY=" + product.VAPIDPublicKey,
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
