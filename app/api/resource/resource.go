package resource

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func DataResponse(w http.ResponseWriter, data interface{}, status int) {
	response := Data{
		Status:  status,
		Message: "success",
		Data:    data,
	}

	ResponseJSON(w, response, status)
}

func MessageResponse(w http.ResponseWriter, message string, status int) {
	response := Message{
		Status:  status,
		Message: message,
	}

	ResponseJSON(w, response, status)
}

func ResponseJSON(w http.ResponseWriter, response interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
