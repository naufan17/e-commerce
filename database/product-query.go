package database

import (
	"context"
	"fmt"
	"log"

	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/config"
)

const (
	table          = "product"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Product, error) {

	var products []models.Product

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var product models.Product

		if err = rowQuery.Scan(&product.ID,
			&product.Name,
		); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}
