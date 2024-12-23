package main

import (
	"cosmink/auth/infra/controller"
	"cosmink/libs/route"
	"encoding/json"
	"net/http"
)

type IController interface {
	RegisterRoutes(server *route.Server)
}

func main() {
	var control IController
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	server := route.NewServer()
	http.HandleFunc("/", server.ServeHTTP)
	userController := controller.UserController{}
	control = userController
	control.RegisterRoutes(server)
	http.ListenAndServe(":8080", nil)
}
