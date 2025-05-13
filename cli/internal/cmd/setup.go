package cmd

import (
	"fmt"
	"os"

	"github.com/monorkin/twice/cli/internal/docker"
	"github.com/spf13/cobra"
)

const (
	CheckMarkIcon = "✅"
	CrossIcon     = "❌"
)

func NewSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup <license-key>",
		Short: "Setup the environment for a product and install it",
		Long:  `Install all the necessary dependencies to run a product and then installs the product associated with the given license key`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			licenseKey := args[0]
			runSetupCmd(licenseKey)
		},
	}

	return cmd
}

func runSetupCmd(licenseKey string) {
	registry := "localhost:5000"

	// Step 1 - Check if docker is installed
	if docker.IsInstalled() {
		println(CheckMarkIcon + " Docker is installed")
	} else {
		println(CrossIcon + " Docker is not installed")

		if err := docker.InstallDocker(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// Step 2 - Check if docker is running
	running, err := docker.IsRunning()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if running {
		println(CheckMarkIcon + " Docker is running")
	} else {
		println(CrossIcon + " Docker is not running")

		if err := docker.Start(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			println(CheckMarkIcon + " Docker is now running")
		}
	}

	// Step 3 - Check the license
	println("Checking license...")

	// Step 4 - Login with the registry
	println("Downloading app...")
	identityToken, err := docker.LoginWithRegistry(registry, "username", licenseKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = docker.PullImageWithIdentityToken("twice", registry, identityToken, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	println(CheckMarkIcon + " App downloaded")
}
