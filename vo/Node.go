package vo

import (
	"fmt"
	"strings"

	"github.com/raiskumar/c2m/common"
)

// https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Node struct {
	HostName      string
	Url           string
	Status        string //healthy, ..
	DocumentCount int    // read from curr_items or curr_items_tot

	Services []string // stores all the services running on the node - "index", "kv","n1ql"
	GetHits  int      // Number of get hits on the node

	DiskUsedByData int64 // couch_docs_actual_disk_size
	RAMUsed        int64 // interestingstats.mem_used
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
	FreeMemory         int64

	Uptime       string
	AutoFailover string
}

func (this Node) GetHeaders(subCommand string) []string {
	if subCommand == "stats" {
		return []string{"Host", "# Item", "CPU Utiliztion", "Total RAM", "RAM used", "Free RAM", "Disk Used By Data", "Up For"}
	}
	return []string{"Host", "Services Running", "Status", "Cluster Membership", "# Item", "# Hits", "Cache Misses", "Auto-Failover"}
}

// Returns the string representation of the Node as an array
// Sprintf does NOT print; it only evaluates and generates string
func (this Node) ToString(subCommand string) []string {
	if subCommand == "stats" {
		return []string{
			fmt.Sprintf("%s", this.HostName),
			fmt.Sprintf("%d", this.DocumentCount),
			fmt.Sprintf("%f", this.CPUUtilizationRate),
			fmt.Sprintf("%s", common.HumanRedableMemory(this.TotalRAM)),
			fmt.Sprintf("%s", common.HumanRedableMemory(this.RAMUsed)),
			fmt.Sprintf("%s", common.HumanRedableMemory(this.FreeRAM)),
			fmt.Sprintf("%s", common.HumanRedableMemory(this.DiskUsedByData)),
			fmt.Sprintf("%s", common.HumandRedableUpFor(this.Uptime))}
	}
	return []string{
		fmt.Sprintf("%s", this.HostName),
		fmt.Sprintf("%s", strings.Join(this.Services[:], ",")), // convert array into CSV
		fmt.Sprintf("%s", this.Status),
		fmt.Sprintf("%s", this.ClusterMembership),
		fmt.Sprintf("%d", this.DocumentCount),
		fmt.Sprintf("%d", this.GetHits),
		fmt.Sprintf("%d", this.CacheMisses),
		fmt.Sprintf("%s", this.AutoFailover)}
}

//Active if 'clusterMemebrship' === 'active'
//Failed over if 'clusterMemebrship' === 'inactiveFailed'
//Down if 'status' !== 'healthy'
// and I think finding Pending nodes are more tricky and it is not enough to find nodes with 'clusterMembership' !== 'active'
