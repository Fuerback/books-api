package repository

import (
	"context"
)

type bookRepository struct {
	// db injection
}

func NewBookRepository() RepoBook {
	return &bookRepository{}
}

func (r *bookRepository) Create(ctx context.Context, book NewBook) error { return nil }
func (r *bookRepository) Read(ctx context.Context, bookID string) (BookDetails, error) {
	return BookDetails{}, nil
}
func (r *bookRepository) Update(ctx context.Context, book BookDetails) error { return nil }
func (r *bookRepository) Delete(ctx context.Context, bookID string) error    { return nil }
