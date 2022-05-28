package book

type bookService struct {
	// db
}

func NewService() Book {
	return &bookService{}
}

func (s *bookService) Create(book BookDetail) error {
	return nil
}

func (s *bookService) Read(bookID string) (error, BookDetail) {
	return nil, BookDetail{}
}

func (s *bookService) Update(book BookDetail) error {
	return nil
}

func (s *bookService) Delete(bookID string) error {
	return nil
}
