package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var validate *validator.Validate = validator.New()

var membersDummyData = Members{Member{ID: 1, Name: "John"}, Member{ID: 2, Name: "Alice"}}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getMembersHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(membersDummyData)
}

func createMemberHandler(w http.ResponseWriter, req *http.Request) {
	var member Member

	err := json.NewDecoder(req.Body).Decode(&member)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	// TODO: Generalize validation process + Improve error messages
	err = validate.Struct(member)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Errorf("Field %v: '%v'", err.Field(), err.Tag())
			log.Print(errorMessage)

			errors = append(errors, errorMessage.Error())
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": false, "errors": errors})
		return
	}

	db.NewRecord(member)
	db.Create(&member)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
