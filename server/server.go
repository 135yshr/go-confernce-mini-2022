package server

import (
    "net/http"
)

type Server struct {
}

func New() *Server {
    return &Server{}
}

func (s *Server) AddHandle(pattern string, handler http.Handler) {
    http.Handle(pattern, handler)
}

func (s *Server) Start() error {
    return http.ListenAndServe(":8080", nil)
}

