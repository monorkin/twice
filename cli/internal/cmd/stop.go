package cmd

import (
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

func runStopCmd(productQuery string) {
	println("Stopping product:", productQuery)
}
