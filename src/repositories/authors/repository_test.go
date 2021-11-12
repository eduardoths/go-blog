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
var author structs.Author = structs.Author{
	Name:  "Test author",
	Email: "test@author.com",
}

func queryAuthor(authorId int) structs.Author {
	var actualAuthor structs.Author
	db.Table("authors").Where("id = ?", authorId).Take(&actualAuthor)
	return actualAuthor
}

func insertAuthor() (int, structs.Author) {
	newAuthor := structs.Author{
		Name:  "Test author",
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
	t.Run("Create author repository container", func(t *testing.T) {
		expected := AuthorRepository{db: db}
		actual := NewAuthorRepository(db)

		tests.AssertEquals(t, expected, actual)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create author function", func(t *testing.T) {
		var expectedReturn error = nil
		expectedAuthor := author
		id, actualReturn := authorRepo.Create(expectedAuthor)
		expectedAuthor.ID = id
		actualAuthor := queryAuthor(id)
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedReturn, actualReturn)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
	t.Run("Create author - sql injection", func(t *testing.T) {
		var expectedErr error = nil
		expectedAuthor := structs.Author{
			Name:  "author",
			Email: "author@email.com;DELETE FROM authors;",
		}
		id, actualErr := authorRepo.Create(expectedAuthor)
		expectedAuthor.ID = id
		actualAuthor := queryAuthor(id)
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
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
	t.Run("Update author", func(t *testing.T) {
		id, _ := insertAuthor()
		var expectedErr error = nil
		expectedAuthor := structs.Author{
			Name:  "Test Author Two",
			Email: "test2@author.com",
		}
		actualErr := authorRepo.Update(id, expectedAuthor)
		actualAuthor := queryAuthor(id)
		actualAuthor.ID = 0 // Don't care
		createdAt := actualAuthor.CreatedAt
		updatedAt := actualAuthor.UpdatedAt
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
		if !updatedAt.After(createdAt) {
			t.Errorf("Updated at wasn't changed")
		}
	})

	t.Run("Update author - SQL Injection", func(t *testing.T) {
		id, _ := insertAuthor()
		var expectedErr error = nil
		expectedAuthor := structs.Author{
			Name:  "Test author three",
			Email: "test@author.com;DELETE FROM posts; DELETE FROM authors;",
		}
		actualErr := authorRepo.Update(id, expectedAuthor)
		actualAuthor := queryAuthor(id)
		actualAuthor.ID = 0
		cleanTimestamp(&actualAuthor)
		tests.AssertEquals(t, expectedErr, actualErr)
		tests.AssertEquals(t, expectedAuthor, actualAuthor)
	})
	t.Run("Update already deleted author", func(t *testing.T) {
		var expectedErr string = "record not found"
		id, _ := insertAuthor()
		db.Where("id = ?", id).Delete(&structs.Author{})
		expectedAuthor := structs.Author{
			ID:    id,
			Name:  "Deleted user",
			Email: "test@author.com",
		}
		actualErr := authorRepo.Update(id, expectedAuthor)
		tests.AssertEquals(t, expectedErr, actualErr.Error())
	})
	t.Run("Update passing extra fields", func(t *testing.T) {
		// I shouldn't be able to change metadata fields such as ID and CreatedAt
		id, _ := insertAuthor()
		fakeId := 4812
		fakeCreatedAt := time.Date(2001, 12, 31, 23, 58, 58, 2, time.UTC)
		var expectedErr error = nil
		updatedAuthor := structs.Author{
			ID:        fakeId,
			Name:      "author",
			Email:     "author@author.com",
			CreatedAt: fakeCreatedAt,
		}
		actualErr := authorRepo.Update(id, updatedAuthor)
		actualAuthor := queryAuthor(id)
		tests.AssertEquals(t, expectedErr, actualErr)
		if actualAuthor.ID == fakeId {
			t.Errorf("Author update is updating id")
		}
		if actualAuthor.CreatedAt == fakeCreatedAt {
			t.Errorf("Author update is updating created_at")
		}
	})
	t.Run("Update unexisting author", func(t *testing.T) {
		fakeId := 4813
		var expectedErr string = "record not found"
		actualErr := authorRepo.Update(fakeId, structs.Author{})
		tests.AssertEquals(t, expectedErr, actualErr.Error())
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
		tests.AssertEquals(t, structs.Author{}, actualAuthor) // Should still exist as we're using soft deletes
	})
	t.Run("Delete already deleted author", func(t *testing.T) {
		id, _ := insertAuthor()
		var expectedErr string = "record not found"
		db.Where("id = ?", id).Delete(&structs.Author{})
		actualErr := authorRepo.Delete(id)
		tests.AssertEquals(t, expectedErr, actualErr.Error())
	})
	t.Run("Delete unexistent author", func(t *testing.T) {
		var expectedErr string = "record not found"
		fakeId := 347823
		actualErr := authorRepo.Delete(fakeId)
		tests.AssertEquals(t, expectedErr, actualErr.Error())
	})
}
