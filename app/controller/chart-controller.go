package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufan17/e-commerce/app/models"
	"github.com/naufan17/e-commerce/app/resource"
	"github.com/naufan17/e-commerce/config"
)

func PostChart(w http.ResponseWriter, r *http.Request) {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	var chart models.Chart
	json.NewDecoder(r.Body).Decode(&chart)

	result, err := db.Exec("INSERT INTO charts(user_id, product_id, count) VALUES(?, ?, ?)", chart.User_ID, chart.Product_ID, chart.Count)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	chart.Chart_ID = int(id)

	resource.ResponseJSON(w, result, http.StatusOK)
}

func DeleteChart(w http.ResponseWriter, r *http.Request) {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)

	result, err := db.Exec("DELETE FROM charts WHERE charts_id = ?", params)
	if err != nil {
		log.Fatal(err)
	}

	resource.ResponseJSON(w, result, http.StatusOK)
}
