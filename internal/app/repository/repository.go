package repository

import (
	"wallet/internal/app/config"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	postgresql *sqlx.DB
}

func New(cfg *config.Config) *Repository {
	dsn := cfg.Dbconfig.Dsn

	postgres, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	return &Repository{
		postgresql: postgres,
	}
}
