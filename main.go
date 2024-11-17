package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlbertClo/pylon/keyboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type Model struct {
	spinner    spinner.Model
	counter    int
	message    string
	quitting   bool
	err        error
	keys       keyboard.Shortcuts
	windowSize struct {
		width  int
		height int
	}
	style lipgloss.Style
}

func initialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#c084fc"))
	return Model{
		spinner: s,
		counter: 0,
		message: "",
		keys:    keyboard.New(),
		style: lipgloss.NewStyle().
			PaddingTop(2).
			PaddingLeft(4).
			Foreground(lipgloss.Color("#10B981")),
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			m.quitting = true
			return m, tea.Quit
		}
		if key.Matches(msg, m.keys.Increment) {
			m.counter++
		}
		if key.Matches(msg, m.keys.Decrement) {
			m.counter--
		}
		if key.Matches(msg, m.keys.Reset) {
			m.counter = 0
		}
		if key.Matches(msg, m.keys.Start) {
			cmd := exec.Command("touch", "test")
			err := cmd.Run()
			if err != nil {
				m.err = err
			}
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.windowSize.width = msg.Width
		m.windowSize.height = msg.Height
		m.message = fmt.Sprintf("Window size: %d x %d", msg.Width, msg.Height)
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

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	str := fmt.Sprintf(`Counter: %d
Controls:
%s
%s
%s
%s
%s`,
		m.counter,
		m.keys.Quit.Help(),
		m.keys.Increment.Help(),
		m.keys.Decrement.Help(),
		m.keys.Reset.Help(),
		m.message,
	)

	if m.quitting {
		return str + "\n"
	}

	return m.style.Width(m.windowSize.width).Height(m.windowSize.height).Render(str)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
