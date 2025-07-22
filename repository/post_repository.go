package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/model"
)

// PostRepository defines the interface for managing post data.
type PostRepository interface {
	// Save inserts a new post record into the database.
	Save(ctx context.Context, post model.Post) error

	// FindAll retrieves all post records from the database.
	FindAll(ctx context.Context) ([]model.Post, error)

	// FindById retrieves a post by its UUID.
	FindById(ctx context.Context, id uuid.UUID) (model.Post, error)

	// Update modifies an existing post identified by UUID.
	Update(ctx context.Context, post model.Post) error

	// Delete removes a post from the database by UUID.
	Delete(ctx context.Context, id uuid.UUID) error
}
