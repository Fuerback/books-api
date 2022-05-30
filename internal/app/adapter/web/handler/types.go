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
	ID     string `json:"_"`
	Title  string `json:"title" validate:"required,gte=3"`
	Author string `json:"author" validate:"omitempty,gte=3"`
	Pages  int    `json:"pages" validate:"omitempty,gte=1"`
}

func (u NewBook) newBookToDomain() domain.NewBook {
	return domain.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}

func (u BookDetails) bookDetailToDomain() domain.BookDetails {
	return domain.BookDetails{
		ID:     u.ID,
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
