package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/controller"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/products", controller.GetProduct).Methods("GET")
	router.HandleFunc("/categories", controller.GetCategory).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
