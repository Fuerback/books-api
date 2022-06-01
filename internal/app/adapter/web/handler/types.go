package handler

import "github.com/Fuerback/books-api/internal/app/domain"

type NewBook struct {
	Title  string `json:"title" validate:"required,gte=3"`
	Author string `json:"author" validate:"omitempty,gte=3"`
	Pages  int    `json:"pages" validate:"omitempty,gte=1"`
}

type BookID struct {
	ID string `json:"id"`
}

type BookDetails struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required,gte=3"`
	Author string `json:"author" validate:"omitempty,gte=3"`
	Pages  int    `json:"pages" validate:"omitempty,gte=1"`
}

type UpdateBookDetails struct {
	Title  *string `json:"title" validate:"omitempty,gte=3"`
	Author *string `json:"author" validate:"omitempty,gte=3"`
	Pages  *int    `json:"pages" validate:"omitempty,gte=1"`
}

type Books struct {
	Page    int           `json:"page"`
	PerPage int           `json:"perPage"`
	Items   []BookDetails `json:"items"`
}

func domainBooksToHandlerBooks(books []domain.BookDetails) []BookDetails {
	booksDetails := make([]BookDetails, 0)
	for _, b := range books {
		book := BookDetails{
			ID:     b.ID,
			Title:  b.Title,
			Author: b.Author,
			Pages:  b.Pages,
		}
		booksDetails = append(booksDetails, book)
	}
	return booksDetails
}

func (u NewBook) newBookToDomain() domain.NewBook {
	return domain.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}

func (u UpdateBookDetails) bookDetailToDomain() domain.UpdateBookDetails {
	return domain.UpdateBookDetails{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
