package controllers

import (
	"log"
	"net/http"

	"github.com/naufan17/e-commerce/app/api/models"
	"github.com/naufan17/e-commerce/app/api/resource"
	"github.com/naufan17/e-commerce/app/config"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	category := r.URL.Query().Get("category")

	if category == "" {
		rows, err := db.Query("SELECT products.product_id, products.product_name, categories.category_name, products.price, products.count FROM products INNER JOIN categories ON products.category_id = categories.category_id ORDER BY product_id ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		products := make([]models.Product, 0)
		for rows.Next() {
			product := models.Product{}

			err := rows.Scan(&product.Product_ID,
				&product.Product_Name,
				&product.Category_Name,
				&product.Price,
				&product.Count)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		if len(products) == 0 {
			resource.ErrorHandler(w, "Product not found", http.StatusNotFound)
			return
		}

		resource.ResponseHandler(w, products, http.StatusOK)

	} else {
		rows, err := db.Query("SELECT products.product_id, products.product_name, categories.category_name, products.price, products.count FROM products INNER JOIN categories ON products.category_id = categories.category_id WHERE category_name LIKE ? ORDER BY product_id ASC", category)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		products := make([]models.Product, 0)
		for rows.Next() {
			product := models.Product{}

			err := rows.Scan(&product.Product_ID,
				&product.Product_Name,
				&product.Category_Name,
				&product.Price,
				&product.Count)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		if len(products) == 0 {
			resource.ErrorHandler(w, "Product not found", http.StatusNotFound)
			return
		}

		resource.ResponseHandler(w, products, http.StatusOK)
	}
}
