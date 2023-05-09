package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Open a database connection
func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		return nil, err
	}

	return db, nil
}
