package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/database"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		products, err := database.GetAllProducts(ctx)

		if err != nil {
			fmt.Println(err)
		}

		resource.ResponseJSON(w, products, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}
