package main

import (
	"log"
	"net/http"
)

var membersDummyData = Members{Member{ID: 1, Name: "John"}, Member{ID: 2, Name: "Alice"}}

func main() {
	handleRequests()
}

func handleRequests() {
	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
