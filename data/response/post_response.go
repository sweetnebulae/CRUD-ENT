package response

import (
	"github.com/google/uuid"
	"time"
)

type PostResponse struct {
	Id          uuid.UUID `json:"id"`
	Tittle      string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time
}
