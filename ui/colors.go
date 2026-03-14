package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(4).
		PaddingRight(4)

	OverdueStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#E88388")) // Red-ish

	NewProjectStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00D7FF")) // Electric Blue

	SuccessStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A8CC8C")).
		Bold(true)
)

func HealthBar(percent float64) string {
	width := 10
	filled := int(percent * float64(width))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	
	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "|"
		} else {
			bar += " "
		}
	}
	bar += "]"

	color := "#A8CC8C" // Green
	if percent < 0.3 {
		color = "#E88388" // Red
	} else if percent < 0.6 {
		color = "#D29034" // Yellow/Orange
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(bar)
}
