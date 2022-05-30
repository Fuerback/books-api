package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type bookRepository struct {
	dB     *sql.DB
	dbName string
}

func NewBookRepository(dbName string) RepoBook {
	return &bookRepository{dbName: dbName}
}

func (s *bookRepository) connectDatabase() error {
	db, err := sql.Open("sqlite3", s.dbName)
	if err != nil {
		return err
	}

	s.dB = db
	return nil
}

func (s *bookRepository) ClearUp() error {
	db, err := sql.Open("sqlite3", s.dbName)
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

func (s *bookRepository) Create(ctx context.Context, book NewBook) (string, error) {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return "", err
	}

	tx, err := s.dB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	stmt, err := tx.PrepareContext(ctx, "insert into book(id, title, author, pages) values (?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return "", err
	}

	bookID := uuid.NewString()
	_, err = stmt.ExecContext(ctx, bookID, book.Title, book.Author, book.Pages)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return bookID, nil
}

func (s *bookRepository) Read(ctx context.Context, bookID string) (BookDetails, error) {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return BookDetails{}, err
	}

	stmt, err := s.dB.PrepareContext(ctx, "select id, title, author, pages from book where id = ? and deleted = 0")
	defer stmt.Close()
	if err != nil {
		return BookDetails{}, err
	}

	book := BookDetails{}
	err = stmt.QueryRow(bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Pages)
	if err != nil {
		return BookDetails{}, err
	}

	return book, nil
}

func (s *bookRepository) Update(ctx context.Context, book BookDetails) error {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return err
	}

	tx, err := s.dB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// TODO: update only sent fields
	stmt, err := tx.PrepareContext(ctx, "update book set title = ?, author = ?, pages = ? where id = ? and deleted = 0")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, book.Title, book.Author, book.Pages, book.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (s *bookRepository) Delete(ctx context.Context, bookID string) error {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return err
	}

	tx, err := s.dB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "update book set deleted = 1 where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, bookID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
