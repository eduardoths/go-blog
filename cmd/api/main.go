package main

import (
	"log"
	"net/http"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/handlers"
	"github.com/eduardothsantos/go-blog/src/httputils"
	"github.com/eduardothsantos/go-blog/src/repositories"
	"github.com/eduardothsantos/go-blog/src/services"
	"github.com/gofiber/fiber/v2"
)

func health(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	response.Data = "Health ok!"
	return ctx.Status(http.StatusOK).JSON(response)
}

func main() {
	server := fiber.New()
	db := databases.Config()
	repositoriesContainer := repositories.NewRepositoryContainer(db)
	servicesContainer := services.GetServices(repositoriesContainer)
	handlersContainer := handlers.NewHandlerContainer(server, servicesContainer)
	handlersContainer.AuthorHandler.Route()
	handlersContainer.PostHandler.Route()

	server.Get("/health", health)

	if err := server.Listen(":3000"); err != nil {
		log.Fatalf("Couldn't start application, err: %v", err.Error())
	}
}
