package commands

import (
	"encoding/json"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Prints cluster details",
	Long: `Prints out details of the Couchbase Cluster
            Command format
			$c2m command subcommand --flag=xyz
            `,
	Run: cluster,
}

// base_url/pools/default/buckets
func cluster(cmd *cobra.Command, args []string) {
	common.ValidateCommand(NodeURL)
	uri := NodeURL + "/pools/default"
	//uri = "http://mocky.io/v2/5986c32d1100009c00fcbe4a" // Test URL

	contents := common.GetRestContent(uri, UserID, Password)

	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)

	cluster := getClusterDetails(obj)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cluster.GetHeaders())

	table.Append(cluster.ToString())
	table.Render()

}

func getClusterDetails(resp vo.PoolResp) vo.Cluster {
	var nodeStatus map[string]string
	nodeStatus = make(map[string]string)
	for i := 0; i < len(resp.Nodes); i++ {
		nodeStatus[resp.Nodes[i].CouchAPIBase] = resp.Nodes[i].ClusterMembership
	}

	cluster := vo.Cluster{
		Name:           resp.ClusterName,
		NodesStatus:    nodeStatus,
		DiskUsedByData: int64(resp.StorageTotals.Hdd.UsedByData),
		FreeDisk:       resp.StorageTotals.Hdd.Free,
		TotalDisk:      resp.StorageTotals.Hdd.Total,
		RAMUsedByData:  int64(resp.StorageTotals.RAM.UsedByData),
		TotalRAM:       resp.StorageTotals.RAM.Total,
		UsedRAM:        resp.StorageTotals.RAM.Used}
	return cluster
}
