package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ent.Schema
}

func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("title").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.String("author"),
		field.String("genre").
			NotEmpty(),
		field.Time("CreatedAt").
			Immutable().
			Default(time.Now),
		field.Time("UpdatedAt").
			Default(time.Now),
	}
}
