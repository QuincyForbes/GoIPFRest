package models

type Metadata struct {
	Cid         string `json:"cid"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Cid struct {
	Cid string `json:"cid"`
}
