package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/naufan17/e-commerce/app/api/routes"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Forward the request to the route handler
	routes.HandleRoutes(r)

	// Get PORT from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Wrap the router with a logging middleware
	loggedRouter := routes.LogRequests(r)

	// Assign port from env file
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// Print listening port in terminal
	fmt.Println("Server listening on port: ", os.Getenv("PORT"))

	// Start the server
	log.Fatal(http.ListenAndServe(addr, loggedRouter))
}
