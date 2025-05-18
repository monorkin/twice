package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	AppName  = "twice"
	FileName = "config.json"
)

func ConfigFilePath() (string, error) {
	configDir, err := ConfigDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, FileName), nil
}

func ConfigDirPath() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "linux":
		configDir = os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			usr, err := user.Current()
			if err != nil {
				return os.Getwd()
			}
			configDir = filepath.Join(usr.HomeDir, ".config")
		}

	case "darwin":
		usr, err := user.Current()
		if err != nil {
			return os.Getwd()
		}
		configDir = filepath.Join(usr.HomeDir, "Library", "Application Support")

	case "windows":
		configDir = os.Getenv("AppData")
		if configDir == "" {
			return os.Getwd()
		}

	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return filepath.Join(configDir, AppName), nil
}
