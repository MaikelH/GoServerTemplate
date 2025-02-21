package repositories

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	// Import the postgres dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type UserRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func NewUserRepository(pgx *sql.DB) *UserRepository {
	dialect := goqu.Dialect("postgres")
	db := dialect.DB(pgx)

	return &UserRepository{db: pgx, goqu: db}
}
