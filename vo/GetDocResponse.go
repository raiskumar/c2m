package vo

import "time"

type DocumentResponse struct {
	Meta struct {
		ID         string `json:"id"`
		Rev        string `json:"rev"`
		Expiration int    `json:"expiration"`
		Flags      int    `json:"flags"`
	} `json:"meta"`
	JSON struct {
		CreatedAt        time.Time `json:"createdAt"`
		SchemaVersion    int64     `json:"schemaVersion"`
		Creditors        []string  `json:"creditors"`
		CreatedDate      int64     `json:"createdDate"`
		CreatedBy        string    `json:"createdBy"`
		SchemeID         string    `json:"schemeId"`
		MaxDecimalPlaces int       `json:"maxDecimalPlaces"`
		Class            string    `json:"_class"`
		UpdatedDate      int64     `json:"updatedDate"`
		Type             string    `json:"type"`
		OwnerID          string    `json:"ownerId"`
	} `json:"json"`
}
