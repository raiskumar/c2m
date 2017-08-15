// Command - bucket
package commands

import (
	"encoding/json"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

// ./c2m node
var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Prints important details/metadata about the buckets stored in couchbase cluster",
	Long: `Command which prints the bucket related details 
            Command format
			$./c2m bucket {optional_bucket_name}
            `,
	Run: bucket,
}

// bucket command
// http://<ip>:8091/pools/default/buckets gives bucket and node details
func bucket(cmd *cobra.Command, args []string) {
	common.ValidateCommand(NodeURL)
	var bucketName string
	if len(args) > 0 {
		bucketName = args[0]
	}
	uri := NodeURL + "/pools/default/buckets"
	//uri = "http://www.mocky.io/v2/598aa61d410000d51d8211bf" // Test URL

	contents := common.GetRestContent(uri, UserID, Password)

	var obj vo.BucketResp
	json.Unmarshal(contents, &obj)

	buckets := getBucketDetails(obj)
	printBucketCommandOutput(buckets, bucketName)
}

// Parse REST response and return the list of nodes struct
func getBucketDetails(resp vo.BucketResp) []vo.Bucket {
	len := len(resp) // number of buckets

	var buckets []vo.Bucket
	for i := 0; i < len; i++ {
		vBucketMap := resp[i].VBucketServerMap.VBucketMap

		count := 0 // count of vBucket
		for range vBucketMap {
			count = count + 1
		}

		buckt := vo.Bucket{
			Name:           resp[i].Name,
			BucketType:     resp[i].BucketType,
			ReplicaNumber:  resp[i].ReplicaNumber,
			Ops:            resp[i].BasicStats.OpsPerSec,
			DiskFetches:    resp[i].BasicStats.DiskFetches,
			ItemCount:      resp[i].BasicStats.ItemCount,
			EvictionPolicy: resp[i].EvictionPolicy,
			NumVBuckets:    count}
		buckets = append(buckets, buckt)
	}
	return buckets
}

func printBucketCommandOutput(buckets []vo.Bucket, bucketName string) {
	//fmt.Println(" verbose =", Verbose)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(buckets[0].GetHeaders())

	for _, val := range buckets {
		if len(bucketName) == 0 {
			table.Append(val.ToString())
		} else if val.Name == bucketName {
			table.Append(val.ToString())
		}
	}
	table.Render()
}
