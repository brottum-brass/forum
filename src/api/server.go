package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
	mux  http.Handler
}

func NewServer(port int, router http.Handler) *Server {
	return &Server{
		port: port,
		mux:  router,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}
