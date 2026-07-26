package main

import (
	"log"

	"alilacream/socialx/internal/env"
)

func main() {
	port, err := env.GetVar("PORT")
	if err != nil {
		log.Println("No port configured gang: ", err.Error())
	}
	app := &app{
		config: Config{
			addr: port,
		},
	}
	log.Fatal(app.run(app.route()))
}
