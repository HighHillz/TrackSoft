/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"tracker/db"
	"tracker/ui"

	"github.com/gen2brain/beeep"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := db.GetAllProjects()
		if err != nil {
			fmt.Printf("Error fetching projects: %v\n", err)
			return
		}

		if len(projects) == 0 {
			fmt.Println("No projects found. Use 'tracker add' to get started!")
			return
		}

		fmt.Println(ui.HeaderStyle.Render("PROJECT TRACKER DASHBOARD"))
		fmt.Println()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Type", "Due", "Health", "Status"})
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t")
		table.SetNoWhiteSpace(true)

		anyOverdue := false
		for _, p := range projects {
			var status string
			if p.IsCompleted {
				status = "Done"
			} else if p.Overdue() {
				status = ui.OverdueStyle.Render("OVERDUE")
				anyOverdue = true
			} else if p.IsNew() {
				status = ui.NewProjectStyle.Render("NEW")
			} else {
				status = "Pending"
			}

			health := fmt.Sprintf("%.1f", p.Health()*100) + "%"
			
			table.Append([]string{
				fmt.Sprintf("%d", p.ID),
				p.Name,
				string(p.Type),
				p.Deadline.Format("2006-01-02"),
				health,
				status,
			})
		}
		table.Render()

		if anyOverdue {
			beeep.Alert("Tracker Alert", "You have overdue projects!", "")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
