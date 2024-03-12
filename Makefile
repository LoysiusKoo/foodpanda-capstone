.PHONY: start-and-migrate
start-and-migrate: stop-and-delete networkcreate postgres-image postgres wait-for-postgres createdb migrateup

.PHONY: stop-and-delete
stop-and-delete:
	-docker stop postgres16
	-docker rm postgres16
	-docker network rm playlist-network

.PHONY: networkcreate
networkcreate:
	-docker network create playlist-network

.PHONY: postgres-image
postgres-image:
	docker pull postgres:16-alpine

.PHONY: postgres
postgres:
	docker run --name postgres16 --network playlist-network -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:16-alpine

.PHONY: createdb
createdb: wait-for-postgres
	docker exec -it postgres16 createdb --username=postgres --owner=postgres foodpanda-playlist

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: server
server:
	go run main.go

.PHONY: migrateup
migrateup:
	migrate -path database/schema -database "postgresql://postgres:password@127.0.0.1:5433/foodpanda-playlist?sslmode=disable" -verbose up

wait-for-postgres:
	sleep 5  # wait time to run make createdb