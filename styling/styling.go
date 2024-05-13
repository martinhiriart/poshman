package styling

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	errStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ff0000"))
	successStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#00b300", Dark: "#00ff00"})
	statusStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#0099cc", Dark: "#00ffff"})
)

func StyleErrMsg(err error) string {
	return strings.TrimSpace(errStyle.Render(err.Error()))
}

func StyleSuccessMsg(text string) string {
	return strings.TrimSpace(successStyle.Render(text))
}

func StyleStatusMsg(text string) string {
	return strings.TrimSpace(statusStyle.Render(text))
}

// func GetTableHeaderStyle() lipgloss.Style {
// 	return tableHeaderStyle
// }
// func GetTableBaseStyle() lipgloss.Style {
// 	return tableBaseStyle
// }
