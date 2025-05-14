package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	imageTypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullImageWithIdentityToken(image string, registry string, username string, password string, printProgress bool) error {
	image = SanitizeImageName(image, registry)

	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	authJSON := map[string]string{
		"username":      username,
		"password":      password,
		"serveraddress": registry,
	}
	authBytes, err := json.Marshal(authJSON)
	if err != nil {
		return fmt.Errorf("failed to marshal auth: %v", err)
	}
	authBase64 := base64.StdEncoding.EncodeToString(authBytes)

	options := imageTypes.PullOptions{
		RegistryAuth: authBase64,
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
