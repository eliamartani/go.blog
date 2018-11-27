package model

// Post is an entity representation from database
type Post struct {
	ID            int    `json:"id"`
	CategoryID    int    `json:"categoryid"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Content       string `json:"content"`
	DateCreation  string `json:"datecreation"`
	DatePublished string `json:"datepublished"`
	Author        string `json:"author"`
	Active        byte   `json:"active"`
}
