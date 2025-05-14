package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	// Step 5 - Configuration
	reader := bufio.NewReader(os.Stdin)

	domain := ""
	enableHTTPS := false

	for {
		fmt.Print("Enter the domain (e.g. example.com) where you'll run the app: ")
		domain, _ := reader.ReadString('\n')
		domain = strings.TrimSpace(domain)

		fmt.Print("Do you want to enable HTTPS? (yes/NO): ")
		httpsAnswer, _ := reader.ReadString('\n')
		httpsAnswer = strings.TrimSpace(strings.ToLower(httpsAnswer))
		enableHTTPS := httpsAnswer == "yes" || httpsAnswer == "y"

		fmt.Printf("   ├──Domain: %s\n", domain)
		fmt.Printf("   └──HTTPS: %t\n", enableHTTPS)

		fmt.Print("Is this correct? (yes/NO): ")
		answer, _ := reader.ReadString('\n')
		correct := strings.TrimSpace(strings.ToLower(answer))
		if correct == "yes" || correct == "y" {
			break
		}
	}

	// Step 6 - Run the app
	err = docker.RunApp(
		license.Product.Repository,
		license.Product.Registry,
		license.Owner.EmailAddress,
		domain,
		enableHTTPS,
	)
	if err != nil {
		println(CrossIcon + " App run failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	println(CheckMarkIcon + " App is running")
}
