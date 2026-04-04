package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	BookID string
	Title  string
	Author string
	Copies int
}

func (b *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}

	b.Copies = copies
	return nil
}

type Catalog map[string]Book

func GetCatalog() Catalog {
	return map[string]Book{
		"a": {
			BookID: "a",
			Title:  "In the company of cheerful ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		"b": {
			BookID: "b",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
}

func (c Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(c))
}

func (c Catalog) GetBook(bookID string) (Book, bool) {
	book, ok := c[bookID]

	return book, ok
}

func (c Catalog) AddBook(book Book) {
	c[book.BookID] = book
}

func (b Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", b.Title, b.Author, b.Copies)
}
