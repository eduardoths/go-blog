package structs

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int
	Title     string
	Text      string
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
