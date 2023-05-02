package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/database"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		products, err := database.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		resource.ResponseJSON(w, products, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var prd models.Product

		if err := json.NewDecoder(r.Body).Decode(&prd); err != nil {
			resource.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := database.Insert(ctx, prd); err != nil {
			resource.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		resource.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
