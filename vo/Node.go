package vo

import "fmt"

// https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Node struct {
	HostName      string
	Url           string
	Status        string //healthy, ..
	DocumentCount int    // read from curr_items or curr_items_tot

	Services []string // stores all the services running on the node - "index", "kv","n1ql"
	GetHits  int      // Number of get hits on the node

	DiskUsedByData int   // couch_docs_actual_disk_size
	RAMUsed        int   // interestingstats.mem_used
	FreeRAM        int64 // mem
	TotalRAM       int64 //memoryTotal

	//active - Nodes are participating in the cluster.
	//inactiveFailed - meaning that the node has failed, and administrator intervention is needed. Critical event.
	ClusterMembership string // Returns state of the node in the cluster
	CacheMisses       int    //ep_bg_fetched

	//If any of the below value shows constrants for any node of the cluster, address it and evaluate if additional nodes are required!
	//https://developer.couchbase.com/documentation/server/4.5/monitoring/monitoring-rest.html
	CPUUtilizationRate float64
	SwapUsed           int
	FreeMemory         int
}

func (this Node) GetHeaders() []string {
	return []string{"URL", "Services", "Status", "# Document", "# Hits", "Cluster Membership", "Free RAM", "Total RAM", "RAM used", "Disk Used By Data", "Cache Misses"}
}

// Returns the string representation of the Node as an array
// Sprintf does NOT print; it only evaluates and generates string
func (this Node) ToString() []string {
	return []string{
		fmt.Sprintf("%s", this.HostName),
		fmt.Sprintf("%v", this.Services),
		fmt.Sprintf("%s", this.Status),
		fmt.Sprintf("%d", this.DocumentCount),
		fmt.Sprintf("%d", this.GetHits),
		fmt.Sprintf("%s", this.ClusterMembership),
		fmt.Sprintf("%d", this.FreeRAM),
		fmt.Sprintf("%d", this.TotalRAM),
		fmt.Sprintf("%d", this.RAMUsed),
		fmt.Sprintf("%d", this.DiskUsedByData),
		fmt.Sprintf("%d", this.CacheMisses)}
}

//Active if 'clusterMemebrship' === 'active'
//Failed over if 'clusterMemebrship' === 'inactiveFailed'
//Down if 'status' !== 'healthy'
// and I think finding Pending nodes are more tricky and it is not enough to find nodes with 'clusterMembership' !== 'active'
