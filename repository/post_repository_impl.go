package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/ent"
	"github.com/sweetnebulae/go_ent/helper"
	"github.com/sweetnebulae/go_ent/model"
)

// PostRepositoryImpl is the Ent-based implementation of PostRepository.
type PostRepositoryImpl struct {
	Db *ent.Client // Ent ORM client
}

// NewPostRepositoryImpl creates and returns a new PostRepositoryImpl.
func NewPostRepositoryImpl(db *ent.Client) *PostRepositoryImpl {
	return &PostRepositoryImpl{Db: db}
}

// Save creates a new post record in the database using Ent.
func (p PostRepositoryImpl) Save(ctx context.Context, post model.Post) error {
	result, err := p.Db.Book.
		Create().
		SetTitle(post.Title).
		SetAuthor(post.Author).
		SetGenre(post.Genre).
		SetDescription(post.Description).
		Save(ctx)

	helper.ErrorPanic(err)
	fmt.Println(result)
	return nil
}

// FindAll retrieves all book records from the database and maps them to model.Post.
func (p PostRepositoryImpl) FindAll(ctx context.Context) ([]model.Post, error) {
	allPosts, err := p.Db.Book.Query().All(ctx)
	helper.ErrorPanic(err)

	var posts []model.Post
	for _, post := range allPosts {
		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Author:      post.Author,
			Genre:       post.Genre,
			Description: post.Description,
		}
		posts = append(posts, postData)
	}
	return posts, nil
}

// FindById retrieves a single post by its UUID.
func (p PostRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (model.Post, error) {
	post, err := p.Db.Book.Get(ctx, id)
	helper.ErrorPanic(err)

	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Author:      post.Author,
		Genre:       post.Genre,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
	}
	return postData, nil
}

// Update modifies an existing post identified by UUID.
// Returns error if the update operation fails.
func (p PostRepositoryImpl) Update(ctx context.Context, post model.Post) error {
	fmt.Printf("Attempting to update book with ID: %v\n", post.Id)
	fmt.Printf("Data to update: %+v\n", post)

	updateBook, err := p.Db.Book.
		UpdateOneID(post.Id).
		SetTitle(post.Title).
		SetAuthor(post.Author).
		SetDescription(post.Description).
		SetGenre(post.Genre).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	fmt.Printf("Updated book: %+v\n", updateBook)
	return nil
}

// Delete removes a post from the database by its UUID.
func (p PostRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := p.Db.Book.
		DeleteOneID(id).
		Exec(ctx)
	helper.ErrorPanic(err)
	return nil
}
