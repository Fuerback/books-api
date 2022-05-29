package db

import (
	"context"
)

type DB interface {
	CreateNewBook(ctx context.Context, book NewBook) (string, error)
	FindBook(ctx context.Context, bookID string) (BookDetails, error)
	UpdateBook(ctx context.Context, book BookDetails) error
	DeleteBook(ctx context.Context, bookID string) error
}
