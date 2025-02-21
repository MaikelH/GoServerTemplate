package server

import (
	"context"
	"database/sql"
	"goservertemplate/httpserver"
	"goservertemplate/persistence"
	"goservertemplate/servicecontainer"
	"goservertemplate/types"
	"log/slog"
)

type Server struct {
	Config    *types.Configuration
	DB        *sql.DB
	container *servicecontainer.Container
}

func (s *Server) Start() error {
	return httpserver.StartHTTPServer(s.container)
}

func NewServer(config *types.Configuration) (*Server, error) {
	server := &Server{
		Config: config,
	}

	// Start Database server
	db, err := persistence.InitDatabase(context.Background(), config)
	if err != nil {
		return nil, err
	}
	server.DB = db
	slog.Info("Database connection established")

	server.container = servicecontainer.NewServiceContainer(config, db)

	return server, nil
}
