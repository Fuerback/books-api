package main

import (
	"fmt"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
	"github.com/Fuerback/books-api/internal/app/adapter/web/handler"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/Fuerback/books-api/internal/app/infrastructure/server"
)

func main() {
	fmt.Println("Starting api server")

	repository := repository.NewBookRepository()
	bookService := domain.NewService(repository)
	httpHandler := handler.NewHttpHandler(bookService)

	server.NewHttpServer(httpHandler).Run()
}
