package authors

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)

type MockAuthorRepository struct {
	interfaces.PostRepository
	FakeAuthor structs.Author
}

func (mar MockAuthorRepository) Create(author structs.Author) (int, error) {
	return 0, nil
}

func (mar MockAuthorRepository) Get(id int) (structs.Author, error) {
	return mar.FakeAuthor, nil
}

func (mar MockAuthorRepository) Update(id int, author structs.Author) error {
	return nil
}

func (mar MockAuthorRepository) Delete(id int) error {
	return nil
}
