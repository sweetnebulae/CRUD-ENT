package service

import (
	"context"

	"github.com/google/uuid"
	data "github.com/sweetnebulae/go_ent/data/request"
	"github.com/sweetnebulae/go_ent/data/response"
)

// PostService defines the business logic layer interface
// for handling operations related to posts (books).
type PostService interface {
	// Create inserts a new post record based on the given request payload.
	//
	// Parameters:
	//   - ctx: Context for managing deadlines and cancellation.
	//   - request: CreateRequest struct containing the post data.
	//
	// Returns:
	//   - error if the operation fails.
	Create(ctx context.Context, request data.CreateRequest) error

	// FindAll retrieves all post records from the database.
	//
	// Parameters:
	//   - ctx: Context for managing deadlines and cancellation.
	//
	// Returns:
	//   - A slice of PostResponse and an error if any.
	FindAll(ctx context.Context) ([]response.PostResponse, error)

	// FindById retrieves a single post by its UUID.
	//
	// Parameters:
	//   - ctx: Context for managing deadlines and cancellation.
	//   - id: UUID of the post to be retrieved.
	//
	// Returns:
	//   - PostResponse containing the post data, and an error if any.
	FindById(ctx context.Context, id uuid.UUID) (response.PostResponse, error)

	// Update modifies an existing post record based on the provided update request.
	//
	// Parameters:
	//   - ctx: Context for managing deadlines and cancellation.
	//   - updateRequest: UpdateRequest struct with the updated post data.
	//
	// Returns:
	//   - error if the operation fails.
	Update(ctx context.Context, updateRequest data.UpdateRequest) error

	// Delete removes a post record by its UUID.
	//
	// Parameters:
	//   - ctx: Context for managing deadlines and cancellation.
	//   - id: UUID of the post to be deleted.
	//
	// Returns:
	//   - error if the operation fails.
	Delete(ctx context.Context, id uuid.UUID) error
}
