package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/monorkin/twice/cli/internal/config"
)

func IsProductRunning(product *config.ProductConfig) (bool, error) {
	containerName := product.ContainerName()
	return IsContainerRunning(containerName)
}

func IsContainerRunning(containerName string) (bool, error) {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return false, fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	dockerContainer, err := dockerClient.ContainerInspect(ctx, containerName)
	if err != nil {
		return false, fmt.Errorf("failed to inspect container: %v", err)
	}

	return dockerContainer.State.Running, nil
}
