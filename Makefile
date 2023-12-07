postgres:
	docker run --name pg-accounting -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pg-accounting createdb --username=root --owner=root db-accounting

dropdb:
	docker exec -it pg-accounting dropdb db-accounting

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/db-accounting?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/db-accounting?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test