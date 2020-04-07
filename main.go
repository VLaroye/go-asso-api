package main

import (
	"log"
	"net/http"

	"github.com/VLaroye/api/app"
)

func main() {
	router := app.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
