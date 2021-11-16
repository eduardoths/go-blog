package posts

import (
	"time"

	"github.com/eduardothsantos/go-blog/src/structs"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

type postUpdate struct {
	Title     string
	Text      string
	UpdatedAt time.Time
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (pr PostRepository) Create(post structs.Post) (int, error) {
	tx := pr.db.Save(&post)
	return post.ID, tx.Error
}

func (pr PostRepository) Get(id int) (structs.Post, error) {
	var post structs.Post
	tx := pr.db.Where("id = ?", id).Take(&post)
	return post, tx.Error
}
func (pr PostRepository) Update(id int, post structs.Post) error {
	postToUpdate := postUpdate{
		Title: post.Title,
		Text:  post.Text,
	}
	tx := pr.db.Model(structs.Post{}).Where("id = ?", id).Updates(postToUpdate)
	return tx.Error
}

func (pr PostRepository) Delete(id int) error {
	tx := pr.db.Where("id = ?", id).Delete(&structs.Post{})
	return tx.Error
}
