# Dev environement to run 
dev: 
	go run cmd/api/*go
# show env of what i want
showenv:
	cat env | grep "$(WANTED)"
# building directly
build:
	go build -o ./bin/main ./cmd/api/*.go 
# my custome migrate func 
migrate:
	docker exec -i socialx_db psql -U alilacream -d social < ./scripts/"$(SCRIPT)"
