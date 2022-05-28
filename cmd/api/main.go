package main

import (
	"fmt"
	"github.com/Fuerback/books-api/internal/app/adapter/booksadapter"
	"github.com/Fuerback/books-api/internal/app/domain/book"
	"github.com/Fuerback/books-api/internal/app/infrastructure/server"
)

func main() {
	fmt.Println("Starting api server")

	// DB
	bookService := book.NewService()
	httpHandler := booksadapter.NewHttpHandler(bookService)

	server.NewHttpServer(httpHandler).Run()
}
