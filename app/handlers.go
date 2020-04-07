package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var membersDummyData = Members{Member{ID: 1, Name: "John"}, Member{ID: 2, Name: "Alice"}}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func getMembersHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(membersDummyData)
}
