package vo

import "fmt"

// https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Node struct {
	HostName      string
	Url           string
	Status        string //healthy, ..
	DocumentCount int    // read from curr_items or curr_items_tot
	MemoryFree    int64
	MemoryTotal   int64
	MemoryUsed    int    //This will tell you how much memory is being used by the bucket, which you can compare to the allocated memory
	Services      string // stores all the services running on the node - "index", "kv","n1ql"
	GetHits       int    // Number of get hits on the node
}

//Active Servers	The number of active servers within the current cluster configuration.
//Servers Failed Over	The number of servers that have failed over due to an issue that should be investigated.
//Servers Down	The number of servers that are down and cannot be contacted.
//Servers Pending Rebalance	The number of servers that are currently waiting to be rebalanced after joining a cluster or being reactivated after failover.

func (this Node) GetHeaders() []string {
	return []string{"URL", "Services", "Status", "# Document", "# Hits"}
}

// Returns the string representation of the Node as an array
// Sprintf does NOT print; it only evaluates and generates string
func (this Node) ToString() []string {
	return []string{
		fmt.Sprintf("%s", this.HostName), fmt.Sprintf("%s", this.Services),
		fmt.Sprintf("%s", this.Status), fmt.Sprintf("%d", this.DocumentCount),
		fmt.Sprintf("%d", this.GetHits)}
}
