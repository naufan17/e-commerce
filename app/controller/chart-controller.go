package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/authentication"
	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/config"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := authentication.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	db, err := config.DBConnect()
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

func PostCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := authentication.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	var id string
	db.QueryRow("SELECT user_id FROM users WHERE username = ?", claims.Username).Scan(&id)

	product_id := r.FormValue("product_id")
	count := r.FormValue("count")

	stmt, err := db.Prepare("INSERT INTO carts(user_id, product_id, count) VALUES(?, ?, ?)")
	_, err = stmt.Exec(id, product_id, count)
	if err != nil {
		fmt.Fprintf(w, "Error inserting cart")
	}

	fmt.Fprintf(w, "Cart added successfully")
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := authentication.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, claims.Username+" Invalid authorization token", http.StatusUnauthorized)
		return
	}

	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	id := params["id"]

	stmt, err := db.Prepare("DELETE FROM carts WHERE cart_id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Fprintf(w, "Error deleting cart")
	}

	fmt.Fprintf(w, "Cart deleted successfully")
}
