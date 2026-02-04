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

.PHONY: run-billing
run-billing:
	cd Billing && go run cmd/server/main.go

.PHONY: run-orchestrator
run-orchestrator:
	cd Orchestrator && go run cmd/server/main.go

.PHONY: gen-ent-identity
gen-ent-identity:
	go generate ./Identity/...

.PHONY: gen-ent-billing
gen-ent-billing:
	go generate ./Billing/...

.PHONY: gen-ent
gen-ent:
	make gen-ent-identity
	make gen-ent-billing

.PHONY: init-db-generation
init-db-generation:
	docker exec -i scylla-node1 cqlsh < Generation/scripts/init_db.cql

.PHONY: init-db-redirection
init-db-redirection:
	docker exec -i scylla-node1 cqlsh < Redirection/scripts/init_db.cql

.PHONY: init-db-identity
init-db-identity:
	docker exec -i golink-postgres psql -U admin -d postgres < Identity/scripts/seed.sql

.PHONY: init-db-billing
init-db-billing:
	docker exec -i golink-postgres psql -U admin -d postgres < Billing/scripts/seed.sql

.PHONY: init-db
init-db:
	make init-db-generation
	make init-db-redirection
	make init-db-identity
	make init-db-billing

.PHONY: init-cdc
init-cdc:
	./Redirection/scripts/init_cdc.sh

.PHONY: init-all
init-all:
	make init-db
	make init-cdc

.PHONY: install-tools
install-tools:
	@echo "Installing tools..."
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: gen-proto
gen:
	cd proto && docker-compose up --build