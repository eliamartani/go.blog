package api

import (
	"fmt"
	"net/http"
	"strconv"

	database "../database"
	"github.com/gorilla/mux"
)

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

// SQL Queries
const getQuery = "SELECT ID, CategoryID, Title, Url, Content, cast(DateCreation as char) DateCreation, ifnull(cast(DatePublished as char), '') DatePublished, Author, Active from post WHERE ID = ?"
const allQuery = "SELECT ID, CategoryID, Title, Url, cast(DateCreation as char) DateCreation, ifnull(cast(DatePublished as char), '') DatePublished, Author, Active from post"
const pagedQuery = `SELECT ID, CategoryID, Title, Url, cast(DateCreation as char) DateCreation, ifnull(cast(DatePublished as char), '') DatePublished, Author, Active from post
where Active = 1 and (DatePublished <= now() or DatePublished is null)
ORDER BY IfNull(DatePublished, DateCreation), ID
LIMIT ? OFFSET ?`
const insertQuery = "INSERT INTO post (ID, CategoryID, Title, Url, Content, Author, DateCreation, DatePublished, Active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

// GetBlog is the main endpoint
func GetBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	// close the connection at the end
	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// get variables
	vars := mux.Vars(r)

	key, err := strconv.Atoi(vars["id"])

	if HasError(err) {
		ResponseJSON(w, ServerError())
		return
	}

	// retrieve object from database
	var post Post

	err = db.QueryRow(getQuery, key).Scan(&post.ID, &post.CategoryID, &post.Title, &post.URL, &post.Content, &post.DateCreation, &post.DatePublished, &post.Author, &post.Active)

	if HasError(err) {
		ResponseJSON(w, NoDataFound())
		return
	}

	// returns json with Response representation
	ResponseJSON(w, ToResponse(post))
}

// ListBlog retrieve all blog posts
func ListBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// retrieve object from database
	results, err := db.Query(allQuery)

	if HasError(err) {
		ResponseJSON(w, ServerError())
		return
	}

	var posts []Post

	for results.Next() {
		var post Post

		err := results.Scan(&post.ID, &post.CategoryID, &post.Title, &post.URL, &post.DateCreation, &post.DatePublished, &post.Author, &post.Active)

		if HasError(err) {
			ResponseJSON(w, ServerError())
			return
		}

		posts = append(posts, post)
	}

	if posts == nil {
		ResponseJSON(w, NoDataFound())
		return
	}

	// returns json with Response representation
	ResponseJSON(w, ToResponse(posts))
}

// ListPagedBlog retrieve all blog posts within an interval
func ListPagedBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// get variables
	vars := mux.Vars(r)

	length, err := strconv.Atoi(vars["length"])

	if HasError(err) {
		ResponseJSON(w, ServerError())
		return
	}

	index, err := strconv.Atoi(vars["index"])

	if HasError(err) {
		ResponseJSON(w, ServerError())
		return
	}

	// retrieve object from database
	results, err := db.Query(pagedQuery, length, index)

	if HasError(err) {
		ResponseJSON(w, ServerError())
		return
	}

	var posts []Post

	for results.Next() {
		var post Post

		err := results.Scan(&post.ID, &post.CategoryID, &post.Title, &post.URL, &post.DateCreation, &post.DatePublished, &post.Author, &post.Active)

		if HasError(err) {
			ResponseJSON(w, ServerError())
			return
		}

		posts = append(posts, post)
	}

	if posts == nil {
		ResponseJSON(w, NoDataFound())
		return
	}

	// returns json with Response representation
	ResponseJSON(w, ToResponse(posts))
}

// InsertBlog inserts a register into database
func InsertBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// TODO
}

// UpdateBlog updates a register from database
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// TODO
}

// DeleteBlog removes a register from database
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// open connection
	db := database.Connect()

	defer db.Close()
	defer fmt.Println("[INFO]", "Closing current connection...")

	// TODO
}
