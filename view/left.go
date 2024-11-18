package view

import (
	"fmt"
	"github.com/AlbertClo/pylon/types"
	"strings"
)

func RenderLeftContent(m Model) string {
	var menuStr strings.Builder

	for i, item := range m.GetMenuItems() {
		if i == m.GetSelectedItem() {
			// Selected item
			menuStr.WriteString(fmt.Sprintf("> %s %s\n", liveness(item), item.Name))
		} else {
			// Unselected item
			menuStr.WriteString(fmt.Sprintf("  %s %s\n", liveness(item), item.Name))
		}
	}

	return menuStr.String()
}

func livenessIndicator(item types.MenuItem) string {
	if item.Running {
		return "🟢"
	}
	return "⚫"
}
