package vo

type Cluster struct {
	name            string // Name of the cluster, if any
	numberOfBuckets int
	bucketsName     string // CSV format
	numberOfNodes   int

	Status        string //healthy, ..
	DocumentCount int    // read from curr_items or curr_items_tot
	MemoryFree    int64
	MemoryTotal   int64
	MemoryUsed    int    //This will tell you how much memory is being used by the bucket, which you can compare to the allocated memory
	Services      string // stores all the services running on the node - "index", "kv","n1ql"
	GetHits       int    // Number of get hits on the node
}

/*
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
*/
