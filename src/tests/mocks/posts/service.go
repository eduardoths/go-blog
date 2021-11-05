package posts

import (
	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)


type MockPostService struct {
	interfaces.PostRepository
}

func (mps MockPostService) Create(post structs.Post) error {
	return nil
} 

func (mps MockPostService) Get(id int) (structs.Post, error) {
	return structs.Post{}, nil
}

func (mps MockPostService) Update(id int, post structs.Post) error {
	return nil
}

func (mps MockPostService) Delete(id int) error {
	return nil
}
