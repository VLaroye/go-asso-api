package app

import (
	"encoding/json"
	"net/http"
)

var membersDummyData = Members{Member{ID: 1, Name: "John"}, Member{ID: 2, Name: "Alice"}}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getMembersHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(membersDummyData)
}
