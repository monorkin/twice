package setup

import (
	"fmt"
	"strings"

	cmd "github.com/monorkin/twice/cli/cmd"
)

func (m *model) dockerView(s *strings.Builder) {
	s.WriteString(cmd.SubtitleStyle.Render("Docker Setup"))
	s.WriteString("\n")

	if m.checkingDocker {
		s.WriteString(m.spinner.View())
		s.WriteString("Checking Docker installation...\n")
		return
	}

	installedStatusIcon := cmd.CheckmarkIcon
	if !m.dockerInstalled {
		installedStatusIcon = cmd.CrossIcon
	}

	runningStatusIcon := cmd.CheckmarkIcon
	if !m.dockerRunning {
		runningStatusIcon = cmd.CrossIcon
	}

	s.WriteString(fmt.Sprintf("%s Docker Installed\n", installedStatusIcon))
	s.WriteString(fmt.Sprintf("%s Docker Running\n", runningStatusIcon))
	s.WriteString("\n")

	if m.dockerInstalled && m.dockerRunning {
		s.WriteString(m.spinner.View())
		s.WriteString("Continuing to next step...\n")
		return
	}
}
