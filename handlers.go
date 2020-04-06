package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func getMembersHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(membersDummyData)
}
