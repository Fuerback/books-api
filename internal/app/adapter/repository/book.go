package repository

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/infrastructure/db"
)

type bookRepository struct {
	dataBase db.DB
}

func NewBookRepository(dataBase db.DB) RepoBook {
	return &bookRepository{dataBase: dataBase}
}

func (r *bookRepository) Create(ctx context.Context, book NewBook) (string, error) {
	return r.dataBase.CreateNewBook(ctx, book.newBookToDomain())
}

func (r *bookRepository) Read(ctx context.Context, bookID string) (BookDetails, error) {
	book, err := r.dataBase.FindBook(ctx, bookID)
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

func (r *bookRepository) Update(ctx context.Context, book BookDetails) error {
	return r.dataBase.UpdateBook(ctx, book.bookDetailToDomain())
}

func (r *bookRepository) Delete(ctx context.Context, bookID string) error {
	return r.dataBase.DeleteBook(ctx, bookID)
}
