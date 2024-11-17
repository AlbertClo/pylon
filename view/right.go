package view

import "fmt"

func RenderRightContent(m Model) string {
	return fmt.Sprintf(m.GetMessage())
}
