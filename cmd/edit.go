/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"time"
	"tracker/db"

	"github.com/spf13/cobra"
	"github.com/tj/go-naturaldate"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Project ID required. Usage: track edit <ID> --due \"new date\"")
			return
		}

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid project ID format.")
			return
		}

		dueStr, _ := cmd.Flags().GetString("due")
		if dueStr == "" {
			fmt.Println("Error: --due flag is required for editing constraints right now.")
			return
		}

		deadline, err := naturaldate.Parse(dueStr, time.Now())
		if err != nil {
			fmt.Printf("Error parsing new date '%s': %v\n", dueStr, err)
			return
		}

		if err := db.UpdateProjectDeadline(id, deadline); err != nil {
			fmt.Printf("Error updating project: %v\n", err)
			return
		}

		fmt.Printf("Project %d deadline updated to %s (%s).\n", id, deadline.Format("2006-01-02"), dueStr)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("due", "d", "", "New due date (natural language like 'next friday' or YYYY-MM-DD)")
}
