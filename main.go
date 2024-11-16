package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
)

type errMsg error

type model struct {
	spinner  spinner.Model
	counter  int
	quitting bool
	err      error
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("q", "quit"),
)

var incrementKeys = key.NewBinding(
	key.WithKeys("up", "k"),
	key.WithHelp("↑/k", "increment"),
)

var decrementKeys = key.NewBinding(
	key.WithKeys("down", "j"),
	key.WithHelp("↓/j", "decrement"),
)

var resetKeys = key.NewBinding(
	key.WithKeys("r"),
	key.WithHelp("r", "reset"),
)

var startKeys = key.NewBinding(
	key.WithKeys("s"),
	key.WithHelp("s", "start"),
)

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#c084fc"))
	return model{spinner: s, counter: 0}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit

		}
		if key.Matches(msg, incrementKeys) {
			m.counter++
		}
		if key.Matches(msg, decrementKeys) {
			m.counter--
		}
		if key.Matches(msg, resetKeys) {
			m.counter = 0
		}
		if key.Matches(msg, startKeys) {
			cmd := exec.Command("touch", "test")
			cmd.Run()
		}
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	str := fmt.Sprintf(`
   Counter: %d

   Controls:
   %s
   %s
   %s
   %s`,
		m.counter,
		incrementKeys.Help(),
		decrementKeys.Help(),
		startKeys.Help(),
		quitKeys.Help())

	if m.quitting {
		return str + "\n"
	}

	return str
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
