package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Akanibekuly/golang_test_task1.git/internal/config"
	_ "github.com/jackc/pgx/stdlib"
)

func NewPostgresDB(cfg *config.DBConf) (*sql.DB, error) {
	db, err := sql.Open(cfg.Dialect, fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Port,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("ping error: postgres", err)
		return nil, err
	}

	return db, nil
}
