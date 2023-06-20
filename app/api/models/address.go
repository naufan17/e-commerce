package models

type (
	Address struct {
		Address_ID       int    `json:"address_id"`
		Shipping_Address string `json:"shipping_address"`
	}
)
