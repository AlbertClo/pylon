package layout

import (
	"github.com/AlbertClo/pylon/color"
	"github.com/charmbracelet/lipgloss"
)

type Layout struct {
	width         int
	height        int
	leftContent   string
	rightContent  string
	bottomContent string
	leftPaneWidth int
	bottomHeight  int
	heightOffset  int
}

func New(width, height int) *Layout {
	return &Layout{
		width:         width,
		height:        height,
		leftPaneWidth: 30,
		bottomHeight:  8,
		heightOffset:  1,
	}
}

func (l *Layout) SetLeftContent(content string) *Layout {
	l.leftContent = content
	return l
}

func (l *Layout) SetRightContent(content string) *Layout {
	l.rightContent = content
	return l
}

func (l *Layout) SetBottomContent(content string) *Layout {
	l.bottomContent = content
	return l
}

func (l *Layout) Render() string {
	leftStyle := lipgloss.NewStyle().
		PaddingTop(1).
		PaddingLeft(3).
		Width(l.leftPaneWidth).
		Height(l.height - l.bottomHeight - l.heightOffset).
		Foreground(color.Secondary).
		BorderRight(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(color.Secondary)

	rightStyle := lipgloss.NewStyle().
		PaddingTop(1).
		PaddingLeft(3).
		Width(l.width - l.leftPaneWidth).
		Height(l.height - l.bottomHeight - l.heightOffset).
		Foreground(color.Secondary)

	bottomStyle := lipgloss.NewStyle().
		PaddingTop(1).
		PaddingLeft(3).
		Width(l.width).
		Height(l.bottomHeight).
		Foreground(color.Secondary).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(color.Secondary)

	leftPane := leftStyle.Render(l.leftContent)
	rightPane := rightStyle.Render(l.rightContent)
	bottomPane := bottomStyle.Render(l.bottomContent)

	topSection := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPane,
		rightPane,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		topSection,
		bottomPane,
	)
}
