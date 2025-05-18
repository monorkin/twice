package cmd

import (
	"os"

	"github.com/monorkin/twice/cli/internal/docker"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NewStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Displays the status of all installed products",
		Long:  `Displays the status of all installed products, including their health and version information.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runStatusCmd()
		},
	}

	return cmd
}

func runStatusCmd() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Product", "License", "Auth server", "Status"})

	for _, product := range cfg.Products {
		status := "Unknown"

		running, err := docker.IsProductRunning(&product)
		if err == nil {
			if running {
				status = "Running"
			} else {
				status = "Stopped"
			}
		}

		data := []string{product.Product, product.LicenseKey, product.AuthServer, status}
		table.Append(data)
	}

	table.Render()
}
