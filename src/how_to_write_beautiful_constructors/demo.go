package main

type Book struct {
	Title string
	ISBN  string
}

type Option func(*Book)

func NewBook_(options ...Option) *Book {
	b := &Book{}
	for _, option := range options {
		option(b)
	}
	return b
}

func WithTitle(title string) Option {
	return func(b *Book) {
		b.Title = title
	}
}

func WithISBN(isbn string) Option {
	return func(b *Book) {
		b.ISBN = isbn
	}
}

func main() {
	_ = NewBook_(
		WithTitle("GO book"),
		WithISBN("123456"),
	)
}
