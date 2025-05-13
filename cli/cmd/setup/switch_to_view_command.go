package setup

import (
	tea "github.com/charmbracelet/bubbletea"
)

type switchViewMsg struct {
	view View
}

func (m model) switchToViewCmd(view View) tea.Cmd {
	return func() tea.Msg {
		return switchViewMsg{view: view}
	}
}
