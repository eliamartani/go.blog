package model

// Tag is an entity representation from database
type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
