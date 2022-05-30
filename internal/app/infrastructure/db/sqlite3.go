package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type sqlite3 struct {
	DB *sql.DB
}

func NewSqlite3() DB {
	return &sqlite3{}
}

func (s *sqlite3) CreateNewBook(ctx context.Context, book NewBook) (string, error) {
	err := s.connectDatabase()
	defer s.DB.Close()
	if err != nil {
		return "", err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
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
func (s *sqlite3) FindBook(ctx context.Context, bookID string) (BookDetails, error) {
	err := s.connectDatabase()
	defer s.DB.Close()
	if err != nil {
		return BookDetails{}, err
	}

	stmt, err := s.DB.PrepareContext(ctx, "select * from book where id = ? and deleted = 0")
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

func (s *sqlite3) UpdateBook(ctx context.Context, book BookDetails) error {
	err := s.connectDatabase()
	defer s.DB.Close()
	if err != nil {
		return err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

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

func (s *sqlite3) DeleteBook(ctx context.Context, bookID string) error {
	err := s.connectDatabase()
	defer s.DB.Close()
	if err != nil {
		return err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
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

func (s *sqlite3) connectDatabase() error {
	db, err := sql.Open("sqlite3", "./skael_db.db")
	if err != nil {
		return err
	}

	s.DB = db
	return nil
}
