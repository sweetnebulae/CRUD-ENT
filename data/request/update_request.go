package request

import (
	"github.com/google/uuid"
)

type UpdateRequest struct {
	Id          uuid.UUID
	Title       string
	Description string
	Author      string
	Genre       string
}
