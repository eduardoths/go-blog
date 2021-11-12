package posts

import (
	"testing"
	"time"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/structs"
	"gorm.io/gorm"
)

var db *gorm.DB = databases.TestConfig()
var author structs.Author = structs.Author{
	Name:  "Test author",
	Email: "test@author.com",
}
var authorId int = 5
var post structs.Post = structs.Post{
	Title: "Test post",
	Text:  "Test text",
}
var postRepo PostRepository = NewPostRepository(db)

func queryPost(postId int) structs.Post {
	var actualPost structs.Post
	db.Where("id = ?", postId).Take(&actualPost)
	return actualPost
}

func insertAuthor() int {
	db.Save(&author)
	return author.ID

}

func insertPost(authorId int) int {
	newPost := structs.Post{
		Title:    "Test post",
		Text:     "Test text",
		AuthorId: authorId,
	}
	db.Save(&newPost)
	return newPost.ID
}

func cleanTimestamp(post *structs.Post) {
	post.DeletedAt = gorm.DeletedAt{}
	post.CreatedAt = time.Time{}
	post.UpdatedAt = time.Time{}
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
	authorId := insertAuthor()

	t.Run("Create post on database", func(t *testing.T) {
		expectedValue := structs.Post{
			Title:    "Test post",
			Text:     "This is a test post",
			AuthorId: authorId,
		}
		id, err := postRepo.Create(expectedValue)
		expectedValue.ID = id
		actualPost := queryPost(id)
		cleanTimestamp(&actualPost)

		tests.AssertEquals(t, nil, err)
		tests.AssertEquals(t, expectedValue, actualPost)
	})
	t.Run("Create post - SQL Injection", func(t *testing.T) {
		expectedValue := structs.Post{
			Title:    "Test post",
			Text:     "This is a test post;DELETE FROM posts; DELETE FROM authors;",
			AuthorId: authorId,
		}
		id, err := postRepo.Create(expectedValue)
		expectedValue.ID = id
		actualPost := queryPost(id)
		cleanTimestamp(&actualPost)

		tests.AssertEquals(t, nil, err)
		tests.AssertEquals(t, expectedValue, actualPost)
	})
}

func TestGet(t *testing.T) {
	t.Run("Retrieve post from database", func(t *testing.T) {
		var expectedErr error = nil
		authorId = insertAuthor()
		postId := insertPost(authorId)
		expectedPost := structs.Post{
			ID:       postId,
			Text:     post.Text,
			Title:    post.Title,
			AuthorId: authorId,
		}
		aPost, err := postRepo.Get(postId)
		cleanTimestamp(&aPost)

		tests.AssertEquals(t, expectedErr, err)
		tests.AssertEquals(t, expectedPost, aPost)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update post on database", func(t *testing.T) {
		authorId := insertAuthor()
		postId := insertPost(authorId)
		expectedPost := structs.Post{
			ID:       postId,
			Text:     "Test post 2",
			Title:    "Updated post",
			AuthorId: authorId,
		}
		var expectedErr error = nil
		actualErr := postRepo.Update(postId, expectedPost)
		actualPost := queryPost(postId)
		cleanTimestamp(&actualPost)

		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedPost, actualPost)
	})

	t.Run("Update post - SQL Injection", func(t *testing.T) {
		authorId := insertAuthor()
		postId := insertPost(authorId)
		expectedPost := structs.Post{
			ID:       postId,
			Text:     "Test post 2",
			Title:    "Updated post;DELETE FROM posts; DELETE FROM authors;",
			AuthorId: authorId,
		}
		var expectedErr error = nil
		actualErr := postRepo.Update(postId, expectedPost)
		actualPost := queryPost(postId)
		cleanTimestamp(&actualPost)

		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedPost, actualPost)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete post", func(t *testing.T) {
		authorId := insertAuthor()
		postId := insertPost(authorId)
		var expectedErr error = nil
		actualErr := postRepo.Delete(postId)
		actualPost := queryPost(postId)
		cleanTimestamp(&actualPost)

		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, structs.Post{}, actualPost)
	})
}
