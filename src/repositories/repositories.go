package repositories

import (
	"database/sql"

	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/repositories/authors"
	"github.com/eduardothsantos/go-blog/src/repositories/posts"
)

type RepositoryContainer struct {
	PostRepository interfaces.PostRepository
	AuthorRepository interfaces.AuthorRepository
}

func NewRepositoryContainer(db *sql.DB) RepositoryContainer {
	return RepositoryContainer{
		PostRepository: posts.NewPostRepository(db),
		AuthorRepository: authors.NewAuthorRepository(db),
	}
}