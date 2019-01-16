package main

import (
	"log"
	"net/http"

	_ "github.com/akshay111meher/sample-go-server/docs"
	server "github.com/akshay111meher/sample-go-server/go"
	"github.com/akshay111meher/sample-go-server/go/db"
	keyMgmt "github.com/akshay111meher/sample-go-server/go/key_management"
	"github.com/go-chi/chi"
	"github.com/gorilla/handlers"
	"github.com/swaggo/http-swagger"
)

func main() {
	r := chi.NewRouter()

	log.Printf("Server started")
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8000", r)

	keyMgmt.StartRedis()
	db.StartMongo()
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
