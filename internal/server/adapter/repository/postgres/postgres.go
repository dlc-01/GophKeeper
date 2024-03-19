package postgres

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/conf"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var EmbedMigrations embed.FS

func NewSQLClient(cfg conf.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("error while opening conection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while ping db: %w", err)
	}

	goose.SetBaseFS(EmbedMigrations)

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, fmt.Errorf("error while seting dialect: %w", err)
	}

	err = goose.Up(db, "migration")
	if err != nil {
		return nil, fmt.Errorf("error migarition: %w", err)
	}

	return db, nil
}
