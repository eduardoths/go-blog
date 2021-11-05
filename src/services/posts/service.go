package posts

import (
	"errors"

	"github.com/eduardothsantos/go-blog/src/interfaces"
	"github.com/eduardothsantos/go-blog/src/structs"
)

type PostService struct {
	Repo interfaces.PostRepository
}

func NewPostService(repo interfaces.PostRepository) PostService {
	return PostService{
		Repo: repo,
	}
}

func (ps PostService) Create(post structs.Post, authorId int) error {
	err := ps.Repo.Create(post, authorId)
	if err != nil {
		return errors.New("INTERNAL SERVER ERROR")
	}
	return err
}

func (ps PostService) Get(id int) (structs.Post, error) {
	var errorToReturn error
	post, err := ps.Repo.Get(id)
	
	if err != nil {
		errorToReturn = errors.New("INTERNAL SERVER ERROR")
	} else {
		errorToReturn = err
	}
	return post, errorToReturn
}

func (ps PostService) Update(id int, post structs.Post) error {
	err := ps.Repo.Update(id, post)
	if err != nil {
		return errors.New("INTERNAL SERVER ERROR")
	}
	return err
}

func (ps PostService) Delete(id int) error {
	err := ps.Repo.Delete(id)
	if err != nil {
		return errors.New("INTERNAL SERVER ERROR")
	}
	return err
}