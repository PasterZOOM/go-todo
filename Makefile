run-db-container:
	docker run --name=todo-db -e POSTGRES_PASSWORD=qweR1234 -p 5432:5432 -d --rm postgres

run-dev-server:
	air