package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/vo"
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
	resp := parseResponse()
	nodesDetails := getAllNodesDetails(resp)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"URL", "Services", "Status", "# Document", "# Hits"})

	for _, val := range nodesDetails {
		table.Append(val.ToString())
	}
	table.Render()

}

func parseResponse() vo.PoolResp {
	url := "http://mocky.io/v2/5986c32d1100009c00fcbe4a"
	var poolResponse vo.PoolResp
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		json.Unmarshal(contents, &poolResponse)
	}
	return poolResponse
}

func getAllNodesDetails(resp vo.PoolResp) []vo.Node {
	fmt.Printf(" ---------------")
	fmt.Println(len(resp.Nodes))
	nodesLen := len(resp.Nodes)
	var nodes []vo.Node
	for i := 0; i < nodesLen; i++ {
		var servicesAsCsv string
		for _, v := range resp.Nodes[i].Services {
			servicesAsCsv = v + "," + servicesAsCsv
		}
		n := vo.Node{
			HostName:      resp.Nodes[i].Hostname,
			Url:           resp.Nodes[i].CouchAPIBase,
			Status:        resp.Nodes[i].Status,
			Services:      servicesAsCsv,
			GetHits:       resp.Nodes[i].InterestingStats.GetHits,
			DocumentCount: resp.Nodes[i].InterestingStats.CurrItems,
			MemoryTotal:   resp.Nodes[i].SystemStats.MemTotal,
			MemoryUsed:    resp.Nodes[i].InterestingStats.MemUsed,
			MemoryFree:    resp.Nodes[i].SystemStats.MemFree}
		nodes = append(nodes, n)
	}
	return nodes
}
