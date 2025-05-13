package setup

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/monorkin/twice/cli/internal/api"
)

type fetchedProductsMsg struct {
	products []api.Product
	err      error
}

func (m model) fetchProductsCmd() tea.Cmd {
	return func() tea.Msg {
		products, err := m.apiClient.ListProducts()
		return fetchedProductsMsg{products: products, err: err}
	}
}
