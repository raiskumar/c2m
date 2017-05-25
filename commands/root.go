package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "c2m",
	Short: "c2m is a couchbase cluster manager",
	Long: ` A compact and fast tool 
			to monitor couchbase cluster written in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" inside root.go")
	},
}
