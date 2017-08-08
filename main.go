package main

import (
	"fmt"
	"os"

	"github.com/raiskumar/c2m/commands"
)

// Main method of the project
func main() {
	err := commands.CcbCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
