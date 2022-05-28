package main

import (
	"fmt"
	"github.com/Fuerback/books-api/internal/app/adapter/booksadapter"
	"github.com/Fuerback/books-api/internal/app/infrastructure/server"
)

func main() {
	fmt.Println("Starting api server")

	httpHandler := booksadapter.NewHttpHandler()

	server.NewHttpServer(httpHandler).Run()
}
