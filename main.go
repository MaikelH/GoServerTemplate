package main

import (
	"log/slog"
	"os"
	"retrolink-backend/config"
	"retrolink-backend/server"
	"retrolink-backend/types"
)

func main() {
	slog.Info("Starting Retrolink backend server")

	slog.Info("Loading configuration")
	configuration, err := config.SetupConfig[types.Configuration]()
	if err != nil {
		slog.Error("error setup configuration", "error", err)
		os.Exit(1)
	}

	srv := server.NewServer(configuration)
	err = srv.Start()
	if err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(2)
	}
}
