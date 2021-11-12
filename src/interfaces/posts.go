package interfaces

import (
	"github.com/eduardothsantos/go-blog/src/structs"
)

type PostRepository interface {
	Create(post structs.Post) (int, error)
	Get(id int) (structs.Post, error)
	Update(id int, post structs.Post) error
	Delete(id int) error
}

type PostService interface {
	Create(post structs.Post) (int, error)
	Get(id int) (structs.Post, error)
	Update(id int, post structs.Post) error
	Delete(id int) error
}

type PostHandler interface {
	Route()
}