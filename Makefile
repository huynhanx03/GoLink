.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: build-gateway
build-gateway:
	envoy -c APIGateway/envoy.yaml

.PHONY: run-generation
run-generation:
	cd Generation && go run cmd/server/main.go

.PHONY: run-redirection
run-redirection:
	cd Redirection && go run cmd/server/main.go

.PHONY: run-identity
run-identity:
	cd Identity && go run cmd/server/main.go

.PHONY: init-db
init-db:
	docker exec -i scylla-node1 cqlsh < Generation/scripts/init_db.cql
	docker exec -i scylla-node1 cqlsh < Redirection/scripts/init_db.cql
	docker exec -i golink-postgres psql -U admin -d identity < Identity/scripts/seed.sql

.PHONY: init-cdc
init-cdc:
	./Redirection/scripts/init_cdc.sh

.PHONY: init-all
init-all:
	make init-db
	make init-cdc