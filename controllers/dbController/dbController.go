package dbcontroller

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
 
 
func DatabaseConnection() {
	log.Println(os.Getenv("DATABASE_URL"))
	// DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open("postgres://postgres@localhost:5432/todo-go-db?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	log.Println("db connection successful")
}