package controllers

import (
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
	defer db.Close()

	rows, err := db.Query("SELECT address.address_id, address.shipping_address FROM address INNER JOIN users ON address.user_id = users.user_id WHERE users.username = ?", claims.Username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	addresses := make([]models.Address, 0)
	for rows.Next() {
		address := models.Address{}

		err := rows.Scan(&address.Address_ID,
			&address.Shipping_Address)
		if err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if len(addresses) == 0 {
		resource.ErrorHandler(w, "Address not found", http.StatusNotFound)
		return
	}

	resource.ResponseHandler(w, addresses, http.StatusOK)
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
	defer db.Close()

	var id string
	db.QueryRow("SELECT user_id FROM users WHERE username = ?", claims.Username).Scan(&id)

	shipping_address := r.FormValue("shipping_address")

	stmt, err := db.Prepare("INSERT INTO address (user_id, shipping_address) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id, shipping_address)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		resource.ResponseHandler(w, "Address added successfully", http.StatusOK)
	} else {
		resource.ErrorHandler(w, "Error inserting address", http.StatusNotFound)
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
	defer db.Close()

	address_id := r.FormValue("address_id")
	shipping_address := r.FormValue("shipping_address")

	stmt, err := db.Prepare("UPDATE address SET shipping_address = ? WHERE address_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(shipping_address, address_id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		resource.ResponseHandler(w, "Address updated successfully", http.StatusOK)
	} else {
		resource.ErrorHandler(w, "Error updated address", http.StatusNotFound)
	}
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
	defer db.Close()

	params := mux.Vars(r)
	address_id := params["address_id"]

	stmt, err := db.Prepare("DELETE FROM address WHERE address_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(address_id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		resource.ResponseHandler(w, "Address deleted successfully", http.StatusOK)
	} else {
		resource.ErrorHandler(w, "Error deleting address", http.StatusNotFound)
	}
}
