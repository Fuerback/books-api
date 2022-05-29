package domain

import "context"

type Book interface {
	Create(ctx context.Context, book NewBook) error
	Read(ctx context.Context, bookID string) (error, BookDetail)
	Update(ctx context.Context, book BookDetail) error
	Delete(ctx context.Context, bookID string) error
}
