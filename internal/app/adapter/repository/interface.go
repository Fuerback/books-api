package repository

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/domain"
)

type Book interface {
	Create(ctx context.Context, book domain.BookDetail) error
	Read(ctx context.Context, bookID string) (error, domain.BookDetail)
	Update(ctx context.Context, book domain.BookDetail) error
	Delete(ctx context.Context, bookID string) error
}
