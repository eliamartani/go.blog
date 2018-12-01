package api

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/eliamartani/go.blog/model"
	repository "github.com/eliamartani/go.blog/repository"
	"github.com/gorilla/mux"
)

var repoBlog = repository.NewRepoBlog()

// GetBlog is the main endpoint
func GetBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	defer fmt.Println("[INFO]", "Closing current connection...")

	// get variables
	vars := mux.Vars(r)

	url := vars["url"]

	post, err := repoBlog.GetByURL(url)

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	// returns json with Response representation
	OKDataResponse(w, "", post)
}

// ListBlog retrieve all blog posts
func ListBlog(w http.ResponseWriter, r *http.Request) {
	processListBlog(w, r, 10, 1)
}

// ListPagedBlog retrieve all blog posts within an interval
func ListPagedBlog(w http.ResponseWriter, r *http.Request) {
	// get variables
	vars := mux.Vars(r)

	length, err := strconv.Atoi(vars["length"])

	//check if length is valid
	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	index, err := strconv.Atoi(vars["index"])

	// check if index is valid
	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	processListBlog(w, r, length, index)
}

// processListBlog process both request, with or without parameter
func processListBlog(w http.ResponseWriter, r *http.Request, length int, index int) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	defer fmt.Println("[INFO]", "Closing current connection...")

	posts, err := repoBlog.List(index, length)

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	if posts == nil {
		NoDataFoundResponse(w, "")
		return
	}

	// returns json with Response representation
	OKDataResponse(w, "", posts)
}

// InsertBlog inserts a register into database
func InsertBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	defer fmt.Println("[INFO]", "Closing current connection...")

	// Get values from Form
	id, err := strconv.Atoi(r.FormValue("id"))
	categoryID, err := strconv.Atoi(r.FormValue("categoryid"))
	title := r.FormValue("title")
	url := r.FormValue("url")
	description := r.FormValue("description")
	imageURL := r.FormValue("imageurl")
	content := r.FormValue("content")
	dateCreation := r.FormValue("datecreation")
	datePublished := r.FormValue("datepublished")
	author := r.FormValue("author")
	active := byte(0)

	if r.FormValue("active") == "1" {
		active = byte(1)
	}

	// Fill the model with data received through POST
	var model = model.Post{
		ID:            id,
		CategoryID:    categoryID,
		Title:         title,
		URL:           url,
		Description:   description,
		ImageURL:      imageURL,
		Content:       content,
		DateCreation:  dateCreation,
		DatePublished: datePublished,
		Author:        author,
		Active:        active,
	}

	result, err := repoBlog.Insert(model)

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	rowsCount, err := result.RowsAffected()

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	// Returns json with Response representation
	if rowsCount > 0 {
		OKResponse(w, "Register inserted successfully")
	} else {
		OKResponse(w, "No rows were affected")
	}
}

// UpdateBlog updates a register from database
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	defer fmt.Println("[INFO]", "Closing current connection...")

	// Get values from Form
	id, err := strconv.Atoi(r.FormValue("id"))
	categoryID, err := strconv.Atoi(r.FormValue("categoryid"))
	title := r.FormValue("title")
	url := r.FormValue("url")
	description := r.FormValue("description")
	imageURL := r.FormValue("imageurl")
	content := r.FormValue("content")
	dateCreation := r.FormValue("datecreation")
	datePublished := r.FormValue("datepublished")
	author := r.FormValue("author")
	active := byte(0)

	if r.FormValue("active") == "1" {
		active = byte(1)
	}

	// Fill the model with data received through POST
	var model = model.Post{
		ID:            id,
		CategoryID:    categoryID,
		Title:         title,
		URL:           url,
		Description:   description,
		ImageURL:      imageURL,
		Content:       content,
		DateCreation:  dateCreation,
		DatePublished: datePublished,
		Author:        author,
		Active:        active,
	}

	result, err := repoBlog.Update(model)

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	rowsCount, err := result.RowsAffected()

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	// Returns json with Response representation
	if rowsCount > 0 {
		OKResponse(w, "Register updated successfully")
	} else {
		OKResponse(w, "No rows were affected")
	}
}

// DeleteBlog removes a register from database
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	defer fmt.Println("[INFO]", "Closing current connection...")

	id, err := strconv.Atoi(r.FormValue("id"))

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	result, err := repoBlog.Delete(id)

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	rowsCount, err := result.RowsAffected()

	if HasError(err) {
		ServerErrorResponse(w, "")
		return
	}

	// Returns json with Response representation
	if rowsCount > 0 {
		OKResponse(w, "Register deleted successfully")
	} else {
		OKResponse(w, "No rows were affected")
	}
}
