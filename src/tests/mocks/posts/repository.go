package posts

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)


type MockPostRepository struct {
	interfaces.PostRepository
}

func (mpr MockPostRepository) Create(structs.Post) (int, error) {
	return 0, nil
} 

func (mpr MockPostRepository) Get(id int) (structs.Post, error) {
	return structs.Post{}, nil
}

func (mpr MockPostRepository) Update(structs.Post) error {
	return nil
}

func (mpr MockPostRepository) Delete(id int) error {
	return nil
}
