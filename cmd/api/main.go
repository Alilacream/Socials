package main

import (
	"log"

	"alilacream/socialx/internal/db"
	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"
)

func main() {
	// getting all variables from the env
	allVars := env.ReturnAll()
	// config variable
	dbConf := models.DBConfig{
		DSN:                allVars.DB_url,
		MaxOpenConnections: allVars.MaxOpenConns,
		MaxIdleConnections: allVars.MaxIdleConns,
		MaxIdleTime:        allVars.MaxIdleTime,
	}
	db, err := db.New(&dbConf)
	if err != nil {
		log.Println("Connection to database is not established, error: ", err.Error())
		return
	}
	log.Println("Connected to DataBase")
	defer db.Close()
	store := store.NewPQStorage(db)
	app := &app{
		store: &store,
		config: Config{
			db:   &dbConf,
			addr: allVars.Port,
		},
	}
	log.Fatal(app.run(app.route()))
}
