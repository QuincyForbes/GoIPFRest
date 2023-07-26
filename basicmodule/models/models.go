// Package models defines data structures used in the application.
package models

// Metadata represents the structure of metadata associated with an item.
type Metadata struct {
	Cid         string `json:"cid"`         // Unique Content ID associated with the metadata
	Image       string `json:"image"`       // URL or path to the associated image
	Description string `json:"description"` // Description of the item
	Name        string `json:"name"`        // Name of the item
}

// Cid represents a structure for a single Content ID.
type Cid struct {
	Cid string `json:"cid"` // Unique Content ID
}
