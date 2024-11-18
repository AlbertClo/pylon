package view

import (
	"github.com/AlbertClo/pylon/types"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func RenderLeftContent(m Model) string {
	var menuStr strings.Builder

	normalTextStyle := lipgloss.NewStyle().
		Width(20)

	selectedTextStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("6")).
		Foreground(lipgloss.Color("0")).
		Width(20)

	indicatorStyle := lipgloss.NewStyle().
		PaddingLeft(0).
		Width(2)

	for i, item := range m.GetMenuItems() {
		indicator := indicatorStyle.Render(renderLiveness(item))

		if i == m.GetSelectedItem() {
			menuStr.WriteString(indicator + " " + selectedTextStyle.Render(" "+item.Name) + "\n")
		} else {
			menuStr.WriteString(indicator + " " + normalTextStyle.Render(" "+item.Name) + "\n")
		}
	}

	return menuStr.String()
}

func renderLiveness(item types.MenuItem) string {
	if item.Running {
		return "🟢"
	}
	return "⚫"
}
