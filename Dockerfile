FROM golang:latest

LABEL maintainer="Fuerback"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build

ENTRYPOINT ["./books-api"]