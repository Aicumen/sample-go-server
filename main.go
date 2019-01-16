package main

import (
	"log"
	"net/http"

	server "github.com/akshay111meher/sample-go-server/go"
	"github.com/akshay111meher/sample-go-server/go/db"
	keyMgmt "github.com/akshay111meher/sample-go-server/go/key_management"
	"github.com/gorilla/handlers"
)

func main() {
	log.Printf("Server started")
	keyMgmt.StartRedis()
	db.StartMongo()
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
