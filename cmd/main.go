package main

import (
	"github.com/brinton0739/golangwebserver/internal/api"
	"github.com/brinton0739/golangwebserver/internal/database"
	"github.com/brinton0739/golangwebserver/internal/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	api.RegisterRoutes(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
