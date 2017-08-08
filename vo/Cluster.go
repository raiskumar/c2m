package vo

//https://developer.couchbase.com/documentation/server/4.5/admin/ui-intro.html
type Cluster struct {
	name    string   // Name of the cluster, if any
	buckets []string // stores names of all buckets
	nodes   []string // stores ip address of all nodes

	//RAM details
	totalRAM                          int64 // total RAM or memory configured for all servers within the cluster
	totalRAMToDataBuckets             int64 // total RAM configured for data buckets
	totalRAMNotAllocatedToDataBuckets int64
	unusedRAM                         int64 // available for storing data
	inUseMemory                       int64 //amount of memory across all buckets that is in use

	// Disk overview
	totalClusterStorage int64 //The total amount of disk storage available across your entire cluster for storing data
}
