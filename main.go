package main

import (
	"fmt"
	"github.com/AlbertClo/pylon/color"
	"github.com/AlbertClo/pylon/keybind"
	"github.com/AlbertClo/pylon/layout"
	"github.com/AlbertClo/pylon/view"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
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
	menuItems  []string
	selected   int
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

func (m Model) GetMenuItems() []string {
	return m.menuItems
}

func (m Model) GetSelectedItem() int {
	return m.selected
}

func initialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(color.Primary)
	return Model{
		spinner:   s,
		counter:   0,
		message:   "",
		keys:      keybind.New(),
		menuItems: []string{"Start", "Stop", "Quit"},
		selected:  0, // Start with first item selected
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			if m.selected > 0 {
				m.selected--
			} else {
				m.selected = len(m.menuItems) - 1
			}
			return m, nil
		case key.Matches(msg, m.keys.Down):
			if m.selected < len(m.menuItems)-1 {
				m.selected++
			} else {
				m.selected = 0
			}
			return m, nil
		case key.Matches(msg, m.keys.Enter):
			switch m.menuItems[m.selected] {
			case "Start":
				m.message = "Starting..."
			case "Stop":
				m.message = "Stopping..."
			case "Quit":
				m.quitting = true
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.windowSize.width = msg.Width
		m.windowSize.height = msg.Height
		m.message = fmt.Sprintf("Window size: %d x %d", msg.Width, msg.Height)
		return m, nil
	case error:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
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
