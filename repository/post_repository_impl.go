package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/ent"
	"github.com/sweetnebulae/go_ent/helper"
	"github.com/sweetnebulae/go_ent/model"
)

type PostRepositoryImpl struct {
	Db *ent.Client
}

func (p PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	result, err := p.Db.Book.
		Create().
		SetTitle(post.Title).
		SetAuthor(post.Author).
		SetGenre(post.Genre).
		SetDescription(post.Description).
		Save(ctx)

	helper.ErrorPanic(err)
	fmt.Println(result)
}

func (p PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
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
	return posts
}

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

func (p PostRepositoryImpl) Update(ctx context.Context, post model.Post) error {
	updateBook, err := p.Db.Book.
		UpdateOneID(post.Id).
		SetTitle(post.Title).
		SetAuthor(post.Author).
		SetDescription(post.Description).
		SetGenre(post.Genre).
		Save(ctx)

	helper.ErrorPanic(err)
	fmt.Println(updateBook)
	return nil
}

func (p PostRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) {
	err := p.Db.Book.
		DeleteOneID(id).
		Exec(ctx)
	helper.ErrorPanic(err)
	return
}
