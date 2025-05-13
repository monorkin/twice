package cmd

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const TwiceLogoAscii = `
████████╗██╗    ██╗██╗ ██████╗███████╗
╚══██╔══╝██║    ██║██║██╔════╝██╔════╝
   ██║   ██║ █╗ ██║██║██║     █████╗
   ██║   ██║███╗██║██║██║     ██╔══╝
   ██║   ╚███╔███╔╝██║╚██████╗███████╗
   ╚═╝    ╚══╝╚══╝ ╚═╝ ╚═════╝╚══════╝
`

var (
	PrimaryColor   = lipgloss.Color("#7D56F4")
	SecondaryColor = lipgloss.Color("#FAFAFA")
	ErrorColor     = lipgloss.Color("#FF0000")
	SuccessColor   = lipgloss.Color("#00AA00")
	NeutralColor   = lipgloss.Color("#626262")

	CheckmarkIcon = "✅"
	CrossIcon     = "❌"

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(SecondaryColor).
			Background(PrimaryColor).
			Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor)

	LogoStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor)

	HelpStyle = lipgloss.NewStyle().
			Foreground(NeutralColor)

	SpinnerStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor)

	AppStyle = lipgloss.NewStyle().
			Padding(1, 2, 1, 2)
)

func RenderLogo() string {
	return LogoStyle.Render(TwiceLogoAscii)
}

func RenderCenteredLogo(width int) string {
	logo := RenderLogo()
	logoLines := strings.Split(logo, "\n")

	var centeredLogo strings.Builder
	for _, line := range logoLines {
		if line != "" {
			centeredLine := lipgloss.PlaceHorizontal(width, lipgloss.Center, line)
			centeredLogo.WriteString(centeredLine + "\n")
		} else {
			centeredLogo.WriteString("\n")
		}
	}

	return centeredLogo.String()
}
