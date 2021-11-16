package authors

import (
	"github.com/eduardothsantos/go-blog/src/domain/input"
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)

func authorValidation(author *structs.Author) error {
	author.Name = input.TransformSingleLine(author.Name)
	author.Email = input.TransformSingleLine(author.Email)
	err := input.ValidateNameField(author.Name)
	if err != nil {
		return err
	}
	err = input.ValidateEmailField(author.Email)
	if err != nil {
		return err
	}
	return nil
}

type AuthorService struct {
	Repo interfaces.AuthorRepository
}

func NewAuthorService(repo interfaces.AuthorRepository) AuthorService {
	return AuthorService{
		Repo: repo,
	}
}

func (as AuthorService) Create(author structs.Author) (int, error) {
	err := authorValidation(&author)
	if err != nil {
		return 0, err
	}
	return as.Repo.Create(author)
}

func (as AuthorService) Get(id int) (structs.Author, error) {
	return as.Repo.Get(id)
}

func (as AuthorService) Update(id int, author structs.Author) error {
	err := authorValidation(&author)
	if err != nil {
		return err
	}
	return as.Repo.Update(id, author)
}

func (as AuthorService) Delete(id int) error {
	return as.Repo.Delete(id)
}
