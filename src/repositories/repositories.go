package repositories

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/repositories/authors"
	"github.com/eduardothsantos/go-blog/src/repositories/posts"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	PostRepository   interfaces.PostRepository
	AuthorRepository interfaces.AuthorRepository
}

func NewRepositoryContainer(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		PostRepository:   posts.NewPostRepository(db),
		AuthorRepository: authors.NewAuthorRepository(db),
	}
}
