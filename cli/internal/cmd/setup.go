package cmd

import (
	"fmt"
	"os"

	"github.com/monorkin/twice/cli/internal/api"
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
	apiClient := api.NewClient("localhost:3000")
	license, err := apiClient.InspectLicense(licenseKey)
	if err != nil {
		println(CrossIcon + " License is not valid")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	println(CheckMarkIcon + " License is valid")
	fmt.Printf("   ├──License key: %s\n", license.Key)
	fmt.Printf("   ├──License owner: %s\n", license.Owner.EmailAddress)
	fmt.Printf("   └──Product: %s\n", license.Product.Name)
	// fmt.Printf("   Repository: %s\n", license.Product.Repository)
	// fmt.Printf("   Registry: %s\n", license.Product.Registry)

	// Step 4 - Download the app image
	err = docker.PullImageWithIdentityToken(
		license.Product.Repository,
		license.Product.Registry,
		license.Owner.EmailAddress,
		licenseKey,
		false,
	)
	if err != nil {
		println(CrossIcon + " App download failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}
	println(CheckMarkIcon + " App downloaded")

	// Step 5 - Run the app
	err = docker.RunApp(license.Product.Repository, license.Product.Registry)
}
