package utils

import (
	"encoding/json"
	"net/http"
)

func WriteErrorJSON(w http.ResponseWriter, err error, code int) {
	WriteJson(w, map[string]string{"error": err.Error()}, code)
}

func WriteJson(w http.ResponseWriter, payload interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(payload)
}
