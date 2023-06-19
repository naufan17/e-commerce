package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/api/middleware"
	"github.com/naufan17/e-commerce/app/api/models"
	"github.com/naufan17/e-commerce/app/api/resource"
	"github.com/naufan17/e-commerce/app/config"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := middleware.VerifyToken(tokenString)
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
	claims, err := middleware.VerifyToken(tokenString)
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

	stmt, err := db.Prepare("INSERT INTO carts (user_id, product_id, count) VALUES (?, ?, ?)")
	_, err = stmt.Exec(id, product_id, count)
	if err != nil {
		fmt.Fprintf(w, "Error inserting cart")
	} else {
		fmt.Fprintf(w, "Cart added successfully")
	}
}

func PutCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := middleware.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	unusedFuncVar := claims
	_ = unusedFuncVar

	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	cart_id := r.FormValue("cart_id")
	count := r.FormValue("count")

	_, err = db.Query("SELECT cart_id FROM carts WHERE cart_id = ?", cart_id)
	if err != nil {
		fmt.Fprintf(w, "Error updating cart")
		return
	}

	stmt, err := db.Prepare("UPDATE carts SET count = ? WHERE cart_id = ?")
	_, err = stmt.Exec(count, cart_id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Cart updated successfully")
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[len("Bearer "):]
	claims, err := middleware.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	unusedFuncVar := claims
	_ = unusedFuncVar

	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	id := params["id"]

	_, err = db.Query("SELECT cart_id FROM carts WHERE cart_id = ?", id)
	if err != nil {
		fmt.Fprintf(w, "Error updating cart")
		return
	}

	stmt, err := db.Prepare("DELETE FROM carts WHERE cart_id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Cart deleted successfully")
}
