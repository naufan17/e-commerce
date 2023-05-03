package models

type (
	Cart struct {
		Cart_ID      int    `json:"cart_id"`
		Product_Name string `json:"product_name"`
		Price        int    `json:"price"`
		Count        int    `json:"count"`
	}
)
