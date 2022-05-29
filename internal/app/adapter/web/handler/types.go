package handler

import "github.com/Fuerback/books-api/internal/app/domain"

type NewBook struct {
	Title  string `json:"title" validate:"required,gte=3"`
	Author string `json:"author" validate:"omitempty,gte=3"`
	Pages  int    `json:"pages" validate:"omitempty,gte=1"`
}

type BookDetails struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (u NewBook) newBookToDomain() domain.NewBook {
	return domain.NewBook{
		Title:  u.Title,
		Author: u.Author,
		Pages:  u.Pages,
	}
}
