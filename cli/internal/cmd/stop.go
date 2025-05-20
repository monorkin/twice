package cmd

import (
	"fmt"
	"os"

	"github.com/monorkin/twice/cli/internal/config"
	"github.com/monorkin/twice/cli/internal/docker"
	"github.com/spf13/cobra"
)

func NewStopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop <ID>",
		Short: "Stops a running product",
		Long:  `Stops a running product, with the given ID, without uninstalling it.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runStopCmd(args[0])
		},
	}

	return cmd
}

func runStopCmd(id string) {
	var product *config.Product
	for _, p := range cfg.Products {
		if p.ContainerName() == id {
			product = &p
		}
	}

	if product == nil {
		println("Product not found:", id)
		os.Exit(1)
	}

	err := docker.StopProductContainer(product)
	if err != nil {
		println(CrossIcon + " Couldn't stop product")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	println(CheckMarkIcon + " Product stopped")
}
