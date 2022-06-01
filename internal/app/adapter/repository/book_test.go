package repository

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	repo RepoBook
)

// TODO: create more scenarios

func init() {
	dbName := "../../../../skael_test_db"
	repo = NewBookRepository(dbName)
	clearUp(dbName)
}

func TestCreateBook(t *testing.T) {
	ID, err := repo.Create(context.Background(), NewBook{Title: "title"})

	assert.NoError(t, err)
	assert.NotEmpty(t, ID)
}

func TestFindBook(t *testing.T) {
	ID, err := repo.Create(context.Background(), NewBook{Title: "title"})
	assert.NoError(t, err)

	book, err := repo.Read(context.Background(), ID)
	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestFindBooks(t *testing.T) {
	_, err := repo.ReadBooks(context.Background(), BooksFilter{
		Page:    0,
		PerPage: 10,
		Title:   "title",
		Author:  "author",
	})
	assert.NoError(t, err)
}

func TestUpdateBook(t *testing.T) {
	ID, err := repo.Create(context.Background(), NewBook{Title: "title"})
	assert.NoError(t, err)

	err = repo.Update(context.Background(), ID, UpdateBookDetails{
		Title: func() *string {
			s := "title 2"
			return &s
		}(),
		Author: func() *string {
			s := "author"
			return &s
		}(),
		Pages: func() *int {
			s := 3
			return &s
		}(),
	})
	assert.NoError(t, err)

	book, err := repo.Read(context.Background(), ID)
	assert.NoError(t, err)
	assert.Equal(t, "title 2", book.Title)
	assert.Equal(t, "author", book.Author)
	assert.Equal(t, 3, book.Pages)
}

func TestDeleteBook(t *testing.T) {
	ID, err := repo.Create(context.Background(), NewBook{Title: "title"})
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), ID)
	assert.NoError(t, err)

	_, err = repo.Read(context.Background(), ID)
	assert.NotNil(t, err)
}

func clearUp(dbName string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from book")
	tx.Commit()
	return err
}
