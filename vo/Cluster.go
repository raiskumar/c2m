package vo

import "fmt"

//https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Cluster struct {
	Name        string            // Name of the cluster, if any
	NodesStatus map[string]string // maps node's ip address to the status {active, failed over, down, pending rebalance}

	DiskUsedByData int   //hdd.usedByData
	FreeDisk       int64 // hdd.free
	TotalDisk      int64 //hdd.total

	RAMUsedByData  int   //ram.usedByData : total RAM configured for data buckets
	TotalRAM       int64 // ram.total     : total RAM or memory configured for all servers within the cluster
	UsedRAM        int64 // ram.used
	evictionPolicy string
}

//Active Servers	The number of active servers within the current cluster configuration.
//Servers Failed Over	The number of servers that have failed over due to an issue that should be investigated.
//Servers Down	The number of servers that are down and cannot be contacted.
//Servers Pending Rebalance	The number of servers that are currently waiting to be rebalanced after joining a cluster or being reactivated after failover.

func (this Cluster) GetHeaders() []string {
	return []string{"Name", "Node -> Membership Status", "RAM Total", "DATA RAM", "Used RAM", "Total Disk", "DATA DISK", "FREE DISK", "Eviction Policy"}
}

func (this Cluster) ToString() []string {
	var status string = ""
	for key, value := range this.NodesStatus {
		status = status + key + "->" + value + "\n"
	}
	return []string{
		fmt.Sprintf("%s", this.Name),
		fmt.Sprintf("%v", status),
		fmt.Sprintf("%d", this.TotalRAM),
		fmt.Sprintf("%d", this.RAMUsedByData),
		fmt.Sprintf("%d", this.UsedRAM),
		fmt.Sprintf("%d", this.TotalDisk),
		fmt.Sprintf("%d", this.DiskUsedByData),
		fmt.Sprintf("%d", this.FreeDisk),
		fmt.Sprintf("%s", this.evictionPolicy)}
}
