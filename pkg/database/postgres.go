package database

import (
	"database/sql"
	"fmt"
	"github.com/TotrazOrstO/Imapro/pkg/config"

	_ "github.com/lib/pq"
)

func InitPostgres(cfg config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User, cfg.Password, cfg.DB, cfg.Host, cfg.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %v", err)
	}

	return db, nil
}
