package authors

import (
	"testing"
	"time"

	"github.com/eduardothsantos/go-blog/pkg/databases"
	"github.com/eduardothsantos/go-blog/src/domain/tests"
	"github.com/eduardothsantos/go-blog/src/structs"
	"gorm.io/gorm"
)

var db *gorm.DB = databases.TestConfig()
var authorRepo AuthorRepository = AuthorRepository{db: db}
var author structs.Author = structs.Author {
	Name: "Test author",
	Email: "test@author.com",
}

func queryAuthor(authorId int) structs.Author {
	var actualAuthor structs.Author
	db.Table("authors").Where("id = ?", authorId).Take(&actualAuthor)
	return actualAuthor
}

func insertAuthor() (int, structs.Author) {
	newAuthor := structs.Author{
		Name: "Test author",
		Email: "test@author.com",
	}
	db.Table("authors").Save(&newAuthor)
	return newAuthor.ID, newAuthor
}

func cleanTimestamp(author *structs.Author) {
	author.CreatedAt = time.Time{}
	author.UpdatedAt = time.Time{}
	author.DeletedAt = gorm.DeletedAt{}
}

func TestNewAuthorRepository(t *testing.T) {
	t.Run("Create author repository container", func (t *testing.T) {
		expected := AuthorRepository{db: db}
		actual := NewAuthorRepository(db)

		tests.AssertEquals(t, expected, actual)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create author function", func (t *testing.T) {
		var expectedReturn error = nil 
		expectedAuthor := author
		id, actualReturn := authorRepo.Create(expectedAuthor)
		expectedAuthor.ID = id
		actualAuthor := queryAuthor(id)
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedReturn, actualReturn)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
}

func TestGet(t *testing.T) {
	t.Run("Get author function", func(t *testing.T) {
		id, expectedAuthor := insertAuthor()
		actualAuthor, err := authorRepo.Get(id)
		cleanTimestamp(&expectedAuthor)
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
		tests.AssertEquals(t, nil, err)
	})
}

func TestUpdate(t *testing.T) {
	id, _ := insertAuthor()
	var expectedErr error = nil
	t.Run("Update author", func(t *testing.T) {
		expectedAuthor := structs.Author{
			Name: "Test Author Two",
			Email: "test2@author.com",
		}
		actualErr := authorRepo.Update(id, expectedAuthor) 
		actualAuthor := queryAuthor(id)
		actualAuthor.ID = 0  // Don't care
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete author function", func(t *testing.T) {
		id, _ := insertAuthor()
		var expectedErr error = nil
		actualErr := authorRepo.Delete(id)
		actualAuthor := queryAuthor(id)
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, structs.Author{}, actualAuthor)  // Should still exist as we're using soft deletes
	})
}