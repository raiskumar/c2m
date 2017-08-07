package vo

import "fmt"

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

func (this Node) ToString() []string {
	/*fmt.Sprintf("Details for Node :%s", this.HostName)

	fmt.Printf("\n Details for Node :%s", this.HostName)
	fmt.Printf("\n URL :%20s", this.Url)
	fmt.Printf("\n Type :%20.20s", this.Services)
	fmt.Printf("\n Number of Documents :%20.20d", this.DocumentCount)
	fmt.Printf("\n Status :%20.20s", this.Status)
	fmt.Printf("\n # Hits :%20.20d", this.GetHits)*/

	data := []string{fmt.Sprintf("%s", this.HostName), fmt.Sprintf("%s", this.Services), fmt.Sprintf("%s", this.Status), fmt.Sprintf("%d", this.DocumentCount), fmt.Sprintf("%d", this.GetHits)}
	return data
}
