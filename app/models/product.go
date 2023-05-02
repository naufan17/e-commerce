package models

type (
	Product struct {
		Product_ID    int    `json:"product_id"`
		Product_Name  string `json:"product_name"`
		Category_Name string `json:"category_name"`
		Price         int    `json:"price"`
		Count         int    `json:"count"`
	}
)
