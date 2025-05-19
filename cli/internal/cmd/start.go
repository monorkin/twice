package cmd

import (
	"fmt"
	"os"

	"github.com/monorkin/twice/cli/internal/config"
	"github.com/monorkin/twice/cli/internal/docker"
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <ID>",
		Short: "Starts an installed product",
		Long:  `Starts an installed, but stopped, product with the given ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runStartCmd(args[0])
		},
	}

	return cmd
}

func runStartCmd(id string) {
	var product *config.ProductConfig
	for _, p := range cfg.Products {
		if p.ContainerName() == id {
			product = &p
		}
	}

	if product == nil {
		println("Product not found:", id)
		os.Exit(1)
	}

	err := docker.StartProductContainer(product)
	if err != nil {
		println(CrossIcon + " Couldn't start product")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	println(CheckMarkIcon + " Product started")
}
