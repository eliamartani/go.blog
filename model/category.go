package model

// Category is an entity representation from database
type Category struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Active byte   `json:"active"`
}
