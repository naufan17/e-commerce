package controllers

import (
	"log"
	"net/http"

	"github.com/naufan17/e-commerce/app/api/models"
	"github.com/naufan17/e-commerce/app/api/resource"
	"github.com/naufan17/e-commerce/app/config"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT category_id, category_name FROM categories Order By category_id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		category := models.Category{}

		err := rows.Scan(&category.Category_ID,
			&category.Category_Name)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if len(categories) == 0 {
		resource.ErrorHandler(w, "Category not found", http.StatusNotFound)
		return
	}

	resource.ResponseHandler(w, categories, http.StatusOK)
}
