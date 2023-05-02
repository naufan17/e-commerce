package models

type (
	Chart struct {
		Chart_ID     int    `json:"chart_id"`
		User_ID      int    `json:"user_id"`
		Product_ID   int    `json:"product_id"`
		Product_Name string `json:"product_name"`
		Price        int    `json:"price"`
		Count        int    `json:"count"`
	}
)
