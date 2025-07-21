package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id          uuid.UUID
	Title       string
	Description string
	Author      string
	Genre       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
