package setup

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup [license-key]",
		Short: "Set up Twice with your license key",
		Long:  `Configure Twice and install products using your license key`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var licenseKey string
			if len(args) > 0 {
				licenseKey = args[0]
			}

			program := tea.NewProgram(
				initialModel(licenseKey),
				tea.WithAltScreen(),
			)

			finalModel, err := program.Run()
			if err != nil {
				fmt.Printf("Error running setup: %v\n", err)
				os.Exit(1)
			}

			model, ok := finalModel.(model)
			if !ok {
				fmt.Println("Error: model is not of type 'model'")
				os.Exit(1)
			}

			if model.err != nil {
				fmt.Printf("Error: %v\n", model.err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
