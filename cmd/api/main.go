package main

import (
	"fmt"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
	"github.com/Fuerback/books-api/internal/app/adapter/web/handler"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/Fuerback/books-api/internal/app/infrastructure/db"
	"github.com/Fuerback/books-api/internal/app/infrastructure/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting api server")

	db := db.NewSqlite3()
	repository := repository.NewBookRepository(db)
	bookService := domain.NewService(repository)
	httpHandler := handler.NewHttpHandler(bookService)

	server.NewHttpServer(httpHandler).Run()
}
