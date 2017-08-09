package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CcbCmd = &cobra.Command{
	Use:   "c2m",
	Short: "c2m is a couchbase cluster manager",
	Long: ` c2m : Couchbase Cluster Manager
Helps in getting key insights of your Couchbase Cluster. 
Ideal for monitoring and managing all the nodes of your Couchbase cluster`,

	// Method will get called if user doesn't provide any command!
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" Please enter command as well ! ")
		fmt.Println(" like... $c2m node")
	},
}

// Initializes all available commands
func init() {
	CcbCmd.AddCommand(clusterCmd)
	CcbCmd.AddCommand(nodeCmd)
	CcbCmd.AddCommand(configCmd)
	CcbCmd.AddCommand(bucketCmd)
}
