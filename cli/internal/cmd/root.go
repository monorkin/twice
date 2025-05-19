package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/monorkin/twice/cli/internal/config"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "twice",
		Short: "Twice CLI - A client for interacting with the Twice distribution system",
		Long:  `A command line interface for the Twice distribution system.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error
			cfg, err = config.LoadOrCreateConfig()
			if err != nil {
				fmt.Println("Couldn't load or create a config")
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(NewSetupCmd())
	rootCmd.AddCommand(NewStatusCmd())
	rootCmd.AddCommand(NewStartCmd())
	rootCmd.AddCommand(NewStopCmd())

	return rootCmd
}
