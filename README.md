# Books-api

The project is a CRUD Golang REST API that uses gorilla/mux as router framework, SQLite as a database and Docker to build.

## TODO

- Improve error handling
- Work with docker-compose
- Create a gRpc for internal usage
- Add audit columns in DB
- Create more unit and integration tests to cover more scenarios
- Work with idempotent routes

# Getting started

## Locally

To run project locally you need to export the environment variable `PORT=:8000` or to other port to your choice and. 
You must have installed:

- Golang >=1.17
- SQLite3

```shell
make run-local
```

```shell
make run-tests
```

## Docker

```shell
make build
make docker-image
make docker-run
```

or just

```shell
make docker-up
```

The default PORT to test via Docker is `8080`

to run **tests on Docker**:

```shell
make docker-tests
```

## Documentation

The API documentation is on [./docs/openapi/swagger.yaml](https://github.com/Fuerback/books-api/blob/main/docs/openapi/swagger.yaml)
The database schema is on [./docs/db/schema_v1.sql](https://github.com/Fuerback/books-api/blob/main/docs/db/schema_v1.sql)