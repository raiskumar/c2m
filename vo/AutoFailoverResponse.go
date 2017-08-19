package vo

type AutoFailovrResp struct {
	Enabled bool `json:"enabled"`
	Timeout int  `json:"timeout"`
	Count   int  `json:"count"`
}
