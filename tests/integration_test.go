package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/Fuerback/books-api/internal/app/adapter/repository"
	"github.com/Fuerback/books-api/internal/app/adapter/web/handler"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TODO: Create error scenario tests

var (
	router      *mux.Router
	httpHandler handler.BooksHandler
)

func TestMain(m *testing.M) {
	dbName := "../skael_test_db"
	repo := repository.NewBookRepository(dbName)
	clearUp(dbName)
	bookService := domain.NewService(repo)
	httpHandler = handler.NewHttpHandler(bookService)

	os.Exit(m.Run())
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

func TestCreateBook(t *testing.T) {
	book := handler.NewBook{Title: "title", Author: "author", Pages: 3}
	body, _ := json.Marshal(book)
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, "/v1/books", requestReader)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/books", httpHandler.Create).Methods("POST")
	router.ServeHTTP(recorder, req)

	var bookID handler.BookID
	err = json.Unmarshal(recorder.Body.Bytes(), &bookID)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.NotEmpty(t, bookID.ID)
}

func TestReadBook(t *testing.T) {
	book := handler.NewBook{Title: "title", Author: "author", Pages: 3}
	body, _ := json.Marshal(book)
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, "/v1/books", requestReader)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/books", httpHandler.Create).Methods("POST")
	router.HandleFunc("/v1/books/{id}", httpHandler.Read).Methods("GET")
	router.ServeHTTP(recorder, req)

	var bookID handler.BookID
	err = json.Unmarshal(recorder.Body.Bytes(), &bookID)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.NotEmpty(t, bookID.ID)

	req, err = http.NewRequest(http.MethodGet, "/v1/books/"+bookID.ID, nil)
	assert.Nil(t, err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var bookDetails handler.BookDetails
	err = json.Unmarshal(recorder.Body.Bytes(), &bookDetails)
	assert.Nil(t, err)
	assert.Equal(t, bookID.ID, bookDetails.ID)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateBook(t *testing.T) {
	book := handler.NewBook{Title: "title", Author: "author", Pages: 3}
	body, _ := json.Marshal(book)
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, "/v1/books", requestReader)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/books", httpHandler.Create).Methods("POST")
	router.HandleFunc("/v1/books/{id}", httpHandler.Update).Methods("PATCH")
	router.ServeHTTP(recorder, req)

	var bookID handler.BookID
	err = json.Unmarshal(recorder.Body.Bytes(), &bookID)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.NotEmpty(t, bookID.ID)

	payload, _ := json.Marshal(handler.UpdateBookDetails{
		Title: func() *string {
			s := "title"
			return &s
		}(),
		Author: func() *string {
			s := "author"
			return &s
		}(),
	})
	req, err = http.NewRequest(http.MethodPatch, "/v1/books/"+bookID.ID, bytes.NewBuffer(payload))
	assert.Nil(t, err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteBook(t *testing.T) {
	book := handler.NewBook{Title: "title", Author: "author", Pages: 3}
	body, _ := json.Marshal(book)
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, "/v1/books", requestReader)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/books", httpHandler.Create).Methods("POST")
	router.HandleFunc("/v1/books/{id}", httpHandler.Delete).Methods("DELETE")
	router.HandleFunc("/v1/books/{id}", httpHandler.Read).Methods("GET")
	router.ServeHTTP(recorder, req)

	var bookID handler.BookID
	err = json.Unmarshal(recorder.Body.Bytes(), &bookID)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.NotEmpty(t, bookID.ID)

	req, err = http.NewRequest(http.MethodDelete, "/v1/books/"+bookID.ID, nil)
	assert.Nil(t, err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	req, err = http.NewRequest(http.MethodGet, "/v1/books/"+bookID.ID, nil)
	assert.Nil(t, err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
