package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/VLaroye/api/app"
)

func main() {
	router := app.NewRouter()

	// TODO: Put db credentials into env file
	db, err := gorm.Open("postgres", "host=db port=5432 user=admin dbname=asso password=admin123 sslmode=disable")

	if err != nil {
		log.Fatalf("Could not connect to db, error: %d", err)
	} else {
		log.Print("Connected to db")
	}

	defer db.Close()

	db.AutoMigrate(&app.Member{})

	log.Fatal(http.ListenAndServe(":8080", router))
}
