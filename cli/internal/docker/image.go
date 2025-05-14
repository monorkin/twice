package docker

import (
	"fmt"
	"strings"
)

func SanitizeImageName(image string, registry string) string {
	if !strings.HasPrefix(image, registry) {
		image = fmt.Sprintf("%s/%s", registry, image)
	}

	if !strings.Contains(image, ":") {
		image = fmt.Sprintf("%s:latest", image)
	}

	return image
}

func ContainerNameFromImage(image string) string {
	parts := strings.Split(image, "/")
	taggedName := parts[len(parts)-1]
	nameParts := strings.Split(taggedName, ":")
	name := nameParts[0]

	return name
}
