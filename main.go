package main

import (
	"fmt"

	"github.com/Fuerback/books-api/internal/app/adapter/repository"
	"github.com/Fuerback/books-api/internal/app/adapter/web/handler"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/Fuerback/books-api/internal/app/infrastructure/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting api server")

	repository := repository.NewBookRepository("root:books_db@tcp(db-mysql:3306)/books_db?charset=utf8mb4&parseTime=True&loc=Local")
	bookService := domain.NewService(repository)
	httpHandler := handler.NewHttpHandler(bookService)

	server.NewHttpServer(httpHandler).Run()
}
