package vo

import "fmt"

// https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Node struct {
	HostName      string
	Url           string
	Status        string //healthy, ..
	DocumentCount int    // read from curr_items or curr_items_tot

	MemoryFree  int64
	MemoryTotal int64
	MemoryUsed  int      //This will tell you how much memory is being used by the bucket, which you can compare to the allocated memory
	Services    []string // stores all the services running on the node - "index", "kv","n1ql"
	GetHits     int      // Number of get hits on the node

	//active - Nodes are participating in the cluster.
	//inactiveFailed - meaning that the node has failed, and administrator intervention is needed. Critical event.
	ClusterMembership string // Returns state of the node in the cluster
}

func (this Node) GetHeaders() []string {
	return []string{"URL", "Services", "Status", "# Document", "# Hits", "Cluster Membership"}
}

// Returns the string representation of the Node as an array
// Sprintf does NOT print; it only evaluates and generates string
func (this Node) ToString() []string {
	return []string{
		fmt.Sprintf("%s", this.HostName), fmt.Sprintf("%v", this.Services),
		fmt.Sprintf("%s", this.Status), fmt.Sprintf("%d", this.DocumentCount),
		fmt.Sprintf("%d", this.GetHits), fmt.Sprintf("%s", this.ClusterMembership)}
}
