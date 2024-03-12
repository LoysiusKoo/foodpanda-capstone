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

.PHONY: sqlc
sqlc:
	sqlc generate

wait-for-postgres:
	sleep 5  # wait time to run make createdb