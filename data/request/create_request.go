// Package request contains request payload definitions for HTTP operations.
package request

// CreateRequest defines the structure of a request to create a new book post.
//
// It includes the title, author, genre, and description of the book.
type CreateRequest struct {
	Title       string // Title of the book
	Author      string // Author of the book
	Genres      string // Genre of the book (note: inconsistent field name, see UpdateRequest)
	Description string // Description or summary of the book
}
