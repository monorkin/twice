package setup

import "strings"

func (m *model) installationView(s *strings.Builder) {
	s.WriteString(m.spinner.View())
	s.WriteString("Installing...\n")
}
