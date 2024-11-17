package view

import "fmt"

func RenderRightContent(m Model) string {
	return fmt.Sprintf(`Window Size: %s`, m.GetMessage())
}
