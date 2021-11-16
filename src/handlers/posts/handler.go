package posts

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

type PostHandler struct {
	serv  interfaces.PostService
	route fiber.Router
}

func NewPostHandler(route fiber.Router, serv interfaces.PostService) PostHandler {
	return PostHandler{
		serv:  serv,
		route: route,
	}
}

func (ph PostHandler) Route() {
	grouter := ph.route.Group("/posts")
	grouter.Post("/", ph.create)
	grouter.Get("/:id", ph.get)
	grouter.Put("/:id", ph.update)
	grouter.Delete("/:id", ph.delete)
}

func (ph PostHandler) create(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	var status int
	var body map[string]interface{}
	if err := ctx.BodyParser(&body); err != nil {
		response.Errors = []interface{}{"Invalid request", err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	defer recoverPanic(ctx)
	authorId := int(body["author_id"].(float64))
	postCreate := structs.Post{
		Title:    body["title"].(string),
		Text:     body["text"].(string),
		AuthorId: authorId,
	}
	id, err := ph.serv.Create(postCreate)
	if err != nil {
		status = http.StatusInternalServerError
		response.Errors = []interface{}{
			err.Error(),
		}
		return ctx.Status(status).JSON(response)
	}

	response.Data = map[string]int{"id": id}
	status = http.StatusCreated
	return ctx.Status(status).JSON(response)
}

func (ph PostHandler) get(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	var status int
	postId := ctx.Params("id")
	id, err := strconv.Atoi(postId)

	if err != nil || postId == "" {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad query params"}
	}
	data, err := ph.serv.Get(id)
	if err != nil {
		status = http.StatusNotFound
		response.Errors = []interface{}{err.Error()}
		return ctx.Status(status).JSON(response)
	}

	status = http.StatusOK
	response.Data = data
	return ctx.Status(status).JSON(response)
}

func (ph PostHandler) update(ctx *fiber.Ctx) error {
	var status int
	var body map[string]interface{}
	response := httputils.NewResponse()

	postId := ctx.Params("id")
	id, err := strconv.Atoi(postId)

	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad query params"}
	}

	if err := ctx.BodyParser(&body); err != nil {
		response.Errors = []interface{}{"Invalid request", err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	defer recoverPanic(ctx)
	postUpdate := structs.Post{
		Title: body["title"].(string),
		Text:  body["text"].(string),
	}

	if err := ph.serv.Update(id, postUpdate); err != nil {
		if err.Error() == "record not found" {
			status = http.StatusNotFound
			response.Errors = []interface{}{err.Error()}
		} else {
			status = http.StatusBadRequest
			response.Errors = []interface{}{
				"Internal Server Error",
			}
		}
		return ctx.Status(status).JSON(response)
	}
	response.Data = "Post changed!"
	status = http.StatusOK
	return ctx.Status(status).JSON(response)
}

func (ph PostHandler) delete(ctx *fiber.Ctx) error {
	response := httputils.NewResponse()
	var status int
	postId := ctx.Params("id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad query params"}
	}
	if err := ph.serv.Delete(id); err != nil {
		if err.Error() == "record not found" {
			status = http.StatusNotFound
			response.Errors = []interface{}{"Post not found"}
		} else {
			status = http.StatusInternalServerError
			response.Errors = []interface{}{"Internal Server Error"}
		}
		return ctx.Status(status).JSON(response)
	}
	status = http.StatusOK
	response.Data = "Post deleted!"
	return ctx.Status(status).JSON(response)
}
