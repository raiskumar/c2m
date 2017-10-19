package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m node
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Gets important details/metadata of nodes of the cluster!",
	Long: `Command which prints the node related details 
Command format:
$./c2m node stats   # stats is optional
`,
	Run: node,
}

// Node command
// http://<ip>:8091/pools/default gives an insight into computing resources consumed per node
// http://<ip>:8091/nodeStatuses // returns status of nodes
func node(cmd *cobra.Command, args []string) {
	subCommand := "stats"
	if len(args) == 0 || args[0] != "stats" {
		subCommand = ""
	}
	common.ValidateCommand(NodeURL)
	uri := NodeURL + "/pools/default"
	//uri = "http://mocky.io/v2/599448371100004001723034" // Test URL

	contents := common.GetRestContent(uri, UserID, Password)
	var obj vo.PoolResp
	json.Unmarshal(contents, &obj)

	nodes := getAllNodes(obj)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(nodes[0].GetHeaders(subCommand))

	for _, val := range nodes {
		table.Append(val.ToString(subCommand))
	}
	table.Render()
	metaInfoNode()
}

// Parse REST response and return the list of nodes struct
func getAllNodes(resp vo.PoolResp) []vo.Node {
	len := len(resp.Nodes)
	var nodes []vo.Node
	var autoFailover string
	for i := 0; i < len; i++ {
		autoFailover = isAutofailoverEnabled(resp.Nodes[i].Hostname)
		n := vo.Node{
			HostName:           resp.Nodes[i].Hostname,
			Url:                resp.Nodes[i].CouchAPIBase,
			Status:             resp.Nodes[i].Status,
			Services:           resp.Nodes[i].Services,
			GetHits:            resp.Nodes[i].InterestingStats.GetHits,
			DocumentCount:      resp.Nodes[i].InterestingStats.CurrItems,
			DiskUsedByData:     int64(resp.Nodes[i].InterestingStats.CouchDocsActualDiskSize), // couch_docs_actual_disk_size
			RAMUsed:            int64(resp.Nodes[i].InterestingStats.MemUsed),                 // interestingstats.mem_used
			FreeRAM:            resp.Nodes[i].MemoryFree,
			TotalRAM:           resp.Nodes[i].MemoryTotal,
			ClusterMembership:  resp.Nodes[i].ClusterMembership,
			CacheMisses:        resp.Nodes[i].InterestingStats.EpBgFetched,
			CPUUtilizationRate: resp.Nodes[i].SystemStats.CPUUtilizationRate,
			SwapUsed:           resp.Nodes[i].SystemStats.SwapUsed,
			FreeMemory:         resp.Nodes[i].SystemStats.MemFree,
			Uptime:             resp.Nodes[i].Uptime,
			AutoFailover:       autoFailover}
		nodes = append(nodes, n)
	}
	return nodes
}

func isAutofailoverEnabled(host string) string {
	uri := host + "/settings/autoFailover"

	if !strings.Contains(uri, "http") {
		uri = "http://" + uri // if url doesn't start with http
	}

	//uri = "http://mocky.io/v2/5995589011000037107232e7" // Test URL
	contents := common.GetRestContent(uri, UserID, Password)
	//fmt.Println(uri + "-------" + string(contents))
	var obj vo.AutoFailovrResp
	json.Unmarshal(contents, &obj)
	if obj.Enabled == true {
		return "Y, " + strconv.Itoa(obj.Timeout) + " sec"
	}
	return "N"
}

func metaInfoNode() {
	fmt.Println("Note--")
	fmt.Println("* Cluster Membership: active -> Node is part of cluster and taking traffic!")
	fmt.Println("* Auto Failover: Promotion of replicas to master after waiting for given time. Cluster Manager detects and determines if a data node is unavailable and then initiate a hard failover. Hard failover is the ability to drop a unavailable or unstable node quickly from the cluster. Dropping a node is achieved by promoting replica vBuckets on the remaining cluster nodes to active.")
}
