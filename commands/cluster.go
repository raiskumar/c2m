package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Couchbase Cluster Manager",
	Long: `Tool to monitor your Couchbase Cluster
            Command format
			$c2m command subcommand --flag=xyz
            `,
	Run: cluster,
}

func cluster(cmd *cobra.Command, args []string) {
	fmt.Println(" inside runCluster method")
	fmt.Println(args)
}
