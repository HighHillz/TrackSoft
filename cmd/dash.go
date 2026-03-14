/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"tracker/db"
	"tracker/models"
	"tracker/ui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type item struct {
	project models.Project
}

func (i item) Title() string       { return i.project.Name }
func (i item) Description() string {
	health := i.project.Health()
	return fmt.Sprintf("%s | Due: %s %s", string(i.project.Type), i.project.Deadline.Format("2006-01-02"), ui.HealthBar(health))
}
func (i item) FilterValue() string { return i.project.Name }

type model struct {
	list     list.Model
	projects []models.Project
	err      error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "d" {
			i, ok := m.list.SelectedItem().(item)
			if ok {
				db.RemoveProject(i.project.ID)
				// Simplified: just refresh or remove from list
				m.list.RemoveItem(m.list.Index())
			}
		}
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Margin(1, 2).GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().Margin(1, 2).Render(m.list.View())
}

var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "Interactive TUI dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := db.GetAllProjects()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		items := make([]list.Item, len(projects))
		for i, p := range projects {
			items[i] = item{project: p}
		}

		m := model{
			list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
			projects: projects,
		}
		m.list.Title = "Projects"

		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, it failed: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
