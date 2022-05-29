package repository

import (
	"github.com/Fuerback/books-api/internal/app/infrastructure/db"
)

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

func (u NewBook) newBookToDomain() db.NewBook {
	return db.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}

func (u BookDetails) bookDetailToDomain() db.BookDetails {
	return db.BookDetails{
		ID:     u.ID,
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
