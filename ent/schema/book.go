package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Book defines the schema for the Book entity.
//
// Each book contains an ID, title, description, author, genre, and timestamps.
type Book struct {
	ent.Schema
}

// Fields defines the fields for the Book entity in the database.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		// Unique identifier for each book, using UUID.
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),

		// Title of the book. This field is required and cannot be empty.
		field.String("title").
			NotEmpty(),

		// Optional description or summary of the book.
		field.String("description").
			Optional(),

		// Author of the book.
		field.String("author"),

		// Genre of the book (e.g., fiction, history). This field is optional.
		field.String("genre").
			Optional(),

		// Timestamp when the book entry was created.
		field.Time("CreatedAt").
			Immutable().
			Default(time.Now),

		// Timestamp when the book entry was last updated.
		field.Time("UpdatedAt").
			Default(time.Now),
	}
}
