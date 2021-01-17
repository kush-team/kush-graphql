package db

import (
	"kush-graphql/app/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func OpenDB(database string) *gorm.DB {

	databaseDriver := os.Getenv("DATABASE_DRIVER")

	db, err := gorm.Open(databaseDriver, database)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if err := Automigrate(db); err != nil {
		panic(err)
	}
	return db
}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{}, &models.Theme{}).Error
}
