package handlers

import (
	"github.com/eduardothsantos/go-blog/src/handlers/authors"
	"github.com/eduardothsantos/go-blog/src/handlers/posts"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/services"
	"github.com/gofiber/fiber/v2"
)

type HandlerContainer struct {
	PostHandler interfaces.PostHandler
	AuthorHandler interfaces.AuthorHandler
}

func NewHandlerContainer(route fiber.Router, servs services.ServiceContainer) HandlerContainer {
	return HandlerContainer{
		AuthorHandler: authors.NewAuthorHandler(route, servs.AuthorService),
		PostHandler: posts.NewPostHandler(route, servs.PostService),
	}
}