package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

}

var config Config

func init() {
	err := godotenv.Load("/Users/Isaac/codes/go-blog/.env")
	if err != nil {
		log.Panicf("Error loading environment variables, err: %v", err.Error())
	}
}

func GetConfig() Config {
	if config == (Config{}) {
		config = Config {
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		}
	}
	return config
}