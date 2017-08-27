package commands

import (
	"encoding/json"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m index
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generates report of your Couchbase Cluster",
	Long: `Command which generates report 
Command format
$./c2m report
`,
	Run: report,
}

// report command
func report(cmd *cobra.Command, args []string) {
	common.ValidateCommand(NodeURL)
	uri := NodeURL + "/indexStatus"
	uri = "http://www.mocky.io/v2/599141c7120000060394645b" // Test URL

	restResponse := common.GetRestContent(uri, UserID, Password)
	var obj vo.IndexResp
	json.Unmarshal(restResponse, &obj)

	indexes := getAllIndexes(obj)

	if len(indexes) == 0 {
		return // No response!
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(indexes[0].GetHeaders())

	for _, val := range indexes {
		table.Append(val.ToString())
	}
	table.Render()

}
