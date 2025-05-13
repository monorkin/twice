package setup

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tea.WindowSize(),
		m.runCmdWithDelay(300*time.Millisecond, m.checkDockerCmd()),
	)
}
