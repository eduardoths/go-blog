package integration

import (
	"io/ioutil"
	"net/http"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/handlers"
	"github.com/eduardothsantos/go-blog/src/repositories"
	"github.com/eduardothsantos/go-blog/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ServerInterface interface {
	ClearDatabase()
	Test(req *http.Request) (string, int)
	Insert(v interface{}) (tx *gorm.DB)
	Get(table string, id int, v interface{})
}

type ServerContainer struct {
	server *fiber.App
	db     *gorm.DB
}

func (sc ServerContainer) ClearDatabase() {
	sc.db.Exec("DELETE FROM posts; DELETE FROM authors;")
}

func (sc ServerContainer) Test(req *http.Request) (string, int) {
	resp, _ := sc.server.Test(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), resp.StatusCode
}

func (sc ServerContainer) Insert(v interface{}) (tx *gorm.DB) {
	tx = sc.db.Save(v)
	return tx
}

func (sc ServerContainer) Get(table string, id int, v interface{}) {
	sc.db.Table(table).Where("id = ?", id).Take(&v)
}

func InitTestServer() ServerInterface {
	server := fiber.New()
	db := databases.TestConfig()
	repositoriesContainer := repositories.NewRepositoryContainer(db)
	servicesContainer := services.GetServices(repositoriesContainer)
	handlersContainer := handlers.NewHandlerContainer(server, servicesContainer)
	handlersContainer.AuthorHandler.Route()
	handlersContainer.PostHandler.Route()

	sc := ServerContainer{
		server: server,
		db:     db,
	}

	return sc
}
