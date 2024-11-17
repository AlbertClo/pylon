package main

import (
	"fmt"
	"github.com/AlbertClo/pylon/color"
	"github.com/AlbertClo/pylon/layout"
	"github.com/AlbertClo/pylon/view"
	"os"
	"os/exec"

	"github.com/AlbertClo/pylon/keybind"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Model struct {
	spinner    spinner.Model
	counter    int
	message    string
	quitting   bool
	err        error
	keys       keybind.Shortcuts
	windowSize struct {
		width  int
		height int
	}
}

func (m Model) GetCounter() int {
	return m.counter
}

func (m Model) GetMessage() string {
	return m.message
}

func (m Model) GetKeys() keybind.Shortcuts {
	return m.keys
}

func initialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(color.Primary)
	return Model{
		spinner: s,
		counter: 0,
		message: "",
		keys:    keybind.New(),
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
	case error:
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

	return layout.New(m.windowSize.width, m.windowSize.height).
		SetLeftContent(view.RenderLeftContent(m)).
		SetRightContent(view.RenderRightContent(m)).
		SetBottomContent(view.RenderBottomContent(m)).
		Render()
}
