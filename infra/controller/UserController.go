package controller

import (
	"encoding/json"
	"net/http"
)

// PingHandler is a simple handler that returns a JSON response with a message
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

// UserHandler is a simple handler that returns a JSON response with a message
func UserHandler(w http.ResponseWriter, r *http.Request) {
	handled := false
	handle := func(handler func(http.ResponseWriter, *http.Request)) {
		handled = true
		handler(w, r)
	}
	
	if r.Method == http.MethodPost {
		if r.URL.Path == "/user/register" {
			handle(UserRegisterHandler)
		}
		if r.URL.Path == "/user/login" {
			handle(UserLoginHandler)
		}
	}
	if r.Method == http.MethodGet {
		handle(UserProfileHandler)
	}
	if !handled {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "register"})
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "login"})
}
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "profile"})
}
