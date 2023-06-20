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
	r.HandleFunc("/address", controllers.GetAddress).Methods("GET")
	r.HandleFunc("/address", controllers.PostAddress).Methods("POST")
	r.HandleFunc("/address", controllers.PutAddress).Methods("PUT")
	r.HandleFunc("/address/{address_id}", controllers.DeleteAddress).Methods("DELETE")
	r.HandleFunc("/categories", controllers.GetCategory).Methods("GET")
	r.HandleFunc("/products", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/carts", controllers.GetCart).Methods("GET")
	r.HandleFunc("/carts", controllers.PostCart).Methods("POST")
	r.HandleFunc("/carts", controllers.PutCart).Methods("PUT")
	r.HandleFunc("/carts/{cart_id}", controllers.DeleteCart).Methods("DELETE")
}
