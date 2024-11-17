package view

import "fmt"

func RenderBottomContent(m Model) string {
	keys := m.GetKeys()
	return fmt.Sprintf(`Controls:
%s
%s
%s
%s`,
		keys.Quit.Help(),
		keys.Increment.Help(),
		keys.Decrement.Help(),
		keys.Reset.Help(),
	)
}
