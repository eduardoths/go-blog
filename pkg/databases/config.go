package databases

import (
	"fmt"
	"log"

	"github.com/eduardothsantos/go-blog/internal/config"
	"github.com/eduardothsantos/go-blog/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Config() *gorm.DB {
	if db != nil {
		return db
	}
	dbConfig := config.GetConfig()
	psqlInfo := postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database))
	var err error
	db, err = gorm.Open(psqlInfo)
	if err != nil {
		log.Fatalf("Couldn't connect to database, error: %v", err.Error())
	}
	return db
}

func TestConfig() *gorm.DB {
	if db != nil {
		return db
	}
	dbConfig := config.GetConfig()
	psqlInfo := postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		"test"))
	var err error
	db, err = gorm.Open(psqlInfo)
	if err != nil {
		log.Fatalf("Couldn't connect to database, error: %v", err.Error())
	}
	db.Exec("DROP TABLE posts; DROP TABLE authors;")
	db.AutoMigrate(&migrations.Author{})
	db.AutoMigrate(&migrations.Post{})
	return db
}
