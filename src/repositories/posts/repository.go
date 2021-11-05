package posts

import (
	"database/sql"

	"github.com/eduardothsantos/go-blog/src/structs"
)

type PostRepository struct {
	db *sql.DB 
}

func NewPostRepository(db *sql.DB) (PostRepository) {
	return PostRepository{
		db:  db,
	}
}

func (pr PostRepository) Create(post structs.Post, authorId int) error {
	_, err := pr.db.Exec("INSERT INTO posts (title, text, author_id) VALUES ($1, $2, $3)",
	                  post.Title, 
					  post.Text, 
					  authorId)
	return err
}

func (pr PostRepository) Get(id int) (structs.Post, error) {
	query := `
		SELECT 
			title, 
			text,
			name, 
			email 
		FROM posts 
		INNER JOIN authors 
		ON posts.author_id = authors.id
		WHERE posts.id = $1`
	var post structs.Post
	err := pr.db.QueryRow(query, id).Scan(&post.Title, &post.Text, &post.Author.Name, &post.Author.Email)
	return post, err
}
func (pr PostRepository) Update(id int, post structs.Post) error {
	_, err := pr.db.Exec("UPDATE posts SET title=$1, text=$2 WHERE id=$3", post.Title, post.Text, id)
	return err
}

func (pr PostRepository) Delete(id int) error {
	_, err := pr.db.Exec("DELETE FROM posts WHERE id=$1", id)
	return err
}
