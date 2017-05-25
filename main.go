package main

import (
	"fmt"

	"os"

	"github.com/raiskumar/c2m/commands"
)

func main() {
	fmt.Println("inside main")
	err := commands.CcbCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
