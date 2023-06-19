package routes

import (
	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/api/controllers"
)

func HandleRoutes(r *mux.Router) {
	// Define API endpoints
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/profile", controllers.ProfileHandler).Methods("GET")
	r.HandleFunc("/categories", controllers.GetCategory).Methods("GET")
	r.HandleFunc("/products", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/carts", controllers.GetCart).Methods("GET")
	r.HandleFunc("/carts", controllers.PostCart).Methods("POST")
	r.HandleFunc("/carts", controllers.PutCart).Methods("PUT")
	r.HandleFunc("/carts/{id}", controllers.DeleteCart).Methods("DELETE")
}
