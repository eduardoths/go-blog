package posts

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

func cleanTimestamps(post *structs.Post) {
	post.CreatedAt = time.Time{}
	post.UpdatedAt = time.Time{}
	post.DeletedAt = gorm.DeletedAt{}
}

func createAuthor(t *testing.T) (structs.Author, int) {
	t.Helper()
	var author structs.Author
	author = structs.Author{
		Email: "test@author.com",
		Name:  "Test Author",
	}
	server.Insert(&author)
	return author, author.ID
}

func populateDatabase(t *testing.T, authorId int) (structs.Post, int) {
	t.Helper()
	post := structs.Post{
		Title:    "Test post",
		Text:     "Lorem ipsum",
		AuthorId: authorId,
	}
	server.Insert(&post)
	return post, post.ID
}

func getPost(t *testing.T, id int) structs.Post {
	t.Helper()
	var post structs.Post
	server.Get("posts", id, &post)
	return post
}

func TestCreatePost(t *testing.T) {
	t.Run("Create post", func(t *testing.T) {
		_, authorId := createAuthor(t)
		expectedPost := structs.Post{
			Title:    "Creating a Test post",
			Text:     "This is a test post",
			AuthorId: authorId,
		}
		req := requests.Post("/posts", expectedPost)
		expectedResponse := responses.StrResponse(map[string]interface{}{"id": 1}, nil)
		actualResponse, actualStatus := server.Test(req)

		actualPost := getPost(t, 1)
		expectedPost.ID = 1
		cleanTimestamps(&actualPost)

		tests.AssertEquals(t, http.StatusCreated, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedPost, actualPost)
	})

	t.Run("Invalid body post", func(t *testing.T) {
		req := requests.Post("/posts", nil)
		expectedResponse := responses.StrResponse(nil, []interface{}{"INTERNAL SERVER ERROR"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, http.StatusInternalServerError, actualStatus)
	})
	server.ClearDatabase()
}

func TestGetPost(t *testing.T) {
	t.Run("Get post information", func(t *testing.T) {
		_, authorId := createAuthor(t)
		data, id := populateDatabase(t, authorId)
		req := requests.Get(fmt.Sprintf("/posts/%d", id))
		expectedResponse := responses.StrResponse(data, nil)
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})

	t.Run("Getting fake post", func(t *testing.T) {
		var fakeId int = 987634
		req := requests.Get(fmt.Sprintf("/posts/%v", fakeId))
		expectedResponse := responses.StrResponse(nil, []interface{}{"record not found"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusNotFound, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})
	server.ClearDatabase()
}

func TestUpdatePost(t *testing.T) {
	t.Run("Updating post", func(t *testing.T) {
		_, authorId := createAuthor(t)
		_, id := populateDatabase(t, authorId)
		expectedPost := structs.Post{
			Title: "Updated title",
			Text:  "Updated text",
		}
		req := requests.Put(fmt.Sprintf("/posts/%d", id), expectedPost)

		expectedPost.AuthorId = authorId
		expectedPost.ID = id

		expectedResponse := responses.StrResponse("Post changed!", nil)
		actualResponse, actualStatus := server.Test(req)
		actualPost := getPost(t, id)
		cleanTimestamps(&actualPost)

		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedPost, actualPost)
	})

	t.Run("Trying to update without params", func(t *testing.T) {
		var id int = 0
		req := requests.Put(fmt.Sprintf("/posts/%v", id), nil)

		expectedResponse := responses.StrResponse(nil, []interface{}{"INTERNAL SERVER ERROR"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusInternalServerError, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})

	t.Run("Trying to update unexisting post", func(t *testing.T) {
		var fakeId int = 768945

		req := requests.Put(fmt.Sprintf("/posts/%v", fakeId), map[string]interface{}{"title": "title", "text": "text"})

		expectedResponse := responses.StrResponse(nil, []interface{}{"record not found"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusNotFound, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})

	server.ClearDatabase()
}

func TestDeletePost(t *testing.T) {
	t.Run("Deleting post", func(t *testing.T) {
		_, authorId := createAuthor(t)
		_, id := populateDatabase(t, authorId)
		req := requests.Delete(fmt.Sprintf("/posts/%d", id), nil)

		expectedResponse := responses.StrResponse("Post deleted!", nil)
		actualResponse, actualStatus := server.Test(req)
		actualPost := getPost(t, id)
		expectedPost := structs.Post{}

		tests.AssertEquals(t, http.StatusOK, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
		tests.AssertEquals(t, expectedPost, actualPost)

	})

	t.Run("Deleting not found post", func(t *testing.T) {
		var fakeId int = 678
		req := requests.Delete(fmt.Sprintf("/posts/%v", fakeId), nil)

		expectedResponse := responses.StrResponse(nil, []interface{}{"record not found"})
		actualResponse, actualStatus := server.Test(req)

		tests.AssertEquals(t, http.StatusNotFound, actualStatus)
		tests.AssertEquals(t, expectedResponse, actualResponse)
	})

	server.ClearDatabase()
}
