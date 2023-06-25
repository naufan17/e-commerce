package controllers

import (
	"log"
	"net/http"

	"github.com/naufan17/e-commerce/app/api/middleware"
	"github.com/naufan17/e-commerce/app/api/resource"
	"github.com/naufan17/e-commerce/app/config"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Get the username and password from the request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate the user (in this example, we just check that the username and password are not empty)
	if username == "" || password == "" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Insert Username and Password into database
	db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)

	// Create a JWT token for the user
	tokenString, err := middleware.CreateToken(username)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response body
	resource.DataResponse(w, tokenString, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	// Get the username and password from the request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate the user (in this example, we just check that the username and password are not empty)
	if username == "" || password == "" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Query database for user password
	var dbPassword string
	db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPassword)

	// Compare password user
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a JWT token for the user
	tokenString, err := middleware.CreateToken(username)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response body
	resource.DataResponse(w, tokenString, http.StatusOK)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
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

	// Return the username in the response body
	resource.DataResponse(w, claims.Username, http.StatusOK)
}
