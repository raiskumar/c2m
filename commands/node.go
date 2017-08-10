package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m node
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Prints important details/metadata of nodes of the cluster!",
	Long: `Command which prints the node related details 
            Command format
			$./c2m node --flag=xyz
            `,
	Run: node,
}

// Node command
// http://<ip>:8091/pools/default gives an insight into computing resources consumed per node
func node(cmd *cobra.Command, args []string) {
	fmt.Println("URI=", os.Getenv("URI"))

	/*for index, value := range os.Environ() {
		name := strings.Split(value, "=") // split by = sign

		fmt.Printf("[%d] %s : %v\n", index, name[0], name[1])
	}*/
	contents := common.GetRestContent(os.Getenv("URI"), os.Getenv("USER"), os.Getenv("PASS"))
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)

	nodes := getAllNodes(obj)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(nodes[0].GetHeaders())

	for _, val := range nodes {
		table.Append(val.ToString())
	}
	table.Render()

}

// Parse REST response and return the list of nodes struct
func getAllNodes(resp vo.PoolResp) []vo.Node {
	len := len(resp.Nodes)
	var nodes []vo.Node
	for i := 0; i < len; i++ {
		n := vo.Node{
			HostName:          resp.Nodes[i].Hostname,
			Url:               resp.Nodes[i].CouchAPIBase,
			Status:            resp.Nodes[i].Status,
			Services:          resp.Nodes[i].Services,
			GetHits:           resp.Nodes[i].InterestingStats.GetHits,
			DocumentCount:     resp.Nodes[i].InterestingStats.CurrItems,
			DiskUsedByData:    resp.Nodes[i].InterestingStats.CouchDocsActualDiskSize, // couch_docs_actual_disk_size
			RAMUsed:           resp.Nodes[i].InterestingStats.MemUsed,                 // interestingstats.mem_used
			FreeRAM:           resp.Nodes[i].MemoryFree,
			TotalRAM:          resp.Nodes[i].MemoryTotal,
			ClusterMembership: resp.Nodes[i].ClusterMembership}
		nodes = append(nodes, n)
	}
	return nodes
}
