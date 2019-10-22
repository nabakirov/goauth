package main

import (
	"log"
	"net/http"

	"goauth/auth"

	"github.com/gorilla/handlers"
)

func main() {
	router := auth.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Printf("running server on PORT %d", 5000)
	// launch server
	log.Fatal(http.ListenAndServe(":5000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
