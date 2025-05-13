package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	imageTypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullImageWithIdentityToken(image string, registry string, identityToken string, printProgress bool) error {
	// The image has to include the name of the registry as a prefix
	if !strings.HasPrefix(image, registry) {
		image = fmt.Sprintf("%s/%s", registry, image)
	}

	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	options := imageTypes.PullOptions{
		RegistryAuth: identityToken,
	}

	responseBody, err := dockerClient.ImagePull(ctx, image, options)
	if err != nil {
		return fmt.Errorf("failed to pull image: %v", err)
	}
	defer responseBody.Close()

	if printProgress {
		_, err = io.Copy(os.Stdout, responseBody)
		if err != nil {
			return fmt.Errorf("error reading pull response: %v", err)
		}
	}

	return nil
}
