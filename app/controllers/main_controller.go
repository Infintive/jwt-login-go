package controllers

import (
	"fmt"
	"net/http"
)

// Base route to access the API Documentation.
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gophers!")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, register!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Login!")
}
