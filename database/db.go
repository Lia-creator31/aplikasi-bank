package database

import (
	"create-api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "bismillah"
	dbPort   = "5432"
	dbname   = "finalassignment"
	db       *gorm.DB
	err      error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
	return db
}

func GetDB() *gorm.DB {
	return db
}
