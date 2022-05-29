package repository

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/domain"
)

type bookRepository struct {
	// db injection
}

func NewBookRepository() Book {
	return &bookRepository{}
}

func (r *bookRepository) Create(ctx context.Context, book domain.BookDetail) error { return nil }
func (r *bookRepository) Read(ctx context.Context, bookID string) (error, domain.BookDetail) {
	return nil, domain.BookDetail{}
}
func (r *bookRepository) Update(ctx context.Context, book domain.BookDetail) error { return nil }
func (r *bookRepository) Delete(ctx context.Context, bookID string) error          { return nil }
