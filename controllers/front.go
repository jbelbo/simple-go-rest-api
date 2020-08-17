package controllers

import (
	"encoding/json"
	"net/http"
)

// RegisterControllers initializes app controllers
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	encoded, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("JSON encoding error"))
	}
	w.Write(encoded)
}
