package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sweetnebulae/go_ent/model"
)

type PostRepository interface {
	Save(ctx context.Context, post model.Post)
	FindAll(ctx context.Context) ([]model.Post, error)
	FindById(ctx context.Context, Id uuid.UUID) (model.Post, error)
	Update(ctx context.Context, post model.Post)
	Delete(ctx context.Context, Id uuid.UUID)
}
