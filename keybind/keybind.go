package keybind

import (
	"github.com/charmbracelet/bubbles/key"
)

type Shortcuts struct {
	Up        key.Binding
	Down      key.Binding
	Enter     key.Binding
	Quit      key.Binding
	StartStop key.Binding
}

func New() Shortcuts {
	return Shortcuts{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		StartStop: key.NewBinding(
			key.WithKeys("s", " "),
			key.WithHelp("s/space", "start/stop"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q/esc/ctrl+c", "quit"),
		),
	}
}
