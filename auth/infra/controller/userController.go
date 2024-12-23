package controller

import (
	"cosmink/auth/core/usecase"
	"cosmink/auth/infra/repository"
	"cosmink/libs/route"
	"encoding/json"
	"log"
	"net/http"
)

type UserController struct{}

func (c UserController) RegisterRoutes(server *route.Server) {
	server.RegisterRoute("/user/register", route.POST, UserRegisterHandler)

}

// PingHandler is a simple handler that returns a JSON response with a message
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

// UserHandler is a simple handler that returns a JSON response with a message
// func UserHandler(w http.ResponseWriter, r *http.Request) {
// 	authUseCase := usecase.NewAuthUseCase(repository.UserRepository{})
// 	handled := false
// 	handle := func(handler func(http.ResponseWriter, *http.Request, usecase.AuthUseCase)) {
// 		handled = true
// 		handler(w, r, authUseCase)
// 	}

// 	if r.Method == http.MethodPost {
// 		if r.URL.Path == "/user/register" {
// 			handle(UserRegisterHandler)
// 		}
// 		if r.URL.Path == "/user/login" {
// 			handle(UserLoginHandler)
// 		}
// 	}
// 	if r.Method == http.MethodGet {
// 		handle(UserProfileHandler)
// 	}
// 	if !handled {
// 		http.Error(w, "Not Found", http.StatusNotFound)
// 	}
// }

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// get the request body (username, password)
	username := r.FormValue("username")
	password := r.FormValue("password")
	authUseCase := usecase.NewAuthUseCase(repository.UserRepository{})
	_, err := authUseCase.Register(username, password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// return a JSON response with a message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "register"})
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request, authUseCase usecase.AuthUseCase) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "login"})
}
func UserProfileHandler(w http.ResponseWriter, r *http.Request, authUseCase usecase.AuthUseCase) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "profile"})
}
