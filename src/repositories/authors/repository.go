package authors

import (
	"database/sql"

	"github.com/eduardothsantos/go-blog/src/structs"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return AuthorRepository{
		db: db,
	}
}

func (ar AuthorRepository) Create(author structs.Author) error {
	_, err := ar.db.Exec("INSERT INTO authors (name, email) VALUES ($1, $2);", author.Name, author.Email)
	return err
}

func (ar AuthorRepository) Get(id int) (structs.Author, error) {
	var author structs.Author
	err := ar.db.QueryRow("SELECT name, email FROM authors WHERE id = $1;", id).Scan(&author.Name, &author.Email)
	return author, err
}

func (ar AuthorRepository) Update(id int, author structs.Author) error {
	_, err := ar.db.Exec("UPDATE authors SET name=$1, email=$2 WHERE id = $3;", 
	                     author.Name, author.Email, id)
	return err
}

func (ar AuthorRepository) Delete(id int) error {
	_, err := ar.db.Exec("DELETE FROM authors WHERE id = $1;", id)
	return err
}