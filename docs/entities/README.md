1. Init migration schema

install golang-migrate

brew install golang-migrate

create the folder /db/migration

migrate create -ext sql -dir db/migration -seq init_schema

2. Generate functions for sql queries

install sqlc

brew install sqlc

sqlc init

configure sqlc.yaml
