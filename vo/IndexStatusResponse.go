package vo

// Response for REST Call hppt://ip:port/indexStatus
type IndexResp struct {
	Indexes []struct {
		StorageMode string   `json:"storageMode"`
		Hosts       []string `json:"hosts"`
		Progress    int      `json:"progress"`
		Definition  string   `json:"definition"`
		Status      string   `json:"status"`
		Bucket      string   `json:"bucket"`
		Index       string   `json:"index"`
		ID          int64    `json:"id"`
	} `json:"indexes"`
	Version  int           `json:"version"`
	Warnings []interface{} `json:"warnings"`
}
