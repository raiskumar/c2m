package vo

import "fmt"

type Trace struct {
	BucketName       string
	DocumentID       string
	Hash             uint32 // CRC or Hash value
	DocumentVal      string
	VBucketNumber    uint32
	PrimaryDataNode  string
	ReplicaDataNodes string // csv
}

func (this Trace) GetHeaders() []string {
	return []string{"Key", "Value"}
}

func (this Trace) GetKeys() []string {
	return []string{"BucketName", "DocumentID", "Hash"}
}

// Returns the string representation of the bucket as an array
func (this Trace) ToString() []string {
	return []string{
		fmt.Sprintf("%s", "Bucket Name"),
		fmt.Sprintf("%s", this.BucketName)}
}
