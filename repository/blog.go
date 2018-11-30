package repository

import (
	"database/sql"

	model "../model"
)

// RepoBlog defines repository access to Blog
// Inherits properties from Repo
type RepoBlog struct {
	Repo
}

// NewRepoBlog returns a representation from RepoBlog
func NewRepoBlog() RepoBlog {
	var repoBlog RepoBlog

	repoBlog.getQuery = "SELECT ID, CategoryID, Title, Url, Description, ImageUrl, Content, cast(DateCreation as char) DateCreation, ifnull(cast(DatePublished as char), cast(DateCreation as char)) DatePublished, Author, Active from post WHERE Url = ?"
	repoBlog.listQuery = `SELECT ID, CategoryID, Title, Url, Description, ImageUrl, cast(DateCreation as char) DateCreation, ifnull(cast(DatePublished as char), cast(DateCreation as char)) DatePublished, Author, Active from post
	where Active = 1 and (DatePublished <= now() or DatePublished is null)
	ORDER BY IfNull(DatePublished, DateCreation), ID
	LIMIT ? OFFSET ?`
	repoBlog.insertQuery = "INSERT INTO post (ID, CategoryID, Title, Url, Description, ImageUrl, Content, Author, DateCreation, DatePublished, Active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	repoBlog.updateQuery = "UPDATE post SET CategoryID = IFNULL(?, CategoryID), Title = IFNULL(?, Title), Url = IFNULL(?, Url), Description = IFNULL(?, Description), ImageUrl = IFNULL(?, ImageUrl), Content = IFNULL(?, Content), Author = IFNULL(?, Author), DateCreation = IFNULL(?, DateCreation), DatePublished = IFNULL(?, DatePublished), Active = IFNULL(?, Active) WHERE ID = ?"
	repoBlog.deleteQuery = "DELETE FROM post WHERE ID = ?"

	return repoBlog
}

// GetByURL returns a model representation obtained from database
func (r RepoBlog) GetByURL(url string) (model.Post, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	var model model.Post

	// Read data from register
	err := db.QueryRow(r.getQuery, url).Scan(&model.ID, &model.CategoryID, &model.Title, &model.URL, &model.Description, &model.ImageURL, &model.Content, &model.DateCreation, &model.DatePublished, &model.Author, &model.Active)

	// Return post model and error
	return model, err
}

// List returns a paged list of model obtained from database
func (r RepoBlog) List(index int, length int) ([]model.Post, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	var list []model.Post

	// Retrieve objects from database
	results, err := db.Query(r.listQuery, length, index)

	// Check if there's something wrong with the query
	if err != nil {
		return list, err
	}

	for results.Next() {
		var model model.Post

		err := results.Scan(&model.ID, &model.CategoryID, &model.Title, &model.URL, &model.Description, &model.ImageURL, &model.DateCreation, &model.DatePublished, &model.Author, &model.Active)

		// Check if there's something wrong with scanning the row
		if err != nil {
			return list, err
		}

		// Add it if it's ok
		list = append(list, model)
	}

	// Returns list of objects obtained from database
	return list, nil
}

// Insert inserts a register into database
func (r RepoBlog) Insert(model model.Post) (sql.Result, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	// Prepare the query to be executed
	command, err := db.Prepare(r.insertQuery)

	// Check if there's something wrong with the query
	if err != nil {
		return nil, err
	}

	// Finally executes the command
	return command.Exec(model.CategoryID, model.Title, model.URL, model.Description, model.ImageURL, model.Content, model.DateCreation, model.DatePublished, model.Author, model.Active)
}

// Update updates a register from database
func (r RepoBlog) Update(model model.Post) (sql.Result, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	// Prepare the query to be executed
	command, err := db.Prepare(r.updateQuery)

	// Check if there's something wrong with the query
	if err != nil {
		return nil, err
	}

	// Finally executes the command
	return command.Exec(model.CategoryID, model.Title, model.URL, model.Description, model.ImageURL, model.Content, model.DateCreation, model.DatePublished, model.Author, model.Active, model.ID)
}

// Delete removes a register from database
func (r RepoBlog) Delete(id int) (sql.Result, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	// Prepare the query to be executed
	command, err := db.Prepare(r.deleteQuery)

	// Check if there's something wrong with the query
	if err != nil {
		return nil, err
	}

	// Finally executes the command
	return command.Exec(id)
}
