package db

import (
	"context"
	"database/sql"
	"time"

	"alilacream/socialx/models"
)

func New(dbConf models.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConf.DSN)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(dbConf.MaxIdleConnections)
	db.SetMaxOpenConns(dbConf.MaxOpenConnections)
	duration, err := time.ParseDuration(dbConf.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	// setting the max time waiting for connecting to db
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
