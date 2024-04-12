package styling

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	errStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ff0000"))
	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}).
				Bold(true)

	TableBaseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}).
			Bold(false)
)

func StyleErrMsg(err error) string {
	return errStyle.Render(err.Error())
}
