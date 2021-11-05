package databases

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/eduardothsantos/go-blog/internal/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Config() (*sql.DB) {
	if db != nil {
		return db
	}
	dbConfig := config.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							dbConfig.Host,
							dbConfig.Port,
							dbConfig.User,
							dbConfig.Password,
							dbConfig.Database)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Couldn't connect to database, error: %v", err.Error())
	}	
	return db
}

func TestConfig() (*sql.DB) {
	if db != nil {
		return db
	}
	dbConfig := config.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							dbConfig.Host,
							dbConfig.Port,
							dbConfig.User,
							dbConfig.Password,
							"test")
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Couldn't connect to database, error: %v", err.Error())
	}	
	query, err := ioutil.ReadFile("/Users/Isaac/codes/go-blog/migrations/migration.sql")
	if err != nil {
		log.Fatalf("Couldn't read file %v, error: %v", "migrations/migration.sql", err.Error())
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatalf("Couldn't execute query, error: %v", err.Error())
	}
	return db
}