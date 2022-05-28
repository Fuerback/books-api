package book

type Book interface {
	Create(book BookDetail) error
	Read(bookID string) (error, BookDetail)
	Update(book BookDetail) error
	Delete(bookID string) error
}
