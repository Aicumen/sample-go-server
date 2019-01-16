package main

import (
	"log"
	"net/http"

	sw "./go"
	"github.com/gorilla/handlers"
)

func main() {
	log.Printf("Server started")
	sw.StartRedis()
	sw.StartMongo()
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
