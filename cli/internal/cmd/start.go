package cmd

import (
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts an installed product",
		Long:  `Starts an installed, but stopped, product.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runStartCmd(args[0])
		},
	}

	return cmd
}

func runStartCmd(productQuery string) {
	println("Starting product:", productQuery)
}
