package db

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
)

type DB interface {
	CreateNewBook(ctx context.Context, book repository.NewBook) (string, error)
	FindBook(ctx context.Context, bookID string) (repository.BookDetails, error)
	UpdateBook(ctx context.Context, book repository.NewBook) error
	DeleteBook(ctx context.Context, bookID string) error
}
