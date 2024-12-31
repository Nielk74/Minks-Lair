package main

import (
	"cosmink/auth/infra/controller"
	"cosmink/libs/route"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	server := route.NewServer()
	server.SocketIoServer.OnConnect("/", func(socket socketio.Conn) error {
		log.Printf("connected: %s", socket.ID())
		return nil
	})
	registerControllers(server, []controller.Registerable{
		&controller.UserController{},
	})
	server.Serve()

}

func registerControllers(server *route.Server, controllers []controller.Registerable) {
	for _, c := range controllers {
		c.RegisterRoutes(server)
	}
}
