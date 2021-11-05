package authors

import (
	"database/sql"
	"testing"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/structs"
)

var db *sql.DB = databases.TestConfig()
var authorRepo AuthorRepository = AuthorRepository{db: db}
var author structs.Author = structs.Author {
	Name: "Test author",
	Email: "test@author.com",
}

func QueryAuthor(t testing.TB, authorId int) (structs.Author, error) {
	t.Helper()
	var actualAuthor structs.Author
	err := db.QueryRow("SELECT name, email FROM authors WHERE id = $1;", authorId).Scan(&actualAuthor.Name, &actualAuthor.Email)
	return actualAuthor, err
}

func InsertAuthor(t testing.TB) {
	t.Helper()
	_, err := db.Exec("INSERT INTO authors (name, email) VALUES ($1, $2);", author.Name, author.Email)
	if err != nil {
		t.Errorf("Test failed to setup database")
	}
}

func TestNewAuthorRepository(t *testing.T) {
	t.Run("Create author repository container", func (t *testing.T) {
		expected := AuthorRepository{db: db}
		actual := NewAuthorRepository(db)

		tests.AssertEquals(t, expected, actual)
	})
}

func TestCreate(t *testing.T) {
	defer db.Exec("DELETE FROM authors;")
	testNum := 1
	t.Run("Create author", func (t *testing.T) {
		var expectedReturn error = nil 
		actualReturn := authorRepo.Create(author)
		tests.AssertEquals(t, expectedReturn, actualReturn)
	})

	t.Run("Author is created on database", func (t *testing.T) {
		actualAuthor, err := QueryAuthor(t, testNum)
		if err != nil {
			t.Errorf("Error querying database: %v", err.Error())
		}
		tests.AssertEquals(t, author, actualAuthor)
	})
}

func TestGet(t *testing.T) {
	defer db.Exec("DELETE FROM authors;")
	InsertAuthor(t)
	t.Run("Get author", func(t *testing.T) {
		actualAuthor, err := authorRepo.Get(2)
		tests.AssertEquals(t, author, actualAuthor)
		tests.AssertEquals(t, nil, err)
	})

}

func TestUpdate(t *testing.T) {
	defer db.Exec("DELETE FROM authors;")
	InsertAuthor(t)
	var testNum int = 3
	var expectedErr error = nil
	expectedAuthor := structs.Author{
		Name: "Test Author Two",
		Email: "test2@author.com",
	}
	t.Run("Update author", func(t *testing.T) {
		actualErr := authorRepo.Update(testNum, expectedAuthor) 
		tests.AssertEquals(t, expectedErr, actualErr)
	})
	t.Run("Update reflected to database", func(t *testing.T) {
		actualAuthor, err := QueryAuthor(t, testNum)
		if err != nil {
			t.Errorf("Error querying database: %v", err.Error())
		}
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
}

func TestDelete(t *testing.T) {
	defer db.Exec("DELETE FROM authors;")
	testNum := 4
	InsertAuthor(t)
	var expectedErr error = nil
	t.Run("Delete author", func(t *testing.T) {
		actualErr := authorRepo.Delete(testNum)
		tests.AssertEquals(t, expectedErr, actualErr)
	})
	t.Run("Delete reflected to database", func(t *testing.T) {
		actualAuthor, _ := QueryAuthor(t, testNum)
		tests.AssertEquals(t, structs.Author{}, actualAuthor)
	})
}