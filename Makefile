build:
	@echo Building...
	@go build -o ./bin/bifrost cmd/main.go
	@echo Building completed...

run:build
	@echo Starting...
	@./bin/bifrost

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

clean:
	@echo Cleaning...
	@rm -rf ./bin
	@echo Cleaning completed...