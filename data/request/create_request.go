package request

import "github.com/google/uuid"

type CreateRequest struct {
	Id          uuid.UUID
	Title       string
	Description string
	Author      string
	Genres      string
}
