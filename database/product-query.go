package database

import (
	"context"
	"fmt"
	"log"

	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/config"
)

const (
	table_product = "products"
)

func GetAllProducts(ctx context.Context) ([]models.Product, error) {

	var products []models.Product

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By product_id ASC", table_product)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var product models.Product

		if err = rowQuery.Scan(&product.Product_ID,
			&product.Product_Name,
			&product.Category_ID,
			&product.Price,
			&product.Count,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
