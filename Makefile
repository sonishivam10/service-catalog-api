.PHONY: setup stop logs

# Start app + db (both in containers)
setup:
	sudo docker-compose up --build

# Stop containers
stop:
	sudo docker-compose down -v

# View logs
logs:
	sudo docker-compose logs -f app

# Generate test token
token:
	go run scripts/generate_token.go

# Regenerate Swagger docs
swag:
	swag init   --generalInfo cmd/server/main.go   --output docs

#Run Go app locally (not in Docker)
run:
	go run cmd/server/main.go

# Run all tests (added only the service test)
test:
	go test ./internal/service