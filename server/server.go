package server

import (
	"net/http"

	"github.com/135yshr/go-confernce-mini-2022/registry"
)

type Server struct {
	port string
}

func New(r *registry.Registry) *Server {
	http.Handle("/api/v1/users", r.UserHandler)
	return &Server{
		port: ":8080",
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.port, nil)
}
