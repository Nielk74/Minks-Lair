package main

import (
	"cosmink/auth/infra/controller"
	"cosmink/libs/route"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	server := route.NewServer()
	http.HandleFunc("/", server.ServeHTTP)

	registerControllers(server, []controller.Registerable{
		&controller.UserController{},
	})
	http.ListenAndServe(":8080", nil)
}

func registerControllers(server *route.Server, controllers []controller.Registerable) {
	for _, c := range controllers {
		c.RegisterRoutes(server)
	}
}