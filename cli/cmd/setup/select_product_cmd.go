package setup

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/monorkin/twice/cli/internal/api"
)

type productSelectedMsg struct {
	product api.Product
}

func (m model) selectProductCmd(i int) tea.Cmd {
	return func() tea.Msg {
		product, ok := m.productList.Items[i].(Product)
		if !ok {
			return productSelectedMsg{}
		}
		return productSelectedMsg{product: product}
	}
}
