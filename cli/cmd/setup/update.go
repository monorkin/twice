package setup

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.processKeyMsg(msg)

	case spinner.TickMsg:
		return m.processSpinnerTickMsg(msg)

	case tea.WindowSizeMsg:
		return m.processWindowSizeMsg(msg)

	case dockerStatusMsg:
		return m.processDockerStatusMsg(msg)

	case switchViewMsg:
		return m.processSwitchViewMsg(msg)

	case fetchedProductsMsg:
		return m.processFetchedProductsMsg(msg)
	}

	return m, tea.Batch(m.spinner.Tick)
}

func (m model) processKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case ProductView:
		var cmd tea.Cmd
		m.productList, cmd = m.productList.Update(msg)

		switch {
		case key.Matches(msg, m.keys.Select):
			if i, ok := m.productList.SelectedItem().(Product); ok {
				return m, m.selectProductCmd(i)
			}
		}
		return m, cmd
	}

	switch {
	case key.Matches(msg, m.keys.Help):
		m.help.ShowAll = !m.help.ShowAll
	case key.Matches(msg, m.keys.Quit):
		m.quitting = true
		return m, tea.Quit
	}

	return m, nil
}

func (m model) processSpinnerTickMsg(msg spinner.TickMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	new_spinner, cmd := m.spinner.Update(msg)
	m.spinner = new_spinner
	return m, cmd
}

func (m model) processWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height
	m.productList.SetSize(m.width, m.height-14)
	return m, nil
}

func (m model) processDockerStatusMsg(msg dockerStatusMsg) (tea.Model, tea.Cmd) {
	m.checkingDocker = false
	m.dockerInstalled = msg.installed
	m.dockerRunning = msg.running

	if msg.error != nil {
		m.err = msg.error
		return m, tea.Quit
	}

	if m.dockerInstalled && m.dockerRunning {
		return m, m.runCmdWithDelay(1*time.Second, m.switchToViewCmd(ProductView))
	}

	return m, nil
}

func (m model) processSwitchViewMsg(msg switchViewMsg) (tea.Model, tea.Cmd) {
	m.currentView = msg.view

	switch m.currentView {
	case ProductView:
		return m, m.fetchProductsCmd()
	}

	return m, nil
}

func (m model) processFetchedProductsMsg(msg fetchedProductsMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		m.err = msg.err
		return m, tea.Quit
	}

	products := make([]list.Item, len(msg.products))
	for i, apiProduct := range msg.products {
		products[i] = Product{ID: apiProduct.ID, Name: apiProduct.Name}
	}

	m.products = msg.products
	m.productList.SetItems(products)
	m.productsLoaded = true

	var cmd tea.Cmd
	m.productList, cmd = m.productList.Update(nil)

	return m, cmd
}
