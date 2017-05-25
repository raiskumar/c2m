package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var ClusterCommand = &cobra.Command{
	Use:   "c2m",
	Short: "Couchbase Cluster Manager",
	Long: `Tool to monitor your Couchbase Cluster
            Command format
			$c2m command subcommand --flag=xyz
            `,
	Run: clusterRun,
}

func clusterRun(cmd *cobra.Command, args []string) {
	fmt.Println(" inside runCluster method")
	fmt.Println(strings.Join(args, ""))
}

func init() {
	RootCommand.AddCommand(ClusterCommand)
}
