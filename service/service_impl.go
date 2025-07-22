package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/data/request"
	"github.com/sweetnebulae/go_ent/data/response"
	"github.com/sweetnebulae/go_ent/model"
	"github.com/sweetnebulae/go_ent/repository"
)

// PostServiceImpl implements the PostService interface.
// It acts as the service layer that contains business logic
// and interacts with the PostRepository to persist data.
type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

// NewPostServiceImpl returns a new instance of PostServiceImpl.
// It requires a concrete implementation of PostRepository.
//
// Parameters:
//   - postRepository: a pointer to PostRepositoryImpl.
//
// Returns:
//   - *PostServiceImpl
func NewPostServiceImpl(postRepository *repository.PostRepositoryImpl) *PostServiceImpl {
	return &PostServiceImpl{PostRepository: postRepository}
}

// Create constructs a new Post from the request and delegates
// the persistence to the repository.
//
// Parameters:
//   - ctx: context for timeout and cancellation control.
//   - request: CreateRequest containing post data.
//
// Returns:
//   - error if creation fails.
func (p PostServiceImpl) Create(ctx context.Context, request request.CreateRequest) error {
	postData := model.Post{
		Title:       request.Title,
		Author:      request.Author,
		Genre:       request.Genres,
		Description: request.Description,
	}
	return p.PostRepository.Save(ctx, postData)
}

// FindAll fetches all posts from the repository and transforms
// them into response DTOs.
//
// Parameters:
//   - ctx: context for timeout and cancellation control.
//
// Returns:
//   - slice of PostResponse
//   - error if retrieval fails
func (p PostServiceImpl) FindAll(ctx context.Context) ([]response.PostResponse, error) {
	posts, _ := p.PostRepository.FindAll(ctx)

	var postResponse []response.PostResponse
	for _, post := range posts {
		post := response.PostResponse{
			Id:          post.Id,
			Title:       post.Title,
			Author:      post.Author,
			Genre:       post.Genre,
			Description: post.Description,
		}
		postResponse = append(postResponse, post)
	}
	return postResponse, nil
}

// FindById retrieves a post by its UUID and transforms
// it into a response DTO.
//
// Parameters:
//   - ctx: context for timeout and cancellation control.
//   - id: UUID of the post to be retrieved.
//
// Returns:
//   - PostResponse DTO
//   - error if not found or repository fails
func (p PostServiceImpl) FindById(ctx context.Context, id uuid.UUID) (response.PostResponse, error) {
	postById, err := p.PostRepository.FindById(ctx, id)
	if err != nil {
		return response.PostResponse{}, err
	}

	post := response.PostResponse{
		Id:          postById.Id,
		Title:       postById.Title,
		Author:      postById.Author,
		Genre:       postById.Genre,
		Description: postById.Description,
		CreatedAt:   postById.CreatedAt,
	}
	return post, nil
}

// Update modifies an existing post with new values.
//
// Parameters:
//   - ctx: context for timeout and cancellation control.
//   - updateRequest: UpdateRequest DTO with new post data.
//
// Returns:
//   - error if update fails or post not found
func (p PostServiceImpl) Update(ctx context.Context, updateRequest request.UpdateRequest) error {
	postData := model.Post{
		Id:          updateRequest.Id,
		Title:       updateRequest.Title,
		Author:      updateRequest.Author,
		Genre:       updateRequest.Genre,
		Description: updateRequest.Description,
	}
	return p.PostRepository.Update(ctx, postData)
}

// Delete removes a post from the database by its UUID.
//
// Parameters:
//   - ctx: context for timeout and cancellation control.
//   - id: UUID of the post to be deleted.
//
// Returns:
//   - error if post not found or deletion fails
func (p PostServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	post, err := p.PostRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	return p.PostRepository.Delete(ctx, post.Id)
}
