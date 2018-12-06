package model

// Posts is a representation from query to list all posts
type Posts struct {
	CategoryTitle string `json:"categorytitle"`
	CategoryURL   string `json:"categoryurl"`
	Tags          string `json:"tags"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Description   string `json:"description"`
	Content       string `json:"content"`
	ImageURL      string `json:"imageurl"`
	DatePublished string `json:"datepublished"`
	Author        string `json:"author"`
}
