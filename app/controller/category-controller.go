package controller

import (
	"log"
	"net/http"

	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/config"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM categories Order By category_id ASC")
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

	resource.ResponseJSON(w, categories, http.StatusOK)
}
