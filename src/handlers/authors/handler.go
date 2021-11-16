package authors

import (
	"net/http"
	"strconv"

	"github.com/eduardothsantos/go-blog/src/httputils"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
	"github.com/gofiber/fiber/v2"
)

func recoverPanic(ctx *fiber.Ctx) {
	if r := recover(); r != nil {
		response := httputils.NewResponse()
		response.Errors = []interface{}{"INTERNAL SERVER ERROR"}
		ctx.Status(http.StatusInternalServerError).JSON(response)
	}
}

type AuthorHandler struct {
	serv  interfaces.AuthorService
	route fiber.Router
}

func NewAuthorHandler(route fiber.Router, serv interfaces.AuthorService) AuthorHandler {
	return AuthorHandler{
		serv:  serv,
		route: route,
	}
}

func (ah AuthorHandler) Route() {
	grouter := ah.route.Group("/authors")
	grouter.Post("/", ah.create)
	grouter.Get("/:id", ah.get)
	grouter.Put("/:id", ah.update)
	grouter.Delete("/:id", ah.delete)
}

func (ah AuthorHandler) create(ctx *fiber.Ctx) error {
	var authorCreate structs.Author
	var body map[string]interface{}
	response := httputils.NewResponse()
	var status int

	if err := ctx.BodyParser(&body); err != nil {
		response.Errors = []interface{}{"Invalid request", err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	defer recoverPanic(ctx)
	authorCreate = structs.Author{
		Name:  body["name"].(string),
		Email: body["email"].(string),
	}
	id, err := ah.serv.Create(authorCreate)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{
			err.Error(),
		}
		return ctx.Status(status).JSON(response)
	}
	response.Data = map[string]int{"id": id}
	status = http.StatusCreated
	return ctx.Status(status).JSON(response)
}

func (ah AuthorHandler) get(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	var status int
	authorId := ctx.Params("id")
	id, err := strconv.Atoi(authorId)
	if err != nil || authorId == "" {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad query params"}
		return ctx.Status(status).JSON(response)

	}
	data, err := ah.serv.Get(id)
	if err != nil {
		status = http.StatusNotFound
		response.Errors = []interface{}{err.Error()}
		return ctx.Status(status).JSON(response)
	}

	status = http.StatusOK
	response.Data = data
	return ctx.Status(status).JSON(response)
}

func (ah AuthorHandler) update(ctx *fiber.Ctx) error {
	var authorUpdate structs.Author
	var body map[string]interface{}
	response := httputils.NewResponse()
	var status int

	authorId := ctx.Params("id")

	id, err := strconv.Atoi(authorId)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad Request"}
		return ctx.Status(status).JSON(response)
	}

	if err := ctx.BodyParser(&body); err != nil {
		response.Errors = []interface{}{"Invalid request", err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	defer recoverPanic(ctx)
	authorUpdate = structs.Author{
		Name:  body["name"].(string),
		Email: body["email"].(string),
	}

	if err := ah.serv.Update(id, authorUpdate); err != nil {
		if err.Error() == "record not found" {
			status = http.StatusNotFound
			response.Errors = []interface{}{err.Error()}
		} else {
			status = http.StatusBadRequest
			response.Errors = []interface{}{
				"Bad body parameters",
				err.Error(),
			}
		}
		return ctx.Status(status).JSON(response)
	}

	response.Data = "Author changed!"
	status = http.StatusOK
	return ctx.Status(status).JSON(response)
}

func (ah AuthorHandler) delete(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	var status int
	authorId := ctx.Params("id")
	id, err := strconv.Atoi(authorId)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad Request"}
		return ctx.Status(status).JSON(response)
	}
	if err = ah.serv.Delete(id); err != nil {
		if err.Error() == "record not found" {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		response.Errors = []interface{}{err.Error()}
		return ctx.Status(status).JSON(response)
	}

	status = http.StatusOK
	response.Data = "User deleted"
	return ctx.Status(status).JSON(response)
}
