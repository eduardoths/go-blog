package posts

import (
	"github.com/eduardothsantos/go-blog/src/domain/input"
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

func (ps PostService) Create(post structs.Post) (int, error) {
	post.Title = input.TransformSingleLine(post.Title)
	id, err := ps.Repo.Create(post)
	return id, err
}

func (ps PostService) Get(id int) (structs.Post, error) {
	return ps.Repo.Get(id)
}

func (ps PostService) Update(id int, post structs.Post) error {
	post.Title = input.TransformSingleLine(post.Title)
	err := ps.Repo.Update(id, post)
	return err
}

func (ps PostService) Delete(id int) error {
	err := ps.Repo.Delete(id)
	return err
}
