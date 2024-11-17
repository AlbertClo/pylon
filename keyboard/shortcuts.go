package keyboard

import "github.com/charmbracelet/bubbles/key"

type Shortcuts struct {
	Quit      key.Binding
	Increment key.Binding
	Decrement key.Binding
	Reset     key.Binding
	Start     key.Binding
}

func New() Shortcuts {
	return Shortcuts{
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q/esc/ctrl+c", "quit"),
		),
		Increment: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "increment"),
		),
		Decrement: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "decrement"),
		),
		Reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset"),
		),
		Start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start"),
		),
	}
}
