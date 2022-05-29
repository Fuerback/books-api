package domain

import (
	"context"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
	"github.com/Fuerback/books-api/internal/app/adapter/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreateNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name      string
		newBook   NewBook
		wantError error
	}{
		{
			name:      "only title",
			newBook:   NewBook{Title: "the book"},
			wantError: nil,
		},
		{
			name:      "title and author",
			newBook:   NewBook{Title: "the book", Author: "author"},
			wantError: nil,
		},
		{
			name:      "title, author and pages",
			newBook:   NewBook{Title: "the book", Author: "author", Pages: 1},
			wantError: nil,
		},
		{
			name:      "author too few letters",
			newBook:   NewBook{Title: "the book", Author: "au"},
			wantError: nil,
		},
		{
			name:      "empty request",
			newBook:   NewBook{},
			wantError: nil,
		},
		{
			name:      "title few letters",
			newBook:   NewBook{Title: "a"},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := mocks.NewMockRepoBook(mockCtrl)
			mockBook.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes()

			service := NewService(repository.NewBookRepository())
			err := service.Create(context.Background(), tt.newBook)

			if err != tt.wantError {
				t.Fatal(err.Error())
			}
		})
	}
}

func TestReadNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name      string
		bookID    string
		wantError error
	}{
		{
			name:      "number id",
			bookID:    "1",
			wantError: nil,
		},
		{
			name:      "empty id",
			bookID:    "",
			wantError: nil,
		},
		{
			name:      "text id",
			bookID:    "test",
			wantError: nil,
		},
		{
			name:      "uuid format",
			bookID:    "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := mocks.NewMockRepoBook(mockCtrl)
			mockBook.EXPECT().Read(gomock.Any(), gomock.Any()).AnyTimes().Return(repository.BookDetails{Title: "example"}, nil)

			service := NewService(repository.NewBookRepository())
			_, err := service.Read(context.Background(), tt.bookID)

			if err != tt.wantError {
				t.Fatal(err.Error())
			}
		})
	}
}

func TestUpdateNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name      string
		newBook   BookDetails
		wantError error
	}{
		{
			name:      "only title",
			newBook:   BookDetails{Title: "the book"},
			wantError: nil,
		},
		{
			name:      "title and author",
			newBook:   BookDetails{Title: "the book", Author: "author"},
			wantError: nil,
		},
		{
			name:      "title, author and pages",
			newBook:   BookDetails{Title: "the book", Author: "author", Pages: 1},
			wantError: nil,
		},
		{
			name:      "author too few letters",
			newBook:   BookDetails{Title: "the book", Author: "au"},
			wantError: nil,
		},
		{
			name:      "empty request",
			newBook:   BookDetails{},
			wantError: nil,
		},
		{
			name:      "title few letters",
			newBook:   BookDetails{Title: "a"},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := mocks.NewMockRepoBook(mockCtrl)
			mockBook.EXPECT().Update(gomock.Any(), gomock.Any()).AnyTimes()

			service := NewService(repository.NewBookRepository())
			err := service.Update(context.Background(), tt.newBook)

			if err != tt.wantError {
				t.Fatal(err.Error())
			}
		})
	}
}

func TestDeleteNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name      string
		bookID    string
		wantError error
	}{
		{
			name:      "number id",
			bookID:    "1",
			wantError: nil,
		},
		{
			name:      "empty id",
			bookID:    "",
			wantError: nil,
		},
		{
			name:      "text id",
			bookID:    "test",
			wantError: nil,
		},
		{
			name:      "uuid format",
			bookID:    "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := mocks.NewMockRepoBook(mockCtrl)
			mockBook.EXPECT().Delete(gomock.Any(), gomock.Any()).AnyTimes()

			service := NewService(repository.NewBookRepository())
			err := service.Delete(context.Background(), tt.bookID)

			if err != tt.wantError {
				t.Fatal(err.Error())
			}
		})
	}
}
