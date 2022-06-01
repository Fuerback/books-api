package handler

import (
	"bytes"
	"encoding/json"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name           string
		newBook        NewBook
		wantStatusCode int
	}{
		{
			name:           "only title",
			newBook:        NewBook{Title: "the book"},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:           "title and author",
			newBook:        NewBook{Title: "the book", Author: "author"},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:           "title, author and pages",
			newBook:        NewBook{Title: "the book", Author: "author", Pages: 1},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:           "author too few letters",
			newBook:        NewBook{Title: "the book", Author: "au"},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "empty request",
			newBook:        NewBook{},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "title few letters",
			newBook:        NewBook{Title: "a"},
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := domain.NewMockBook(mockCtrl)
			mockBook.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)
			payload, _ := json.Marshal(tt.newBook)

			req, err := http.NewRequest(
				http.MethodPost,
				"/v1/books",
				bytes.NewBuffer(payload),
			)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books", handler.Create).Methods("POST")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
		})
	}
}

func TestReadBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name           string
		bookID         string
		wantStatusCode int
	}{
		{
			name:           "number 1 ID",
			bookID:         "1",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "test ID",
			bookID:         "test",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "empty ID",
			bookID:         "",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "UUID format",
			bookID:         "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := domain.NewMockBook(mockCtrl)
			mockBook.EXPECT().Read(gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)

			req, err := http.NewRequest(
				http.MethodGet,
				"/v1/books/"+tt.bookID,
				nil,
			)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books/{id}", handler.Read).Methods("GET")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
		})
	}
}

func TestReadBooks_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name           string
		params         string
		wantStatusCode int
	}{
		{
			name:           "page and perPage",
			params:         "?page=0&perPage=10",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "title and author",
			params:         "?title=test&author=test",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "page, perPage, title and author",
			params:         "?page=0&perPage=10&title=test&author=test",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "empty",
			params:         "",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "invalid",
			params:         "unknown",
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := domain.NewMockBook(mockCtrl)
			mockBook.EXPECT().ReadBooks(gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)

			req, err := http.NewRequest(
				http.MethodGet,
				"/v1/books"+tt.params,
				nil,
			)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books", handler.ReadBooks).Methods("GET")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
		})
	}
}

func TestUpdateNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name           string
		bookID         string
		book           UpdateBookDetails
		wantStatusCode int
	}{
		{
			name:   "number 1 ID and complete book",
			bookID: "1",
			book: UpdateBookDetails{
				Title: func() *string {
					s := "title"
					return &s
				}(),
				Author: func() *string {
					s := "author"
					return &s
				}(),
				Pages: func() *int {
					s := 1
					return &s
				}(),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "test ID and only title",
			bookID: "test",
			book: UpdateBookDetails{
				Title: func() *string {
					s := "title"
					return &s
				}(),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "empty ID and title",
			bookID: "",
			book: UpdateBookDetails{
				Title: func() *string {
					s := "title"
					return &s
				}(),
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:   "UUID format and incomplete author",
			bookID: "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			book: UpdateBookDetails{
				Title: func() *string {
					s := "title"
					return &s
				}(),
				Author: func() *string {
					s := "au"
					return &s
				}(),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "UUID format and empty book",
			bookID:         "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			book:           UpdateBookDetails{},
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := domain.NewMockBook(mockCtrl)
			mockBook.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)
			payload, _ := json.Marshal(tt.book)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/v1/books/"+tt.bookID,
				bytes.NewBuffer(payload),
			)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books/{id}", handler.Update).Methods("PATCH")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
		})
	}
}

func TestDeleteNewBook_Table(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name           string
		bookID         string
		wantStatusCode int
	}{
		{
			name:           "number 1 ID",
			bookID:         "1",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "test ID",
			bookID:         "test",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "empty ID",
			bookID:         "",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "UUID format",
			bookID:         "837dcd08-26d7-4886-9a2e-c9827a6d68f0",
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := domain.NewMockBook(mockCtrl)
			mockBook.EXPECT().Delete(gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/v1/books/"+tt.bookID,
				nil,
			)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books/{id}", handler.Delete).Methods("DELETE")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
		})
	}
}
