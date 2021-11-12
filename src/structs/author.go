package structs

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
