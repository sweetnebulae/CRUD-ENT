package request

import (
	"github.com/google/uuid"
)

// UpdateRequest defines the structure of a request to update an existing book post.
//
// It includes the ID of the book and updated information such as title, author, genre, and description.
type UpdateRequest struct {
	Id          uuid.UUID `json:"id"`          // UUID of the book to be updated
	Title       string    `json:"title"`       // Updated title of the book
	Author      string    `json:"author"`      // Updated author of the book
	Genre       string    `json:"genre"`       // Updated genre of the book
	Description string    `json:"description"` // Updated description of the book
}
