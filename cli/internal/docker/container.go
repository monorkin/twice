package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/monorkin/twice/cli/internal/config"
)

func IsProductRunning(product *config.Product) (bool, error) {
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

func StartProductContainer(product *config.Product) error {
	containerName := product.ContainerName()
	return StartContainer(containerName)
}

func StartContainer(containerName string) error {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	err = dockerClient.ContainerStart(ctx, containerName, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}

	return nil
}

func StopProductContainer(product *config.Product) error {
	containerName := product.ContainerName()
	return StopContainer(containerName)
}

func StopContainer(containerName string) error {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	err = dockerClient.ContainerStop(ctx, containerName, container.StopOptions{})
	if err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}

	return nil
}
