package main

import (
	"log"
	"net/http"

	_ "github.com/akshay111meher/sample-go-server/docs"
	"github.com/akshay111meher/sample-go-server/go"
	"github.com/akshay111meher/sample-go-server/go/db"
	"github.com/akshay111meher/sample-go-server/go/key_management"
	"github.com/gorilla/handlers"
	"github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /v1
// @host localhost:8000
func main() {

	// go startChi()
	keymanagement.StartRedis()
	db.StartMongo()
	router := server.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

// func startChi() {
// 	r := chi.NewRouter()

// 	r.Get("/swagger/*", httpSwagger.WrapHandler)

// 	http.ListenAndServe(":8000", r)
// }
