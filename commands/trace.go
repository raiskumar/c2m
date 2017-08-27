package commands

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/raiskumar/c2m/common"
	"github.com/raiskumar/c2m/vo"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Traces a document details with the document id",
	Long: `Command which traces all details of the document with it's id
 Command format
$./c2m trace {bucket_name} {doc_id}
`,
	Run: trace,
}

// trace command
// http://<ip>:8091/pools/default/buckets gives bucket and node details
func trace(cmd *cobra.Command, args []string) {
	common.ValidateCommand(NodeURL)
	var bucketName, docId string
	if len(args) == 2 {
		bucketName = args[0]
		docId = args[1]
	} else {
		fmt.Println(" Sub-command missing !")
		fmt.Println(" $c2m trace {bucket_name} {doc_id}")
		os.Exit(1)
	}
	uri := NodeURL + "/pools/default/buckets"
	uri = "http://www.mocky.io/v2/598aa61d410000d51d8211bf" // Test URL

	hash, vBucketNum := getVBucketNumber(docId)
	fmt.Println(bucketName, uri, vBucketNum)

	contents := common.GetRestContent(uri, UserID, Password)

	var obj vo.BucketResp
	json.Unmarshal(contents, &obj)

	v := getTraceDetails(obj, bucketName, docId, hash, vBucketNum)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(v.GetHeaders())
	table.Append([]string{"Bucket Name", fmt.Sprint(v.BucketName)})
	table.Append([]string{"Document Key", v.DocumentID})
	table.Append([]string{"(CRC) Hash of Key", fmt.Sprint(v.Hash)})
	table.Append([]string{"vBucket which owns Document", fmt.Sprint(v.VBucketNumber)})
	table.Append([]string{"Primary Data Node", fmt.Sprint(v.PrimaryDataNode)})
	table.Append([]string{"Replica Data Nodes", v.ReplicaDataNodes})
	//table.Append([]string{"Document Value", v.DocumentVal})

	table.Render()
}

var crcTable *crc32.Table
var castagnoliTable = crc32.MakeTable(crc32.Castagnoli) // see http://golang.org/pkg/hash/crc32/#pkg-constants
func getVBucketNumber(dockID string) (uint32, uint32) {
	checksum := crc32.ChecksumIEEE([]byte(dockID))
	fmt.Println(" checksum = ", checksum)

	crc := crc32.New(castagnoliTable)
	crc.Write([]byte(dockID))
	modulo := crc.Sum32() % 1024

	fmt.Printf("Sum32 : %x \n", crc.Sum32())
	fmt.Println(" # bucket#", modulo)

	crc32Table := crc32.MakeTable(0xD5828281)
	crc32Int := crc32.Checksum([]byte(dockID), crc32Table)
	modulo1 := checksum % 1024
	fmt.Println("CRC32 String is", crc32Int, " checksum bucket#", modulo1)

	return crc.Sum32(), modulo // test with modulo1 as well
}

//https://forums.couchbase.com/t/connect-to-a-specific-server-node/7928/2

// Parse REST response and return the list of nodes struct
func getTraceDetails(resp vo.BucketResp, bucketName, docId string, hash, vBucketNum uint32) vo.Trace {
	len := len(resp) // number of buckets
	var response vo.Trace
	for i := 0; i < len; i++ {
		if resp[i].Name == bucketName {
			vBucketServerMap := resp[i].VBucketServerMap
			nodeLocator := vBucketServerMap.VBucketMap[vBucketNum]
			var primary, replicas string
			for i, val := range nodeLocator {
				if i == 0 {
					primary = vBucketServerMap.ServerList[val]
				} else {
					if val != -1 {
						replicas = replicas + "," + vBucketServerMap.ServerList[val]
					} else {
						replicas = vBucketServerMap.ServerList[0]
					}
				}
			}
			response = vo.Trace{
				BucketName:       bucketName,
				DocumentID:       docId,
				Hash:             hash,
				DocumentVal:      "",
				VBucketNumber:    vBucketNum,
				PrimaryDataNode:  primary,
				ReplicaDataNodes: replicas}
		}
	}
	return response
}
