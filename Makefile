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

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Odvin/go-accounting-service/db/sqlc Store

migration_file:
	migrate create -ext sql -dir db/migration -seq $(name)

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migration_file proto