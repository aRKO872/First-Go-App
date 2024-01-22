package dbcontroller

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/aRKO872/first-go-app/config"
)

var DB *gorm.DB
var err error
 
 
func DatabaseConnection() {
	DB, err = gorm.Open(postgres.Open(config.GetConfig().DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	log.Println("db connection successful")
}