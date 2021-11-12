package main

import (
	"log"

	"github.com/eduardothsantos/go-blog/migrations"
	"github.com/eduardothsantos/go-blog/pkg/databases"
)

func main() {
	db := databases.Config()
	err := db.AutoMigrate(&migrations.Author{})
	if err != nil {
		log.Fatalf("Couldn't migrate authors, error: %v", err.Error())
	}

	err = db.AutoMigrate(&migrations.Post{})
	if err != nil {
		log.Fatalf("Couldn't migrate posts, error: %v", err.Error())
	}
}
