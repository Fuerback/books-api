package domain

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
)

type bookService struct {
	repo repository.RepoBook
}

func NewService(repo repository.RepoBook) Book {
	return &bookService{repo: repo}
}

func (s *bookService) Create(ctx context.Context, book NewBook) (string, error) {
	return s.repo.Create(ctx, book.newBookToDomain())
}

func (s *bookService) ReadBooks(ctx context.Context, bookFilter BooksFilter) ([]BookDetails, error) {
	books, err := s.repo.ReadBooks(ctx, bookFilter.newBookFiltersToDomain())
	if err != nil {
		return nil, err
	}

	return repoBooksToDomainBooks(books), nil
}

func (s *bookService) Read(ctx context.Context, bookID string) (BookDetails, error) {
	book, err := s.repo.Read(ctx, bookID)
	if err != nil {
		return BookDetails{}, err
	}

	bookDetails := BookDetails{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Pages:  book.Pages,
	}
	return bookDetails, nil
}

func (s *bookService) Update(ctx context.Context, bookID string, book UpdateBookDetails) error {
	return s.repo.Update(ctx, bookID, book.bookDetailsToDomain())
}

func (s *bookService) Delete(ctx context.Context, bookID string) error {
	return s.repo.Delete(ctx, bookID)
}
