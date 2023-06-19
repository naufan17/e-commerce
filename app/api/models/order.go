package models

type (
	Order struct {
		Order_ID         int    `json:"order_id"`
		Product_Name     string `json:"product_name"`
		Price            int    `json:"price"`
		Count            int    `json:"count"`
		Shipping_Address string `json:"shipping_address"`
	}
)
