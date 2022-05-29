package repository

import (
	"context"
)

type bookRepository struct {
	// db injection
}

func NewBookRepository() Book {
	return &bookRepository{}
}

func (r *bookRepository) Create(ctx context.Context, book BookDetail) error { return nil }
func (r *bookRepository) Read(ctx context.Context, bookID string) (error, BookDetail) {
	return nil, BookDetail{}
}
func (r *bookRepository) Update(ctx context.Context, book BookDetail) error { return nil }
func (r *bookRepository) Delete(ctx context.Context, bookID string) error   { return nil }
