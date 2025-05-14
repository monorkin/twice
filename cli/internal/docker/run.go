package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunApp(image string, registry string, emailAddress string, domain string, tls_enabled bool) error {
	image = SanitizeImageName(image, registry)
	containerName := ContainerNameFromImage(image)

	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	config := &container.Config{
		Image: image,
		Env: []string{
			"APP_DOMAIN=" + domain,
			"TLS_EMAIL=" + emailAddress,
			fmt.Sprintf("TLS_ENABLED=%t", tls_enabled),
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
		containerName,
	)
	if err != nil {
		return fmt.Errorf("failed to create app container: %w", err)
	}

	if err := dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start app: %w", err)
	}

	return nil
}
