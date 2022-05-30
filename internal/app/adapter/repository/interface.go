package repository

import (
	"context"
)

type RepoBook interface {
	Create(ctx context.Context, book NewBook) (string, error)
	Read(ctx context.Context, bookID string) (BookDetails, error)
	ReadBooks(ctx context.Context, bookFilter BooksFilter) ([]BookDetails, error)
	Update(ctx context.Context, book BookDetails) error
	Delete(ctx context.Context, bookID string) error
}
