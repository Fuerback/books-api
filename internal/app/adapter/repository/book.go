package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type bookRepository struct {
	dB     *sql.DB
	dbName string
}

func NewBookRepository(dbName string) RepoBook {
	repo := &bookRepository{dbName: dbName}
	repo.ClearUp()
	return repo
}

func (s *bookRepository) connectDatabase() error {
	db, err := sql.Open("mysql", s.dbName)
	if err != nil {
		return err
	}

	s.dB = db
	return nil
}

func (s *bookRepository) ClearUp() error {
	db, err := sql.Open("mysql", s.dbName)
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

	stmt, err := tx.PrepareContext(ctx, "insert into books(id, title, author, pages) values (?,?,?,?)")
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

func (s *bookRepository) ReadBooks(ctx context.Context, bookFilter BooksFilter) ([]BookDetails, error) {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return []BookDetails{}, err
	}

	filters := ""
	params := []interface{}{}
	if bookFilter.Title != "" {
		filters += "and title=\"" + bookFilter.Title + "\""
		params = append(params, bookFilter.Title)
	}
	if bookFilter.Author != "" {
		filters += "and author=\"" + bookFilter.Author + "\""
		params = append(params, bookFilter.Author)
	}

	params = append(params, bookFilter.PerPage)
	params = append(params, bookFilter.Page)

	query := fmt.Sprintf("select id, title, author, pages from books where deleted = 0 %s limit %d offset %d", filters, bookFilter.PerPage, bookFilter.Page)
	rows, err := s.dB.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return []BookDetails{}, err
	}

	var result []BookDetails
	for rows.Next() {
		book := BookDetails{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Pages)
		if err != nil {
			return nil, err
		}

		result = append(result, book)
	}
	return result, nil
}

func (s *bookRepository) Read(ctx context.Context, bookID string) (BookDetails, error) {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return BookDetails{}, err
	}

	stmt, err := s.dB.PrepareContext(ctx, "select id, title, author, pages from books where id = ? and deleted = 0")
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

func (s *bookRepository) Update(ctx context.Context, bookID string, book UpdateBookDetails) error {
	err := s.connectDatabase()
	defer s.dB.Close()
	if err != nil {
		return err
	}

	tx, err := s.dB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	fields, args := getFieldsAndParams(bookID, book)
	stmt, err := tx.PrepareContext(ctx, "update books set "+fields+" where id = ? and deleted = 0")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func getFieldsAndParams(bookID string, model interface{}) (string, []interface{}) {
	values := reflect.ValueOf(model)
	typeOfS := values.Type()
	updateFields := ""
	params := make([]interface{}, 0)

	for i := 0; i < values.NumField(); i++ {
		if !values.Field(i).IsNil() {
			updateFields += strings.ToLower(typeOfS.Field(i).Name) + " = ?,"
			params = append(params, values.Field(i).Interface())
		}
	}

	if updateFields != "" {
		// remove last comma
		updateFields = updateFields[:len(updateFields)-1]
	}

	params = append(params, bookID)

	return updateFields, params
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

	stmt, err := tx.PrepareContext(ctx, "update books set deleted = 1 where id = ?")
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
