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

func (u NewBook) newBookToDomain() repository.NewBook {
	return repository.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}

func (u BookDetails) bookDetailsToDomain() repository.BookDetails {
	return repository.BookDetails{
		ID:     u.ID,
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
