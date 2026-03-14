/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"tracker/db"
	"tracker/models"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tj/go-naturaldate"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		dueStr, _ := cmd.Flags().GetString("due")
		projectType, _ := cmd.Flags().GetString("type")

		if name == "" {
			if len(args) > 0 {
				name = args[0]
			} else {
				fmt.Println("Error: Project name is required.")
				return
			}
		}

		if projectType == "CLI" && !cmd.Flags().Changed("type") {
			prompt := promptui.Select{
				Label: "Select Project Type",
				Items: []string{"UI/UX", "Web", "CLI", "GUI", "Script", "Core"},
			}
			_, result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Selection failed: %v\n", err)
				return
			}
			projectType = result
		}

		// Phase 3: Natural Date Parsing
		var deadline time.Time
		if dueStr == "" {
			dueStr = "next friday" // Default if not provided
		}
		
		deadline, err := naturaldate.Parse(dueStr, time.Now())
		if err != nil {
			fmt.Printf("Error parsing date '%s': %v\n", dueStr, err)
			return
		}

		newProject := &models.Project{
			Name:      name,
			Type:      models.ProjectType(projectType),
			Deadline:  deadline,
			CreatedAt: time.Now(),
		}

		if err := db.AddProject(newProject); err != nil {
			fmt.Printf("Error saving project: %v\n", err)
			return
		}

		fmt.Printf("\033[1;32m\033[1mProject '%s' saved successfully!\033[0m\n", name)
		fmt.Printf("Due: %s (%s)\n", deadline.Format("2006-01-02"), dueStr)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Name of the project")
	addCmd.Flags().StringP("due", "d", "", "Due date (natural language like 'next friday' or YYYY-MM-DD)")
	addCmd.Flags().StringP("type", "t", "CLI", "Project type (UI/UX, Web, CLI, etc.)")
}
