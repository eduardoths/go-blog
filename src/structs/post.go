package structs

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Text      string         `json:"text"`
	AuthorId  int            `json:"author_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
