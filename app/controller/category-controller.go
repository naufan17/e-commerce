package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/database"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		products, err := database.GetAllCategories(ctx)

		if err != nil {
			fmt.Println(err)
		}

		resource.ResponseJSON(w, products, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}
