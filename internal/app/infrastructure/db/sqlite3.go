package db

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
)

type sqlite3 struct {
}

func NewSqlite3() DB {
	return &sqlite3{}
}

func (s *sqlite3) CreateNewBook(ctx context.Context, book repository.NewBook) (string, error) {
	return "", nil
}
func (s *sqlite3) FindBook(ctx context.Context, bookID string) (repository.BookDetails, error) {
	return repository.BookDetails{}, nil
}
func (s *sqlite3) UpdateBook(ctx context.Context, book repository.NewBook) error { return nil }
func (s *sqlite3) DeleteBook(ctx context.Context, bookID string) error           { return nil }