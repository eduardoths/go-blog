package structs

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
