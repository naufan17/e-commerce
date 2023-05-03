package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/middleware"
	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/config"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	// Get the JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	// Parse and verify the JWT token
	tokenString := authHeader[len("Bearer "):]
	claims, err := middleware.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT carts.cart_id, products.product_name, products.price, carts.count FROM carts INNER JOIN products ON carts.product_id = products.product_id INNER JOIN users ON carts.user_id = users.user_id WHERE users.username = ? ORDER BY cart_id ASC", claims.Username)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	carts := make([]models.Cart, 0)
	for rows.Next() {
		cart := models.Cart{}

		err := rows.Scan(&cart.Cart_ID,
			&cart.Product_Name,
			&cart.Price,
			&cart.Count)

		if err != nil {
			log.Fatal(err)
		}
		carts = append(carts, cart)
	}
	resource.ResponseJSON(w, carts, http.StatusOK)
}

// func PostCart(w http.ResponseWriter, r *http.Request) {
// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var cart models.Cart
// 	err = json.NewDecoder(r.Body).Decode(&cart)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	result, err := db.Exec("INSERT INTO carts(user_id, product_id, count) VALUES(?, ?, ?)", cart.User_ID, cart.Product_ID, cart.Count)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	cart_id, _ := result.LastInsertId()
// 	cart.Cart_ID = int(cart_id)

// 	resource.ResponseJSON(w, cart, http.StatusOK)
// }

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// Get the JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	// Parse and verify the JWT token
	tokenString := authHeader[len("Bearer "):]
	claims, err := middleware.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, claims.Username+" Invalid authorization token", http.StatusUnauthorized)
		return
	}

	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	id := params["id"]

	stmt, err := db.Prepare("DELETE FROM carts WHERE cart_id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Fprintf(w, "Error deleting chart")
		return
	}

	fmt.Fprintf(w, "Chart deleted successfully")
}
