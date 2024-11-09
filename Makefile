run-db-container:
	docker run --name=todo-db -e POSTGRES_PASSWORD=qweR1234 -p 5432:5432 -d --rm postgres

run-dev-server:
	air

run-migrate:
	migrate -path ./schema -database 'postgres://postgres:qweR1234@localhost:5432/postgres?sslmode=disable' up

run-drop-db:
	migrate -path ./schema -database 'postgres://postgres:qweR1234@localhost:5432/postgres?sslmode=disable' down

gen-docs:
	swag init -g ./cmd/todo/main.go -o cmd/docs

gen-mocks:
	mockery

run-tests: get-mocks
	go test ./... -v

lint:
	golangci-lint run