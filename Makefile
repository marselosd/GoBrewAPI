postgres:
	docker run --name gobrew1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpwd -d postgres:17

createdb:
	docker exec -it gobrew1 createdb --username=root --owner=root go_brew

dropdb:
	docker exec -it gobrew1 dropdb go_brew

migrateup:
	migrate -path db/migration -database postgresql://root:secretpwd@localhost:5432/go_brew?sslmode=disable -verbose up

migratedown:
	migrate -path db/migration -database postgresql://root:secretpwd@localhost:5432/go_brew?sslmode=disable -verbose down

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, sqlc, test
