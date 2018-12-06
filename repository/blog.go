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

	repoBlog.getQuery = `SELECT
			c.Title AS CategoryTitle,
			c.URL AS CategoryURL,
			GROUP_CONCAT(t.Title ORDER BY t.Title) AS Tags,
			p.Title,
			p.Url,
			p.ImageUrl,
			p.Description,
			p.Content,
			IFNULL(CAST(p.DatePublished AS CHAR), CAST(p.DateCreation AS CHAR)) DatePublished,
			p.Author
		FROM post p
		INNER JOIN category c ON c.ID = p.CategoryID
		INNER JOIN post_tag pt ON p.ID = pt.PostID
		INNER JOIN tag t ON t.ID = pt.TagID
		WHERE
			p.Url = ?`
	repoBlog.listQuery = `SELECT
			c.Title AS CategoryTitle,
			c.URL AS CategoryURL,
			posts.Tags,
			p.Title,
			p.Url,
			p.ImageUrl,
			p.Description,
			IFNULL(CAST(p.DatePublished AS CHAR), CAST(p.DateCreation AS CHAR)) DatePublished,
			p.Author
		FROM post p
		INNER JOIN (
			SELECT post.ID, GROUP_CONCAT(t.Title ORDER BY t.Title) AS Tags
			FROM post
			INNER JOIN post_tag pt ON post.ID = pt.PostID
			INNER JOIN tag t ON t.ID = pt.TagID
			WHERE
				post.Active = 1 AND (post.DatePublished <= NOW() OR post.DatePublished IS NULL)
			GROUP BY post.ID
		) posts ON posts.ID = p.ID
		INNER JOIN category c ON c.ID = p.CategoryID
		ORDER BY IFNULL(p.DatePublished, p.DateCreation), p.ID
		LIMIT ? OFFSET ?`
	repoBlog.insertQuery = "INSERT INTO post (ID, CategoryID, Title, Url, Description, ImageUrl, Content, Author, DateCreation, DatePublished, Active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	repoBlog.updateQuery = "UPDATE post SET CategoryID = IFNULL(?, CategoryID), Title = IFNULL(?, Title), Url = IFNULL(?, Url), Description = IFNULL(?, Description), ImageUrl = IFNULL(?, ImageUrl), Content = IFNULL(?, Content), Author = IFNULL(?, Author), DateCreation = IFNULL(?, DateCreation), DatePublished = IFNULL(?, DatePublished), Active = IFNULL(?, Active) WHERE ID = ?"
	repoBlog.deleteQuery = "DELETE FROM post WHERE ID = ?"

	return repoBlog
}

// GetByURL returns a model representation obtained from database
func (r RepoBlog) GetByURL(url string) (model.Posts, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	var model model.Posts

	// Read data from register
	err := db.QueryRow(r.getQuery, url).Scan(
		&model.CategoryTitle,
		&model.CategoryURL,
		&model.Tags,
		&model.Title,
		&model.URL,
		&model.ImageURL,
		&model.Description,
		&model.Content,
		&model.DatePublished,
		&model.Author)

	// Return post model and error
	return model, err
}

// List returns a paged list of model obtained from database
func (r RepoBlog) List(index int, length int) ([]model.Posts, error) {
	// Open connection
	db := r.Connect()

	// Close the connection at the end
	defer db.Close()

	var list []model.Posts

	// Retrieve objects from database
	results, err := db.Query(r.listQuery, length, (index-1)*length)

	// Check if there's something wrong with the query
	if err != nil {
		return list, err
	}

	for results.Next() {
		var model model.Posts

		err := results.Scan(
			&model.CategoryTitle,
			&model.CategoryURL,
			&model.Tags,
			&model.Title,
			&model.URL,
			&model.ImageURL,
			&model.Description,
			&model.DatePublished,
			&model.Author)

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
