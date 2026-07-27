package main

import (
	"log"

	"alilacream/socialx/internal/db"
	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"
)

func main() {
	allVars := env.ReturnAll()
	app := &app{
		config: Config{
			db: models.DBConfig{
				DSN:                allVars.DB_url,
				MaxOpenConnections: allVars.MaxOpenConns,
				MaxIdleConnections: allVars.MaxIdleConns,
				MaxIdleTime:        allVars.MaxIdleTime,
			},
			addr: allVars.Port,
		},
	}
	db, err := db.New(app.config.db)
	if err != nil {
		log.Println("Setting the connection to the database has gone wrong")
	}
	store := store.NewPQStorage(db)
	app.store = store
	defer db.Close()
	log.Fatal(app.run(app.route()))
}
