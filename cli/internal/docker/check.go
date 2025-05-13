package docker

import (
	"fmt"
	"os/exec"
	"strings"
)

func FindDockerBin() (string, error) {
	return exec.LookPath("docker")
}

func IsInstalled() bool {
	_, err := FindDockerBin()

	return err == nil
}

func IsRunning() (bool, error) {
	binPath, err := FindDockerBin()
	if err != nil {
		return false, fmt.Errorf("docker is not installed: %w", err)
	}

	out, err := exec.Command(binPath, "info").Output()
	if err != nil {
		return false, fmt.Errorf("docker is not running: %w", err)
	}

	if strings.Contains(string(out), "Server Version") {
		return true, nil
	}

	return false, nil
}
