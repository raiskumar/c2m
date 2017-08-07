package vo

type Node struct {
	hostName      string
	url           string
	status        string //healthy, ..
	documentCount uint64
	memoryFree    uint64
	memoryTotal   uint64
	services      string // csv values like "index", "kv","n1ql"
}
