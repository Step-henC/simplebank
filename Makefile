postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine
createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown: 
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

psqlcli:
	docker exec -it postgres16 psql -U root -d simple_bank
test: 
	go test -v -cover ./...
dropdb: 
	docker exec -it postgres16 dropdb simple_bank

server:
	go run main.go

.PHONY: postgres dropdb createdb migrateup migratedown sqlc psqlcli server