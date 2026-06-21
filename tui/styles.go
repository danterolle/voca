package tui

import "github.com/charmbracelet/lipgloss"

var (
	headerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("203")).Bold(true)
	subtleStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	inputLabelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("215")).Bold(true)
	outputLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("209")).Bold(true)
	helpStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	inputBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("215")).
			Padding(0, 1)

	outputBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("209")).
			Padding(0, 1)
)
