package setup

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	docker "github.com/monorkin/twice/cli/internal/docker"
)

type dockerStatusMsg struct {
	installed bool
	running   bool
	error     error
}

func (m model) checkDockerAfterDelayCmd(delay time.Duration) tea.Cmd {
	return tea.Tick(delay, func(t time.Time) tea.Msg {
		return m.checkDockerCmd()()
	})
}

func (m model) checkDockerCmd() tea.Cmd {
	return func() tea.Msg {
		installed := docker.IsInstalled()
		if !installed {
			return dockerStatusMsg{installed: installed, running: false}
		}

		running, err := docker.IsRunning()
		if err != nil {
			return dockerStatusMsg{installed: installed, running: running, error: err}
		}

		return dockerStatusMsg{installed: installed, running: running}
	}
}
