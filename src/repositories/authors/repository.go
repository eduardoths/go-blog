package authors

import (
	"errors"

	"github.com/eduardothsantos/go-blog/src/structs"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

type authorUpdate struct {
	Name string
	Email string
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return AuthorRepository{
		db: db,
	}
}

func (ar AuthorRepository) Create(author structs.Author) (int, error) {
	tx := ar.db.Save(&author)
	return author.ID, tx.Error
}

func (ar AuthorRepository) Get(id int) (structs.Author, error) {
	var author structs.Author
	tx := ar.db.Where("id = ?", id).Take(&author)
	return author, tx.Error
}

func (ar AuthorRepository) Update(id int, author structs.Author) error {
	authorToUpdate := authorUpdate{
		Name: author.Name,
		Email: author.Email,
	}
	tx := ar.db.Model(structs.Author{}).Where("id = ?", id).Updates(&authorToUpdate)
	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return tx.Error
}

func (ar AuthorRepository) Delete(id int) error {
	tx := ar.db.Where("id = ?", id).Delete(&structs.Author{})
	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return tx.Error
}