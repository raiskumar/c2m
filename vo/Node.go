package vo

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

func (this Node) GetHeader() string {
	return "Host Name" + "   URL" + "    Status" + "    Document Count" + "    Running Services"
}

func (this Node) ToString() string {
	//base 10
	return this.HostName + "   " + this.Url + "   " + this.Status + "   " + "   " + this.Services
}
