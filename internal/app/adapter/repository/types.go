package repository

type NewBook struct {
	Title  string
	Author string
	Pages  int
}

type BookDetails struct {
	ID     string
	Title  string
	Author string
	Pages  int
}

type UpdateBookDetails struct {
	Title  *string
	Author *string
	Pages  *int
}

type BooksFilter struct {
	Page    int
	PerPage int
	Title   string
	Author  string
}
