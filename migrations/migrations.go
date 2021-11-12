package migrations

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string
	Email string
	Posts []Post `gorm:"foreignKey:AuthorID"`
}

type Post struct {
	gorm.Model
	Title string
	Text string
	AuthorID int
}