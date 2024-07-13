# Books-api

The project is a CRUD Golang REST API that uses gorilla/mux as router framework, MySQL as a database and Docker to build.

## TODO

- Improve error handling
- Create a gRpc for internal usage
- Add audit columns in DB
- Use testcontainers in e2e tests: example https://eltonminetto.dev/en/post/2024-02-15-using-test-helpers/
- Work with idempotent routes

# Getting started

## Locally

To run project locally you need to export the environment variable `PORT=:8080` or to other port to your choice and. 
You must have installed:

- Golang >=1.17
- MySQL >=5.7

```shell
make run-tests
```

## Docker Compose

```shell
make docker-compose
```

The default PORT to test via Docker is `8080`

## Documentation

The API documentation is on [./docs/openapi/swagger.yaml](https://github.com/Fuerback/books-api/blob/main/docs/openapi/swagger.yaml)
The database schema is on [./docs/db/schema_v1.sql](https://github.com/Fuerback/books-api/blob/main/docs/db/schema_v1.sql)