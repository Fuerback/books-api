package domain

import "context"

type Book interface {
	Create(ctx context.Context, book NewBook) (string, error)
	Read(ctx context.Context, bookID string) (BookDetails, error)
	Update(ctx context.Context, book BookDetails) error
	Delete(ctx context.Context, bookID string) error
}
