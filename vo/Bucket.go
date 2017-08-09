package vo

import "fmt"

type Bucket struct {
	Name           string // Name of the bucket, if any
	BucketType     string //Once a memcached or couchbase bucket has been created, its type cannot be changed.
	Ops            int    //Operations per second
	DiskFetches    int    // Disk fetches  ; //Indicates how frequently Couchbase Server is reaching to disk to retrieve information instead of using the information stored in RAM.
	ItemCount      int
	ReplicaNumber  int
	EvictionPolicy string // "valueOnly"
}

//Get bucket stat for a node
// pools/default/buckets/{bucket_name}/nodes/ip:8091/stats     //or %3A in place of :
// above api can even tell what all documents id are there at the node

func (this Bucket) GetHeaders() []string {
	return []string{"Name", "Type", "# Item", "# Replicas", "# Disk Fetches", "OPS"}
}

// Returns the string representation of the bucket as an array
func (this Bucket) ToString() []string {
	return []string{
		fmt.Sprintf("%s", this.Name),
		fmt.Sprintf("%v", this.BucketType),
		fmt.Sprintf("%d", this.ItemCount),
		fmt.Sprintf("%d", this.ReplicaNumber),
		fmt.Sprintf("%d", this.DiskFetches),
		fmt.Sprintf("%d", this.Ops)}
}
