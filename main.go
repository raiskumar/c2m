package main

import (
	"fmt"
	"os"

	"github.com/raiskumar/c2m/commands"
)

// Main method used for initializing Cobra / CLI
func main() {
	err := commands.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
