package domain

import "github.com/Fuerback/books-api/internal/app/adapter/repository"

type NewBook struct {
	Title  string
	Author string
	Pages  int
}

type BookDetails struct {
	ID     string
	Title  string
	Author string
	Pages  int
}

type UpdateBookDetails struct {
	Title  *string
	Author *string
	Pages  *int
}

type BooksFilter struct {
	Page    int
	PerPage int
	Title   string
	Author  string
}

func (u BooksFilter) newBookFiltersToDomain() repository.BooksFilter {
	return repository.BooksFilter{
		Page:    u.Page,
		PerPage: u.PerPage,
		Title:   u.Title,
		Author:  u.Author,
	}
}

func repoBooksToDomainBooks(books []repository.BookDetails) []BookDetails {
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

func (u NewBook) newBookToDomain() repository.NewBook {
	return repository.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}

func (u UpdateBookDetails) bookDetailsToDomain() repository.UpdateBookDetails {
	return repository.UpdateBookDetails{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
