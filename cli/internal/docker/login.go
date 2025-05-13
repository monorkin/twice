package docker

import (
	"context"
	"fmt"

	registryTypes "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

func LoginWithRegistry(registry string, username string, password string) (string, error) {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	authConfig := registryTypes.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: registry,
	}

	resp, err := dockerClient.RegistryLogin(ctx, authConfig)
	if err != nil {
		return "", fmt.Errorf("failed to login to registry: %v", err)
	}

	return resp.IdentityToken, nil
}
