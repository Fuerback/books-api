package repository

import (
	"context"
)

type Book interface {
	Create(ctx context.Context, book BookDetail) error
	Read(ctx context.Context, bookID string) (error, BookDetail)
	Update(ctx context.Context, book BookDetail) error
	Delete(ctx context.Context, bookID string) error
}
