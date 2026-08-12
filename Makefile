# Dev environementc to run 
dev:  
	@echo "Starting Dev Environement for testing..."
	@make -j2 backend frontend	
	
frontend:
	@echo "📦 Starting React frontend..."
	cd web && npm run dev 
backend:	
	@echo "⚙️  Starting Go backend with Air..."
	air
# show env of what i want
showenv:
	cat env | grep "$(WANTED)"
# building directly
build:
	go build -o ./bin/main ./cmd/api/*.go 
# my custome migrate func 
migrate:
	docker exec -i socialx_db psql -U alilacream -d social < ./scripts/"$(SCRIPT)"
