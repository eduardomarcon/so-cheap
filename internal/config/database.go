package config

import (
	"database/sql"
	"fmt"
	"time"
)

func OpenConnection() (*sql.DB, error) {
	dbConfig := GetDB()
	db, err := sql.Open("postgres", dbConfig.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	db.SetMaxOpenConns(dbConfig.Max)
	db.SetMaxIdleConns(dbConfig.MaxIdle)
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTime))

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	return db, err
}
