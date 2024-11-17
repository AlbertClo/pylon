package color

import "github.com/charmbracelet/lipgloss"

const (
	whiteColorString  = "#ffffff"
	purpleColorString = "#c084fc"
	greenColorString  = "#10B981"
	grayColorString   = "#64748b"
	blueColorString   = "#3b82f6"
	redColorString    = "#ef4444"
)

var (
	white  = lipgloss.Color(whiteColorString)
	purple = lipgloss.Color(purpleColorString)
	green  = lipgloss.Color(greenColorString)
	gray   = lipgloss.Color(grayColorString)
	blue   = lipgloss.Color(blueColorString)
	red    = lipgloss.Color(redColorString)
)

var (
	Primary   = green
	Secondary = white
	Muted     = gray
	Error     = red
)

var (
	ErrorStyle = lipgloss.NewStyle().
			Foreground(red).
			Bold(true)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true)

	TextStyle = lipgloss.NewStyle().
			Foreground(gray)
)
