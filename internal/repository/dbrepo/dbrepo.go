package dbrepo

import (
	"database/sql"

	"github.com/flaviusp23/bookings/internal/config"
	"github.com/flaviusp23/bookings/internal/repository"
)

// here we can add many more DB connects like MySQL, MariaDB etc
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
