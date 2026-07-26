package main

import "log"

func main() {
	app := &app{
		config: Config{
			addr: ":8080",
		},
	}
	log.Fatal(app.run(app.route()))
}
