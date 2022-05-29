package domain

import "github.com/Fuerback/books-api/internal/app/adapter/repository"

type bookService struct {
	repo repository.Book
}

func NewService(repo repository.Book) Book {
	return &bookService{repo: repo}
}

func (s *bookService) Create(book BookDetail) error {
	return nil
}

func (s *bookService) Read(bookID string) (error, BookDetail) {
	return nil, BookDetail{}
}

func (s *bookService) Update(book BookDetail) error {
	return nil
}

func (s *bookService) Delete(bookID string) error {
	return nil
}
