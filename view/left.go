package view

import (
	"fmt"
	"strings"
)

func RenderLeftContent(m Model) string {
	var menuStr strings.Builder

	for i, item := range m.GetMenuItems() {
		if i == m.GetSelectedItem() {
			// Selected item
			menuStr.WriteString(fmt.Sprintf("> %s\n", item.Name))
		} else {
			// Unselected item
			menuStr.WriteString(fmt.Sprintf("  %s\n", item.Name))
		}
	}

	return menuStr.String()
}
