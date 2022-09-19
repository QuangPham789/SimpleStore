postgres:
	docker run --name postgres12-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12-alpine createdb --username=root --owner=root simple_store

dropdb:
	docker exec -it postgres12-alpine dropdb simple_store

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_store?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_store?sslmode=disable" -verbose down 

test:
	go test -v -cover ./...

sqlboiler:
	sqlboiler psql -c sqlboiler.toml --wipe --no-tests

.PHONY:	postgres createdb dropdb migrateup migratedown sqlboiler
