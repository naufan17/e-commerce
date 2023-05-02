package models

import (
	"time"
)

type (
	Product struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Price     int       `json:"price"`
		Count     int       `json:"count"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
