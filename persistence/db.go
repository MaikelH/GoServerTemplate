package persistence

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	// Import the pgx driver
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"goservertemplate/types"
	"log/slog"
)

// Embedded files
//
//go:embed migrations/*
var embeddedMigrations embed.FS

// InitDatabase initializes the database connection
// It also runs the database migrations if the configuration specifies to do so.
func InitDatabase(ctx context.Context, config *types.Configuration) (*sql.DB, error) {
	slog.Info("Connecting to database")
	if config.DatabaseURL == "" {
		slog.Error("Database URL is empty")
		return nil, fmt.Errorf("database URL is empty")
	}
	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	slog.Info("Database connection established")
	// Check database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if config.RunMigrations {
		err = runMigrations(ctx, config)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// runMigrations runs the database migrations
// The migrations are embedded in the binary
func runMigrations(_ context.Context, config *types.Configuration) error {
	slog.Info("Running migrations")

	goose.SetBaseFS(embeddedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		return err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			slog.Error("Failed to close database connection", "service_error", err)
		}
	}()

	err = goose.Up(db, "migrations")
	if err != nil {
		return err
	}

	return nil
}
