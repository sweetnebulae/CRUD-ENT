package service

import (
	"context"
	"github.com/google/uuid"
	data "github.com/sweetnebulae/go_ent/data/request"
	"github.com/sweetnebulae/go_ent/data/response"
)

type PostService interface {
	Create(ctx context.Context, request data.CreateRequest)
	FindAll(ctx context.Context, Id uuid.UUID) []response.Response
	FindById(ctx context.Context, Id uuid.UUID) response.PostResponse
	Update(ctx context.Context, request data.UpdateRequest)
	Delete(ctx context.Context, Id uuid.UUID)
}
