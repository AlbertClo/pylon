package color

import "github.com/charmbracelet/lipgloss"

const (
	purpleColorString = "#c084fc"
	greenColorString  = "#10B981"
	grayColorString   = "#64748b"
	blueColorString   = "#3b82f6"
	redColorString    = "#ef4444"
)

var (
	Purple = lipgloss.Color(purpleColorString)
	Green  = lipgloss.Color(greenColorString)
	Gray   = lipgloss.Color(grayColorString)
	Blue   = lipgloss.Color(blueColorString)
	Red    = lipgloss.Color(redColorString)
)

var (
	Primary   = Green
	Secondary = Purple
	Muted     = Gray
	Error     = Red
)

var (
	ErrorStyle = lipgloss.NewStyle().
			Foreground(Red).
			Bold(true)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true)

	TextStyle = lipgloss.NewStyle().
			Foreground(Gray)
)
