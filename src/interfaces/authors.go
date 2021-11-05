package interfaces

import (
	"github.com/eduardothsantos/go-blog/src/structs"
)

type AuthorRepository interface {
	Create(author structs.Author) error
	Get(id int) (structs.Author, error)
	Update(id int, author structs.Author) error
	Delete(id int) error
}

type AuthorService interface {
	Create(author structs.Author) error
	Get(id int) (structs.Author, error)
	Update(id int, author structs.Author) error
	Delete(id int) error
}

type AuthorHandler interface {
	Route()
}