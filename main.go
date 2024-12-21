package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var db = make(map[string]string)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/user/")
		value, ok := db[name]
		response := make(map[string]string)
		response["user"] = name
		if ok {
			response["value"] = value
		} else {
			response["status"] = "no value"
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkCredentials(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var payload struct {
			Value string `json:"value"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Value == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid JSON or missing value"}`))
			return
		}

		db[user] = payload.Value
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.ListenAndServe(":8080", nil)
}

func checkCredentials(user, pass string) bool {
	credentials := map[string]string{
		"foo":  "bar",
		"manu": "123",
	}
	if p, ok := credentials[user]; ok {
		return p == pass
	}
	return false
}
