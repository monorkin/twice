package setup

import (
	"fmt"
	"strings"

	cmd "github.com/monorkin/twice/cli/cmd"
)

type Product struct {
	ID   int64
	Name string
}

func (p Product) Title() string       { return p.Name }
func (p Product) Description() string { return fmt.Sprintf("ID: %d", p.ID) }
func (p Product) FilterValue() string { return p.Name }

func (m *model) productView(s *strings.Builder) {
	if !m.productsLoaded {
		s.WriteString(cmd.SubtitleStyle.Render("Select a products"))
		s.WriteString("\n")
		s.WriteString(m.spinner.View())
		s.WriteString("Fetching available products...\n")
		return
	}

	s.WriteString(cmd.SubtitleStyle.Render("Select a products"))
	s.WriteString("\n")
	s.WriteString(m.productList.View())
	return
}
