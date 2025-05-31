DB_URL=postgresql://root:secret@localhost:5432/user-service?sslmode=disable

network:
	docker network create dnd-network

postgres:
	docker run --name postgres12 --network dnd-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

restartpg:
	docker restart postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root user-db

dropdb:
	docker exec -it postgres12 dropdb user-db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateversion:
	migrate -path db/migration/ -database "$(DB_URL)" force 1

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateversion