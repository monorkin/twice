package setup

import (
	"strings"

	cmd "github.com/monorkin/twice/cli/cmd"
)

func (m model) View() string {
	if m.quitting {
		return "Quitting setup...\n"
	}

	var s strings.Builder
	s.WriteString(cmd.RenderLogo())
	s.WriteString("\n")

	switch m.currentView {
	case DockerView:
		m.dockerView(&s)
	case ProductView:
		m.productView(&s)
	case InstallationView:
		m.installationView(&s)
	default:
		return cmd.ErrorStyle.Render("Unknown step")
	}

	s.WriteString("\n")
	s.WriteString(m.help.View(m.keys))

	return s.String()
}
