build:
	cd cmd/server && go build -o ../../bin/micro

run: build
	./bin/micro

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

migrate_up:
	migrate -source file:./internal/migrations -database cassandra://localhost:9042/smudge up

migrate_down:
	migrate -source file:./internal/migrations -database cassandra://localhost:9042/smudge down

migrate_create:
	migrate create --ext cql -dir internal/migrations -tz utc $(filter-out $@, $(MAKECMDGOALS))