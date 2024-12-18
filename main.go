package main

import (
	"goservertemplate/config"
	"goservertemplate/server"
	"goservertemplate/types"
	"log/slog"
	"os"
)

func main() {
	configuration, err := config.SetupConfig[types.Configuration]()
	if err != nil {
		slog.Error("error setup configuration", "error", err)
		os.Exit(1)
	}

	setupLogger(configuration)

	slog.Info("Starting server")
	srv := server.NewServer(configuration)
	err = srv.Start()
	if err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(2)
	}
}

func setupLogger(c *types.Configuration) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(c.LogLevel)}))

	slog.SetDefault(logger)
	slog.Info("Setting logger to level", "level", c.LogLevel)
}
