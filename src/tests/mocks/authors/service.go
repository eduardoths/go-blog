package authors

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)


type MockAuthorService struct {
	interfaces.PostRepository
}

func (mas MockAuthorService) Create(post structs.Post) (int, error) {
	return 0, nil
} 

func (mas MockAuthorService) Get(id int) (structs.Post, error) {
	return structs.Post{}, nil
}

func (mas MockAuthorService) Update(id int, post structs.Post) error {
	return nil
}

func (mas MockAuthorService) Delete(id int) error {
	return nil
}

