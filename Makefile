init:
	go run github.com/99designs/gqlgen init
generate:
	go run github.com/99designs/gqlgen generate
postgresup:
	docker run --name hackernews-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:latest
postgresdown:
	docker kill hackernews-postgres
	docker rm hackernews-postgres