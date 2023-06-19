package models

type (
	Address struct {
		Address_ID       int    `json:"address_id"`
		Username         string `json:"name"`
		Shipping_Address string `json:"shipping_address"`
	}
)
