package main

import (
	"io/ioutil"
	"log"

	"github.com/eduardothsantos/go-blog/pkg/databases"
)

func main() {
	db := databases.Config()
	defer db.Close()
	query, err := ioutil.ReadFile("migrations/migration.sql")
	if err != nil {
		log.Panicf("Couldn't read file %v, error: %v", "migrations/migration.sql", err.Error())
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Printf("Couldn't execute query, error: %v", err.Error())
	}
}