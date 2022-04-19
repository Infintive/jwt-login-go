package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Infintive/predictive-go/app/models"
	"github.com/Infintive/predictive-go/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "Infinitive007"

// Base route to access the API Documentation.
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gophers!")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Declare a new User struct.
	var u models.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// encrypt the password now
	passEncrypt, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	// set it back on the struct, before saving to DB
	u.Password = string(passEncrypt)

	// Do something with the Person struct...
	fmt.Fprintf(w, "new User: %+v", u)
	database.AddUser(u)

	fmt.Printf("New user registerd! %+v\n", u)
}

// Handler for login, will accept username and pass as map
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	// parse the input body as map
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := database.GetUserByEmail(data["email"])

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "User not found in DB"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Incorrect password for the user"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Incorrect username or password combination"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	// login success return the token on the output writer
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Login success"
	resp["token"] = token
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
