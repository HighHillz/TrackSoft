/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"tracker/cmd"
	"tracker/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database: %v\n", err)
		os.Exit(1)
	}
	cmd.Execute()
}
