package vo

import (
	"fmt"

	"github.com/raiskumar/c2m/common"
)

//https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Cluster struct {
	Name        string            // Name of the cluster, if any
	NodesStatus map[string]string // maps node's ip address to the status {active, failed over, down, pending rebalance}

	DiskUsedByData int64 //hdd.usedByData
	FreeDisk       int64 // hdd.free
	TotalDisk      int64 //hdd.total

	RAMUsedByData int64 //ram.usedByData : total RAM configured for data buckets
	TotalRAM      int64 // ram.total     : total RAM or memory configured for all servers within the cluster
	UsedRAM       int64 // ram.used
}

//Active Servers	The number of active servers within the current cluster configuration.
//Servers Failed Over	The number of servers that have failed over due to an issue that should be investigated.
//Servers Down	The number of servers that are down and cannot be contacted.
//Servers Pending Rebalance	The number of servers that are currently waiting to be rebalanced after joining a cluster or being reactivated after failover.

func (this Cluster) GetHeaders() []string {
	return []string{"Name", "Node -> Membership Status", "Total RAM", "Data RAM", "Used RAM", "Total Disk", "Data Disk", "Free Disk"}
}

func (this Cluster) ToString() []string {
	var status string
	for key, value := range this.NodesStatus {
		status = status + key + "->" + value + "\n"
	}
	return []string{
		fmt.Sprintf("%s", this.Name),
		fmt.Sprintf("%v", status),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.TotalRAM)),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.RAMUsedByData)),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.UsedRAM)),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.TotalDisk)),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.DiskUsedByData)),
		fmt.Sprintf("%s", common.HumanRedableMemory(this.FreeDisk))}
}
