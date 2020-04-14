package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

func InitDB(dbURI string) {
	var err error
	// TODO: Put db credentials into env file
	db, err = gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Could not connect to db, error: %d", err)
	} else {
		log.Print("Connected to db")
	}

	// defer db.Close()

	db.AutoMigrate(&Member{})
}

func CloseDB() {
	err := db.Close()

	if err != nil {
		log.Fatalf("Error closing db connection: %d", err)
	}
}
