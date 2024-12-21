package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	http.ListenAndServe(":8080", nil)
}
