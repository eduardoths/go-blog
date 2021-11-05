package posts

import (
	"net/http"

	"github.com/eduardothsantos/go-blog/src/domain/myhttp"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	serv interfaces.PostService
	route fiber.Router
}

func NewPostHandler(route fiber.Router, serv interfaces.PostService) PostHandler{
	return PostHandler{
		serv: serv,
		route: route,
	}
}

func (ph PostHandler) Route() {
	grouter := ph.route.Group("/posts")
	grouter.Post("/", create)
	grouter.Get("/:id", get)
	grouter.Put("/:id", update)
	grouter.Delete("/:id", delete)
}

func create(ctx *fiber.Ctx) error {
	response := myhttp.New()
	status := http.StatusOK
	return ctx.Status(status).JSON(response)
}

func get(ctx *fiber.Ctx) error {
	response := myhttp.New()
	status := http.StatusOK
	return ctx.Status(status).JSON(response)
}

func update(ctx *fiber.Ctx) error {
	response := myhttp.New()
	status := http.StatusOK
	return ctx.Status(status).JSON(response)
}

func delete(ctx *fiber.Ctx) error {
	response := myhttp.New()
	status := http.StatusOK
	return ctx.Status(status).JSON(response)
}