package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/monorkin/twice/cli/internal/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "twice",
		Short: "Twice CLI - A client for interacting with the Twice distribution system",
		Long:  `A command line interface for the Twice distribution system.`,
	}

	rootCmd.AddCommand(cmd.NewSetupCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
