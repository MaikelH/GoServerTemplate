package server

import (
	"goservertemplate/http_server"
	"goservertemplate/types"
)

type Server struct {
	Config *types.Configuration
}

func (s *Server) Start() error {
	return http_server.StartHTTPServer(s.Config)
}

func NewServer(config *types.Configuration) *Server {
	return &Server{
		Config: config,
	}
}
