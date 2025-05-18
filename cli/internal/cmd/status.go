package cmd

import "github.com/spf13/cobra"

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
	println("Running status command...")
}
