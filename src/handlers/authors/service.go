package authors

import (
	"net/http"
	"strconv"

	"github.com/eduardothsantos/go-blog/src/domain/myhttp"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	serv interfaces.AuthorService
	route fiber.Router
}

func NewAuthorHandler(route fiber.Router, serv interfaces.AuthorService) AuthorHandler{
	return AuthorHandler{
		serv: serv,
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
	response := myhttp.New()
	status := 0
	
	if err := ctx.BodyParser(&body); err != nil {
		response.Errors = []interface{}{"Invalid request", err.Error()}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	authorCreate = structs.Author{
		Name: body["name"].(string),
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
	response := myhttp.New()
	status := http.StatusOK
	authorId := ctx.Params("id")
	id, err := strconv.Atoi(authorId)
	if err != nil || authorId == "" {
		status = http.StatusBadRequest
		response.Errors =[]interface{}{"Bad query params"}

	} else if data, err := ah.serv.Get(id); err != nil {
		if err.Error()[:3] == "pq:" {  // Errors coming from postgres
			status = http.StatusNotFound
			response.Errors = []interface{}{"Author not found"}
		} else {
			status = http.StatusInternalServerError
			response.Errors = []interface{}{"INTERNAL SERVER ERROR"}
		}
	} else {
		status = http.StatusOK
		response.Data = data
	}
	return ctx.Status(status).JSON(response)
}

func (ah AuthorHandler) update(ctx *fiber.Ctx) error {
	var authorUpdate structs.Author
	response := myhttp.New()
	status := http.StatusOK
	authorId := ctx.Params("id")
	id, err := strconv.Atoi(authorId)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad Request"}
	} else if err := ctx.BodyParser(&authorUpdate); err != nil {
		response.Errors = []interface{}{"Invalid request"}
		status = http.StatusBadRequest

	} else if err := ah.serv.Update(id, authorUpdate); err != nil {
		if err.Error()[:3] == "pq:" {  // Errors coming from postgres
			status = http.StatusInternalServerError
			response.Errors = []interface{}{"Internal server error"}
		} else {
			status = http.StatusBadRequest
			response.Errors = []interface{}{
				"Bad body parameters",
				err.Error(),
			}
		}

	} else {
		response.Data = "Author changed!"
		status = http.StatusOK
	}
	return ctx.Status(status).JSON(response)
}

func (ah AuthorHandler) delete(ctx *fiber.Ctx) error {
	response := myhttp.New()
	status := http.StatusOK
	authorId := ctx.Params("id")
	id, err := strconv.Atoi(authorId)
	if err != nil {
		status = http.StatusBadRequest
		response.Errors = []interface{}{"Bad Request"}
	} else if err = ah.serv.Delete(id); err != nil {
		if err.Error()[:3] == "pq:" {
			status = http.StatusInternalServerError
			response.Errors = []interface{}{"User not found"}
		} else {
			status = http.StatusInternalServerError
			response.Errors = []interface{}{err.Error()}
		}
	} else {
		status = http.StatusOK
		response.Data = "User deleted"
	}
	return ctx.Status(status).JSON(response)
}