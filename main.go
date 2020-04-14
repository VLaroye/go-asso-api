package main

import (
	"log"
	"net/http"

	"github.com/VLaroye/api/app"
)

func main() {
	router := app.NewRouter()

	app.InitDB("host=db port=5432 user=admin dbname=asso password=admin123 sslmode=disable")

	defer app.CloseDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
