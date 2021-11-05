package authors

import (
	"testing"

	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
	"github.com/eduardothsantos/go-blog/src/tests/mocks/authors"
)

var fakeAuthor structs.Author = structs.Author{
	Name: "Test author",
	Email: "test@author.com",
}
var repo interfaces.AuthorRepository = authors.MockAuthorRepository{
	FakeAuthor: fakeAuthor,
}
var authorService AuthorService = AuthorService{Repo: repo}

func TestNewAuthorService(t *testing.T) {
	t.Run("Test creation of author service struct", func(t *testing.T) {
		repo := authors.MockAuthorRepository{}
		expectedResult := AuthorService{Repo: repo}
		actualResult := NewAuthorService(repo)
		tests.AssertEquals(t, expectedResult, actualResult)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test author creation service", func(t *testing.T) {
		var expectedErr error = nil
		actualErr := authorService.Create(fakeAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)  
	})

	t.Run("Test author creation with invalid name", func(t *testing.T) {
		var expectedErr string = "name.invalid"
		expectedAuthor := structs.Author{
			Name: "  ",
			Email: "test@author.com",
		}
		actualErr := authorService.Create(expectedAuthor)
		tests.AssertEquals(t, expectedErr, actualErr.Error())
	})
}

func TestGet(t *testing.T) {
	t.Run("Test author get service", func(t *testing.T) {
		var expectedErr error = nil
		expectedAuthor := fakeAuthor
		actualAuthor, actualErr := authorService.Get(0)
		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test author update service", func(t *testing.T) {
		var expectedErr error = nil
		actualErr := authorService.Update(0, fakeAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
	})

	t.Run("Test author update with invalid name", func(t *testing.T) {
		expectedErr := "name.invalid"
		expectedAuthor := structs.Author{
			Name: "     3ed",
			Email: "test@author.com",
		}
		actualErr := authorService.Update(0, expectedAuthor)
		tests.AssertEquals(t, expectedErr, actualErr.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete author update service", func(t *testing.T) {
		var expectedErr error = nil
		actualErr := authorService.Update(0, fakeAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
	})
}