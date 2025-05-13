package setup

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) runCmdWithDelay(delay time.Duration, cmd tea.Cmd) tea.Cmd {
	return tea.Tick(delay, func(t time.Time) tea.Msg {
		return cmd()
	})
}
