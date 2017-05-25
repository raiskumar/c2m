package main

import (
	"fmt"

	"os"

	"github.com/raiskumar/c2m/commands"
)

func main() {
	fmt.Println("inside main")
	if err := commands.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := commands.ClusterCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
