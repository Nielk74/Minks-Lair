package route

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	SocketIoServer *socketio.Server
	routes         map[string]map[string]func(socketio.Conn, string)
	NotFound       func(socketio.Conn, string)
	OnErr          func(socketio.Conn, string)
}

func (server *Server) RegisterEvent(namespace string, event string, handler func(socketio.Conn, string)) error {
	if _, ok := server.routes[namespace]; !ok {
		server.routes[namespace] = make(map[string]func(socketio.Conn, string))
	}
	if _, ok := server.routes[namespace][event]; !ok {
		server.routes[namespace][event] = handler
		server.SocketIoServer.OnEvent(namespace, event, handler)
		log.Printf("Event %s registered on %s", event, namespace)
		return nil
	}
	return fmt.Errorf("event already exists")

}

func NewServer() *Server {
	SocketIoServer := socketio.NewServer(nil)
	return &Server{
		SocketIoServer: SocketIoServer,
		routes:         make(map[string]map[string]func(socketio.Conn, string)),
		NotFound:       NotFoundHandler,
		OnErr:          OnErrHandler,
	}
}

func (server *Server) Serve() {
	go server.SocketIoServer.Serve()
	defer server.SocketIoServer.Close()
	http.Handle("/socket.io/", server.SocketIoServer)
	http.ListenAndServe(":8080", nil)
}

func NotFoundHandler(socket socketio.Conn, message string) {
	log.Printf("404 Not Found: %s", message)
	socket.Emit("error", "404 Not Found")
}

func OnErrHandler(socket socketio.Conn, message string) {
	log.Printf("Internal Server Error: %s", message)
	socket.Emit("error", "Internal Server Error")
}
