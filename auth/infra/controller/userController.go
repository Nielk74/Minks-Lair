package controller

import (
	"cosmink/auth/core/usecase"
	"cosmink/auth/infra/repository"
	"cosmink/libs/route"
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type UserController struct{}

func (c UserController) RegisterRoutes(server *route.Server) {
	server.RegisterEvent("/user", "register", UserRegisterHandler)
	server.RegisterEvent("/user", "login", UserLoginHandler)
	server.RegisterEvent("/user", "profile", UserProfileHandler)
}

func UserRegisterHandler(socket socketio.Conn, message string) {
	// get the request body (username, password)
	var data map[string]string
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		log.Printf("failed to unmarshal message: %v", err)
		socket.Emit("error", "Could not parse message")
		return
	}
	username := data["username"]
	password := data["password"]

	authUseCase := usecase.NewAuthUseCase(repository.UserRepository{})
	_, err := authUseCase.Register(username, password)
	if err != nil {
		log.Printf("failed to register user: %v", err)
		socket.Emit("error", "Internal Server Error")
		return
	}
	socket.Emit("message", "User registered successfully")
}
func UserLoginHandler(socket socketio.Conn, message string) {
	var data map[string]string
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		log.Printf("failed to unmarshal message: %v", err)
		socket.Emit("error", "Could not parse message")
		return
	}
	username := data["username"]
	password := data["password"]

	authUseCase := usecase.NewAuthUseCase(repository.UserRepository{})
	token, err := authUseCase.Login(username, password)
	if err != nil {
		log.Printf("failed to login user: %v", err)
		socket.Emit("error", "Internal Server Error")
		return
	}
	response := map[string]string{"message": "Login successful", "token": token.TokenString}
	socket.Emit("message", response)

}
func UserProfileHandler(socket socketio.Conn, message string) {
	tokenString, err := getToken([]byte(message))
	if err != nil {
		log.Printf("failed to get token: %v", err)
		socket.Emit("error", "Could not parse message")
		return
	}
	authUseCase := usecase.NewAuthUseCase(repository.UserRepository{})
	user, err := authUseCase.GetUserByToken(tokenString)
	if err != nil {
		log.Printf("failed to get user by token: %v", err)
		socket.Emit("error", "Internal Server Error")
		return
	}
	response := map[string]string{"message": "User profile", "username": user.Username}
	socket.Emit("message", response)	
}

func getToken(message []byte) (string, error) {
	var data map[string]string
	if err := json.Unmarshal(message, &data); err != nil {
		return "", err
	}
	return data["token"], nil
}