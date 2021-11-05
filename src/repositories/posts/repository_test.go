package posts

import (
	"database/sql"
	"testing"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/structs"
)

var db *sql.DB = databases.TestConfig()
var author structs.Author = structs.Author {
	Name: "Test author",
	Email: "test@author.com",
}
var authorId int = 5
var post structs.Post = structs.Post {
	Title: "Test post",
	Text: "Test text",
}
var postRepo PostRepository = NewPostRepository(db)

func queryPost(t testing.TB, postId int) (structs.Post, int, error) {
	t.Helper()
	var actualPost structs.Post
	var actualAuthorId int
	query := `
		SELECT 
			title, 
			text, 
			authors.id 
		FROM posts 
		INNER JOIN authors ON authors.id = posts.author_id
		WHERE posts.id = $1;`
	err := db.QueryRow(query, postId).Scan(
			&actualPost.Title, 
			&actualPost.Text, 
			&actualAuthorId) 
	return actualPost, actualAuthorId, err
}

func insertAuthor(t testing.TB) int {
	t.Helper()
	var authorId int
	err := db.QueryRow("INSERT INTO authors (name, email) VALUES ($1, $2) RETURNING id;", author.Name, author.Email).Scan(&authorId)
	if err != nil {
		t.Errorf("Test failed to setup database, error %v", err.Error())
	}
	return authorId

}

func insertPost(t testing.TB, authorId int) int {
	t.Helper()
	var postId int
	err := db.QueryRow("INSERT INTO posts (title, text, author_id) VALUES ($1, $2, $3) RETURNING id;", post.Title, post.Text, authorId).Scan(&postId)
	if err != nil {
		t.Errorf("Test failed to setup database %v", err.Error())
	}
	return postId
}

func TestNewPostRepository(t *testing.T) {
	t.Run("Create post repository container", func(t *testing.T) {
		expected := PostRepository{db: nil}
		actual := NewPostRepository(nil)

		if expected != actual {
			t.Errorf(tests.AssertFailed, expected, actual)
		}
	})
}


func TestCreate(t *testing.T) {
	defer db.Exec("DELETE FROM posts; DELETE FROM authors;")
	authorId := insertAuthor(t)

	t.Run("Create post on database", func(t *testing.T) {
		err := postRepo.Create(post, authorId)
		tests.AssertEquals(t, nil, err)
		actualPost, actualAuthorId, err := queryPost(t, 1)
		if err != nil {
			t.Errorf("Failed to retrieve data from database, err %v", err)
		}
		tests.AssertEquals(t, post, actualPost)
		tests.AssertEquals(t, authorId, actualAuthorId)
	})
}

func TestGet(t *testing.T) {
	t.Run("Retrieve post from database", func(t *testing.T) {
		defer db.Exec("DELETE FROM posts; DELETE FROM authors;")
		authorId = insertAuthor(t)
		postId := insertPost(t, authorId)
		expectedPost := structs.Post{
			Text: post.Text,
			Title: post.Title,
			Author: author,
		}
		var expectedErr error = nil
		actualPost, err := postRepo.Get(postId)
		tests.AssertEquals(t, expectedErr, err)
		tests.AssertEquals(t, expectedPost, actualPost)
	})
}

func TestUpdate(t *testing.T) {
	defer db.Exec("DELETE FROM posts; DELETE FROM authors;")
	t.Run("Update post on database", func(t *testing.T) {
		authorId := insertAuthor(t)
		postId := insertPost(t, authorId)
		expectedPost := structs.Post{
			Text: "Test post 2",
			Title: "Updated post",
		}
		var expectedErr error = nil 
		actualErr := postRepo.Update(postId, expectedPost)
		tests.AssertEquals(t, expectedErr, actualErr)
		actualPost, _, err := queryPost(t, postId)
		if err != nil {
			t.Errorf("Couldn't retrieve data from database, error %v", err.Error())
		}
		tests.AssertEquals(t, expectedPost, actualPost)
	})
}

func TestDelete(t *testing.T) {
	defer db.Exec("DELETE FROM posts; DELETE FROM authors;")
	authorId := insertAuthor(t)
	postId := insertPost(t, authorId)
	var expectedErr error = nil
	t.Run("Delete post", func(t *testing.T) {
		actualErr := postRepo.Delete(postId)
		tests.AssertEquals(t, expectedErr, actualErr)
		actualPost, _, _ := queryPost(t, postId)
		tests.AssertEquals(t, structs.Post{}, actualPost)
	})
}