package main

import (
	"github.com/gorilla/mux"
	"internal/api"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	api.RegisterRoutes(r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
