package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m node
var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Prints important details/metadata about the buckets stored in couchbase cluster",
	Long: `Command which prints the bucket related details 
            Command format
			$./c2m bucket --ip={val}    # for a specific node
            `,
	Run: bucket,
}

// Node command
// http://<ip>:8091/pools/default/buckets gives an insight into computing resources consumed per node
func bucket(cmd *cobra.Command, args []string) {
	fmt.Println(os.Getenv("URI") + " " + os.Getenv("USER") + " " + os.Getenv("PASS"))
	uri := os.Getenv("URI") + "/pools/default/buckets"
	uri = "http://www.mocky.io/v2/598aa61d410000d51d8211bf"

	contents := GetContent(uri, os.Getenv("USER"), os.Getenv("PASS"))

	var obj vo.BucketResp
	json.Unmarshal(contents, &obj)

	buckets := getBucketDetails(obj)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(buckets[0].GetHeaders())

	for _, val := range buckets {
		table.Append(val.ToString())
	}
	table.Render()

}

// Parse REST response and return the list of nodes struct
func getBucketDetails(resp vo.BucketResp) []vo.Bucket {
	len := len(resp) // number of buckets

	var buckets []vo.Bucket
	for i := 0; i < len; i++ {
		buckt := vo.Bucket{
			Name:           resp[i].Name,
			BucketType:     resp[i].BucketType,
			ReplicaNumber:  resp[i].ReplicaNumber,
			Ops:            resp[i].BasicStats.OpsPerSec,
			DiskFetches:    resp[i].BasicStats.DiskFetches,
			ItemCount:      resp[i].BasicStats.ItemCount,
			EvictionPolicy: resp[i].EvictionPolicy}
		buckets = append(buckets, buckt)
	}
	return buckets
}
