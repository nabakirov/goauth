package main

import (
	"log"
	"net/http"
	"os"

	"goauth/auth"

	"github.com/gorilla/handlers"
)

var port = os.Getenv("PORT")
var tokenFilePath = os.Getenv("TOKENS_FILE_PATH")

func main() {
	if len(port) == 0 {
		port = "5000"
	}
	auth.LoadTokens(tokenFilePath, &auth.Tokens)

	router := auth.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Printf("running server on PORT %s", port)
	log.Printf("available tokens %+v", auth.Tokens)

	// launch server
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
