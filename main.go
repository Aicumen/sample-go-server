package main

import (
	"log"
	"net/http"

	"github.com/akshay111meher/sample-go-server/go"
	"github.com/akshay111meher/sample-go-server/go/db"
	"github.com/akshay111meher/sample-go-server/go/key_management"
	"github.com/gorilla/handlers"
)

func main() {

	keymanagement.StartRedis()
	db.StartMongo()
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
