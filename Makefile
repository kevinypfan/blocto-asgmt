DB_URL=postgresql://root:secret@localhost:5432/blocto-asgmt?sslmode=disable

network:
	docker network create blocto-asgmt

postgres:
	docker run --name postgres --network blocto-asgmt -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root blocto-asgmt

dropdb:
	docker exec -it postgres dropdb blocto-asgmt

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlcgen:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown test server
