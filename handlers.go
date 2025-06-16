package main

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Hello from the chessare-backend"}
	json.NewEncoder(w).Encode(response)
}