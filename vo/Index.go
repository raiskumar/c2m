package vo

import "fmt"

type Index struct {
	Name        string
	Bucket      string
	Host        string // IP address of the host
	StorageMode string // can be forestdb
	Status      string // Ready
	Progress    int
}

func (this Index) GetHeaders() []string {
	return []string{"Name", "Bucket", "Host", "Storage Mode", "Status", "Progress"}
}

// Provides array of string of index details
func (this Index) ToString() []string {
	return []string{
		fmt.Sprintf("%s", this.Name),
		fmt.Sprintf("%s", this.Bucket),
		fmt.Sprintf("%s", this.Host),
		fmt.Sprintf("%s", this.StorageMode),
		fmt.Sprintf("%s", this.Status),
		fmt.Sprintf("%d", this.Progress)}
}
