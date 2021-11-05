package services

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/repositories"
	"github.com/eduardothsantos/go-blog/src/services/authors"
	"github.com/eduardothsantos/go-blog/src/services/posts"
)

type ServiceContainer struct {
	PostService interfaces.PostService
	AuthorService interfaces.AuthorService
}

func GetServices(repos repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		PostService: posts.NewPostService(repos.PostRepository),
		AuthorService: authors.NewAuthorService(repos.AuthorRepository),
	}
}