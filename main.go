package main

import (
	"cosmink/auth/infra/controller"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	// route start with user wildcard
	http.HandleFunc("/user/", controller.UserHandler)
	http.ListenAndServe(":8080", nil)
}
