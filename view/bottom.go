package view

import "fmt"

func RenderBottomContent(m Model) string {
	keys := m.GetKeys()
	return fmt.Sprintf(`Controls:

%s
%s
%s
%s`,
		keys.Up.Help(),
		keys.Down.Help(),
		keys.Enter.Help(),
		keys.Quit.Help(),
	)
}
