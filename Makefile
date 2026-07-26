# Dev environement to run 
dev: 
	go run cmd/api/*go
showenv:
	cat env | grep "$(WANTED)"
build:
	go build -o ./bin/main ./cmd/api/*.go 
