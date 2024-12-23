package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// registerRoute nethod; path; handler
const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

type Server struct {
	routes   map[string]map[string]http.HandlerFunc
	NotFound http.HandlerFunc
	OnErr    http.HandlerFunc
}

func (s *Server) RegisterRoute(path string, method string, handler func(http.ResponseWriter, *http.Request)) bool {
	if s.routes[method] == nil {
		s.routes[method] = make(map[string]http.HandlerFunc)
	}
	s.routes[method][path] = handler
	return true
}

func NewServer() *Server {
	return &Server{
		routes:   make(map[string]map[string]http.HandlerFunc),
		NotFound: http.NotFound,
		OnErr:    OnErrHandler,
	}
}

func OnErrHandler(w http.ResponseWriter, r *http.Request) {
	simulatedError := func() error {
		return fmt.Errorf("an example error occurred")
	}
	err := simulatedError()
	if err != nil {
		errObj := struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errObj); err != nil {
			log.Printf("failed to encode error: %s\n", err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request handled successfully"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := s.routes[r.Method][r.URL.Path]
	if !ok {
		s.NotFound(w, r)
		return
	}
	handler(w, r)
}
