package main

import (
	"log"
	"net/http"

	"github.com/naufan17/e-commerce/app/controller"
)

func main() {
	http.HandleFunc("/product", controller.GetProduct)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
