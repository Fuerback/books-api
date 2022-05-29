package repository

type NewBook struct {
	Title  string
	Author string
	Pages  int
}

type BookDetail struct {
	ID     string
	Title  string
	Author string
	Pages  int
}
