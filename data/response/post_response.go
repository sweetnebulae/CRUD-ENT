package response

import (
	"github.com/google/uuid"
	"time"
)

// PostResponse represents the response structure for a single book record.
//
// It contains metadata such as ID, title, author, description, genre, and creation timestamp.
type PostResponse struct {
	Id          uuid.UUID `json:"id"`          // Unique identifier for the book
	Title       string    `json:"title"`       // Title of the book
	Author      string    `json:"author"`      // Author of the book
	Description string    `json:"description"` // Description or summary of the book
	Genre       string    `json:"genre"`       // Genre category of the book
	CreatedAt   time.Time `json:"created_at"`  // Timestamp of when the book entry was created
}
