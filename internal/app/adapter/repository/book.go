package repository

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/infrastructure/db"
)

type bookRepository struct {
	dataBase db.DB
}

func NewBookRepository(dataBase db.DB) RepoBook {
	return &bookRepository{dataBase: dataBase}
}

func (r *bookRepository) Create(ctx context.Context, book NewBook) error { return nil }
func (r *bookRepository) Read(ctx context.Context, bookID string) (BookDetails, error) {
	return BookDetails{}, nil
}
func (r *bookRepository) Update(ctx context.Context, book BookDetails) error { return nil }
func (r *bookRepository) Delete(ctx context.Context, bookID string) error    { return nil }
