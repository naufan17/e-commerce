package resource

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseHandler(w http.ResponseWriter, data interface{}, status int) {
	response := Response{
		Status:  status,
		Message: "success",
		Data:    data,
	}

	ResponseJSON(w, response, http.StatusOK)
}

func ErrorHandler(w http.ResponseWriter, message string, status int) {
	response := Response{
		Status:  status,
		Message: message,
	}

	ResponseJSON(w, response, status)
}

func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
