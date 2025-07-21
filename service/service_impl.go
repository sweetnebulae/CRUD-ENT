package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/data/request"
	"github.com/sweetnebulae/go_ent/data/response"
	"github.com/sweetnebulae/go_ent/helper"
	"github.com/sweetnebulae/go_ent/model"
	"github.com/sweetnebulae/go_ent/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{PostRepository: postRepository}
}

func (p PostServiceImpl) Create(ctx context.Context, request request.CreateRequest) {
	postData := model.Post{
		Title:       request.Title,
		Author:      request.Author,
		Genre:       request.Genres,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

func (p PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts, _ := p.PostRepository.FindAll(ctx)

	var result []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Tittle:      value.Title,
			Author:      value.Author,
			Genre:       value.Genre,
			Description: value.Description,
		}
		result = append(result, post)
	}
	return result
}

func (p PostServiceImpl) FindById(ctx context.Context, id uuid.UUID) response.PostResponse {
	postById, err := p.PostRepository.FindById(ctx, id)
	helper.ErrorPanic(err)

	post := response.PostResponse{
		Id:          postById.Id,
		Tittle:      postById.Title,
		Author:      postById.Author,
		Genre:       postById.Genre,
		Description: postById.Description,
		CreatedAt:   postById.CreatedAt,
	}
	return post
}

func (p PostServiceImpl) Update(ctx context.Context, updateRequest request.UpdateRequest) {
	postData := model.Post{
		Title:       updateRequest.Title,
		Author:      updateRequest.Author,
		Genre:       updateRequest.Genre,
		Description: updateRequest.Description,
	}
	p.PostRepository.Update(ctx, postData)
}

func (p PostServiceImpl) Delete(ctx context.Context, id uuid.UUID) {
	post, err := p.PostRepository.FindById(ctx, id)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}
