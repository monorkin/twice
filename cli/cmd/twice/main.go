package main

import (
	"fmt"
	"os"

	"github.com/monorkin/twice/cli/internal/cmd"
	"github.com/monorkin/twice/cli/internal/config"
)

var cfg *config.Config

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
