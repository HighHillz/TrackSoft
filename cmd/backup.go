/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"tracker/db"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Export projects to a JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := db.GetAllProjects()
		if err != nil {
			fmt.Printf("Error fetching projects: %v\n", err)
			return
		}

		configDir, _ := os.UserConfigDir()
		backupPath := filepath.Join(configDir, "tracker", "backup.json")

		data, err := json.MarshalIndent(projects, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling data: %v\n", err)
			return
		}

		if err := os.WriteFile(backupPath, data, 0644); err != nil {
			fmt.Printf("Error writing backup: %v\n", err)
			return
		}

		fmt.Printf("Database backed up to: %s\n", backupPath)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
