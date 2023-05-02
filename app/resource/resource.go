package resource

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}
