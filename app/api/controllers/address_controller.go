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

func GetAddress(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Query("SELECT address.address_id, users.username, address.shipping_address FROM address INNER JOIN users ON address.user_id = users.user_id WHERE users.username = ? ", claims.Username)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	addresses := make([]models.Address, 0)
	for rows.Next() {
		address := models.Address{}

		err := rows.Scan(&address.Address_ID,
			&address.Username,
			&address.Shipping_Address)
		if err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, address)
	}

	resource.ResponseJSON(w, addresses, http.StatusOK)
}

func PostAddress(w http.ResponseWriter, r *http.Request) {
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

	shipping_address := r.FormValue("shipping_address")

	stmt, err := db.Prepare("INSERT INTO address (user_id, shipping_address) VALUES (?, ?)")
	_, err = stmt.Exec(id, shipping_address)
	if err != nil {
		fmt.Fprintf(w, "Error inserting address")
	} else {
		fmt.Fprintf(w, "Address added successfully")
	}
}

func PutAddress(w http.ResponseWriter, r *http.Request) {
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

	address_id := r.FormValue("address_id")
	shipping_address := r.FormValue("shipping_address")

	_, err = db.Query("SELECT address_id FROM address WHERE address_id = ?", address_id)
	if err != nil {
		fmt.Fprintf(w, "Error updating address")
		return
	}

	stmt, err := db.Prepare("UPDATE address SET shipping_address = ? WHERE address_id = ?")
	_, err = stmt.Exec(shipping_address, address_id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Address updated successfully")
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
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

	_, err = db.Query("SELECT address_id FROM address WHERE address_id = ?", id)
	if err != nil {
		fmt.Fprintf(w, "Error deleting address")
		return
	}

	stmt, err := db.Prepare("DELETE FROM address WHERE address_id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Address deleted successfully")
}
