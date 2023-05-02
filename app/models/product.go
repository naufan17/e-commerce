package models

type (
	Product struct {
		Product_ID   int    `json:"product_id"`
		Product_Name string `json:"product_name"`
		Category_ID  int    `json:"category_id"`
		Price        int    `json:"price"`
		Count        int    `json:"count"`
	}
)
