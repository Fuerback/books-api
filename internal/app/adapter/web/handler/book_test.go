package handler

import (
	"bytes"
	"encoding/json"
	"github.com/Fuerback/books-api/internal/app/domain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
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
			mockBook := mocks.NewMockBook(mockCtrl)
			mockBook.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes()

			handler := NewHttpHandler(mockBook)
			payload, _ := json.Marshal(tt.newBook)

			req, err := http.NewRequest(
				http.MethodPost,
				"/v1/books",
				bytes.NewBuffer(payload),
			)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/v1/books", handler.Create).Methods("POST")
			router.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatal(rr.Body.String())
			}
		})
	}
}
