package server

import (
	"retrolink-backend/http_server"
	"retrolink-backend/types"
)

type Server struct {
	Config *types.Configuration
}

func (s *Server) Start() error {
	http_server.StartHTTPServer(s.Config)
	return nil
}

func NewServer(config *types.Configuration) *Server {
	return &Server{
		Config: config,
	}
}
