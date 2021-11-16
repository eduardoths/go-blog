package authors

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/structs"
	"github.com/eduardothsantos/go-blog/src/tests/integration"
	"github.com/eduardothsantos/go-blog/src/tests/utils/requests"
	"github.com/eduardothsantos/go-blog/src/tests/utils/responses"
	"gorm.io/gorm"
)

var server = integration.InitTestServer()

func cleanTimestamps(author *structs.Author) {
	author.CreatedAt = time.Time{}
	author.UpdatedAt = time.Time{}
	author.DeletedAt = gorm.DeletedAt{}
}

func populateDatabase(t *testing.T) (structs.Author, int) {
	t.Helper()
	var author structs.Author
	author = structs.Author{
		Email: "test@get.com",
		Name:  "test get",
	}
	server.Insert(&author)
	return author, author.ID
}

func getAuthor(t *testing.T, id int) structs.Author {
	t.Helper()
	var author structs.Author
	server.Get("authors", id, &author)
	return author
}

func TestCreateAuthor(t *testing.T) {
	t.Run("Create author", func(t *testing.T) {
		author := structs.Author{
			Name:  " Test\n     Author ",
			Email: " test@test.com ",
		}
		req := requests.Post("/authors", author)
		expectedResponse := responses.StrResponse(map[string]int{"id": 1}, nil)
		actualResponse, actualStatus := server.Test(req)

		actualAuthor := getAuthor(t, 1)
		expectedAuthor := structs.Author{
			Name:  "Test Author",
			Email: "test@test.com",
			ID:    1,
		}
		cleanTimestamps(&actualAuthor)

		tests.AssertEquals(t, http.StatusCreated, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})

	t.Run("Invalid body author post", func(t *testing.T) {
		req := requests.Post("/authors", nil)
		expectedResponse := responses.StrResponse(nil, []interface{}{"INTERNAL SERVER ERROR"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, http.StatusInternalServerError, actualStatus)
	})

	t.Run("Invalid name", func(t *testing.T) {
		req := requests.Post("/authors", structs.Author{
			Email: "test@valid.com",
			Name:  "3s   ",
		})
		expectedResponse := responses.StrResponse(nil, []interface{}{"name.invalid"})
		actualResponse, actualStatus := server.Test(req)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, http.StatusBadRequest, actualStatus)
	})
	t.Run("Invalid email", func(t *testing.T) {
		req := requests.Post("/authors", structs.Author{
			Email: "invalid",
			Name:  "Valid name",
		})
		expectedResponse := responses.StrResponse(nil, []interface{}{"email.invalid"})
		actualResponse, actualStatus := server.Test(req)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, http.StatusBadRequest, actualStatus)
	})
	server.ClearDatabase()
}

func TestGetAuthor(t *testing.T) {
	t.Run("Get author information", func(t *testing.T) {
		data, id := populateDatabase(t)
		req := requests.Get(fmt.Sprintf("/authors/%d", id))
		expectedResponse := responses.StrResponse(data, nil)
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})
	server.ClearDatabase()
}

func TestUpdateAuthor(t *testing.T) {
	t.Run("Updating author", func(t *testing.T) {
		_, id := populateDatabase(t)
		expectedAuthor := structs.Author{
			ID:    id,
			Email: "updated@email.com",
			Name:  "Updated Author Name",
		}
		req := requests.Put(fmt.Sprintf("/authors/%d", id), expectedAuthor)
		expectedResponse := responses.StrResponse("Author changed!", nil)
		actualResponse, actualStatus := server.Test(req)
		actualAuthor := getAuthor(t, id)
		cleanTimestamps(&actualAuthor)
		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})

	t.Run("Trying to update with invalid query params", func(t *testing.T) {
		_, id := populateDatabase(t)
		var fakeId int = 548968
		var fakeTimestamp = time.Date(2001, 12, 31, 23, 58, 58, 2, time.UTC)
		newAuthor := structs.Author{
			ID:        fakeId,
			Email:     "updated2@test.com",
			Name:      "Updated author name",
			UpdatedAt: fakeTimestamp,
			CreatedAt: fakeTimestamp,
		}

		req := requests.Put(fmt.Sprintf("/authors/%d", id), newAuthor)
		_, actualStatus := server.Test(req)
		actualAuthor := getAuthor(t, id)
		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, id, actualAuthor.ID)
		if actualAuthor.UpdatedAt == fakeTimestamp {
			t.Errorf("Input updating metadata at PUT /authors/%d", id)
		}
	})

	server.ClearDatabase()
}

func TestDeleteAuthor(t *testing.T) {
	t.Run("Deleting post", func(t *testing.T) {
		_, id := populateDatabase(t)
		req := requests.Delete(fmt.Sprintf("/authors/%d", id), nil)
		expectedResponse := responses.StrResponse("User deleted", nil)
		actualResponse, actualStatus := server.Test(req)

		actualAuthor := getAuthor(t, id)
		expectedAuthor := structs.Author{}

		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
	t.Run("Deleting unexisting post", func(t *testing.T) {
		var fakeId int = 456468468
		req := requests.Delete(fmt.Sprintf("/authors/%d", fakeId), nil)
		expectedResponse := responses.StrResponse(nil, []interface{}{"record not found"})
		actualResponse, actualStatus := server.Test(req)
		tests.AssertEquals(t, http.StatusNotFound, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})
}
