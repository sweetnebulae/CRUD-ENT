package model

import (
	"github.com/google/uuid"
	"time"
)

// Post represents the domain model for a book or article record in the system.
// It contains metadata such as title, author, genre, and timestamps for tracking creation and updates.
type Post struct {
	// Id is the unique identifier for the post, generated using UUID.
	Id uuid.UUID

	// Title is the title of the post or book.
	Title string

	// Description provides a brief overview or synopsis of the content.
	Description string

	// Author denotes the person who wrote or created the post.
	Author string

	// Genre categorizes the type or style of the content (e.g., fiction, thriller).
	Genre string

	// CreatedAt stores the timestamp when the post was first created.
	CreatedAt time.Time

	// UpdatedAt stores the timestamp of the most recent update to the post.
	UpdatedAt time.Time
}
