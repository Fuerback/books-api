package domain

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
)

type bookService struct {
	repo repository.Book
}

func NewService(repo repository.Book) Book {
	return &bookService{repo: repo}
}

func (s *bookService) Create(ctx context.Context, book NewBook) error {
	return nil
}

func (s *bookService) Read(ctx context.Context, bookID string) (error, BookDetail) {
	return nil, BookDetail{}
}

func (s *bookService) Update(ctx context.Context, book BookDetail) error {
	return nil
}

func (s *bookService) Delete(ctx context.Context, bookID string) error {
	return nil
}
