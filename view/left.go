package view

import "fmt"

func RenderLeftContent(m Model) string {
	return fmt.Sprintf(`Counter: %d`, m.GetCounter())
}
