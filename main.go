package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/naufan17/e-commerce/app/controller"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	router.HandleFunc("/categories", controller.GetCategory).Methods("GET")
	router.HandleFunc("/products", controller.GetProduct).Methods("GET")
	router.HandleFunc("/carts", controller.GetCart).Methods("GET")
	router.HandleFunc("/carts", controller.PostCart).Methods("POST")
	router.HandleFunc("/carts/{id}", controller.DeleteCart).Methods("DELETE")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Start the server
	log.Fatal(http.ListenAndServe(addr, router))
}
