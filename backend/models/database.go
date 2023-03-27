package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() (*gorm.DB, error) {

	dbUrl := fmt.Sprint(os.Getenv("DATABASE_URL"))

	d, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	d.AutoMigrate(&User{})
	db = d
	return d, err
}

func GetDB() *gorm.DB {
	return db
}
