package commands

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m index
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Gets important details/metadata about indexes!",
	Long: `Command which prints the Index related details 
            Command format
			$./c2m index
            `,
	Run: index,
}

// index command
// http://<ip>:8091/indexStatus
func index(cmd *cobra.Command, args []string) {
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

// Parse REST response and return the list of Index struct
func getAllIndexes(resp vo.IndexResp) []vo.Index {
	len := len(resp.Indexes)
	var indexes []vo.Index
	for i := 0; i < len; i++ {
		in := vo.Index{
			Name:        resp.Indexes[i].Index,
			Bucket:      resp.Indexes[i].Bucket,
			Host:        strings.Join(resp.Indexes[i].Hosts[:], ","), // convert array into string
			StorageMode: resp.Indexes[i].StorageMode,
			Status:      resp.Indexes[i].Status,
			Progress:    resp.Indexes[i].Progress}
		indexes = append(indexes, in)
	}
	return indexes
}
