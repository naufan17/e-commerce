package database

import (
	"context"
	"fmt"
	"log"

	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/config"
)

const (
	table_category = "categories"
)

func GetAllCategories(ctx context.Context) ([]models.Category, error) {

	var categories []models.Category

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By category_id ASC", table_category)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category models.Category

		if err = rowQuery.Scan(&category.Category_ID,
			&category.Category_Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
