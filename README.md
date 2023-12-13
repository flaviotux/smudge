# Smudge API

Golang todo application with GraphQL, gRPC and RestFul interfaces.

## How to

### Install

To install dependencies is recomended to use [goenv-shell](https://github.com/go-nv/goenv)

```shell
goenv install 1.21.4
```

```shell
go mod tidy
```

### Build

```shell
go build -v -o ./bin/server ./cmd/server...
```

### Binary Run

```shell
./bin/server
```

## Usage

### Migrations

To run migrations it's used the [migrate cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

#### Create

```shell
migrate create --ext cql -dir cmd/migrate/migrations -tz utc create_todo_table
```

#### Up

```shell
migrate -source file:./migrations -database cassandra://localhost:9042/smudge up
```

#### Down

```shell
migrate -source file:./migrations -database cassandra://localhost:9042/smudge down
```

### Local

#### Run

```shell
go run ./cmd/server/main.go
```
