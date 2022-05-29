package handler

import (
	"context"
	"encoding/json"
	"github.com/Fuerback/books-api/internal/app/adapter/errors"
	"github.com/Fuerback/books-api/internal/app/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type httpHandler struct {
	bookService domain.Book
}

func NewHttpHandler(bookService domain.Book) BooksHandler {
	return &httpHandler{bookService: bookService}
}

func (c *httpHandler) Create(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	book := new(NewBook)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	log.Println("handling BookCreation")

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("error unmarshalling the request - " + err.Error()))
		log.Println("BookCreation - error unmarshalling the request")
		return
	}

	v := validator.New()
	err = v.Struct(book)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		var error errors.Error
		for _, err := range err.(validator.ValidationErrors) {
			message := "validation error on " + err.Namespace()
			error.Message = append(error.Message, message)
		}
		json.NewEncoder(resp).Encode(error)
		log.Println("BookCreation - error validating input")
		return
	}

	err = c.bookService.Create(ctx, book.newBookToDomain())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.NewError("error creating book - " + err.Error()))
		log.Println("BookCreation finished with StatusInternalServerError")
	}

	resp.WriteHeader(http.StatusCreated)
	log.Println("BookCreation finished")
}

func (c *httpHandler) Read(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	log.Println("handling BookReading")

	vars := mux.Vars(r)
	ID := vars["id"]
	if ID == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("Invalid id parameter"))
		return
	}

	book, err := c.bookService.Read(ctx, ID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.NewError("error reading book - " + err.Error()))
		log.Println("BookReading finished with StatusInternalServerError")
	}

	result, err := json.Marshal(book)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.NewError("error marshaling book result - " + err.Error()))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
	log.Println("BookReading finished")
}

func (c *httpHandler) Update(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	book := new(BookDetails)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	log.Println("handling BookUpdating")

	vars := mux.Vars(r)
	ID := vars["id"]
	if ID == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("Invalid id parameter"))
		return
	}
	book.ID = ID

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("error unmarshalling the request - " + err.Error()))
		log.Println("BookUpdating - error unmarshalling the request")
		return
	}

	v := validator.New()
	err = v.Struct(book)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		var error errors.Error
		for _, err := range err.(validator.ValidationErrors) {
			message := "validation error on " + err.Namespace()
			error.Message = append(error.Message, message)
		}
		json.NewEncoder(resp).Encode(error)
		log.Println("BookUpdating - error validating input")
		return
	}

	err = c.bookService.Update(ctx, book.bookDetailToDomain())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.NewError("error creating book - " + err.Error()))
		log.Println("BookUpdating finished with StatusInternalServerError")
	}

	resp.WriteHeader(http.StatusOK)
	log.Println("BookUpdating finished")
}

func (c *httpHandler) Delete(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	log.Println("handling BookRemoving")

	vars := mux.Vars(r)
	ID := vars["id"]
	if ID == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("Invalid id parameter"))
		return
	}

	err := c.bookService.Delete(ctx, ID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.NewError("error reading book - " + err.Error()))
		log.Println("BookRemoving finished with StatusInternalServerError")
	}

	resp.WriteHeader(http.StatusOK)
	log.Println("BookRemoving finished")
}
