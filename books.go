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

func GetCatalog() map[string]Book {
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

func GetBook(catalog map[string]Book, bookID string) (Book, bool) {
	book, ok := catalog[bookID]

	return book, ok
}

func GetAllBooks(catalog map[string]Book) []Book {
	return slices.Collect(maps.Values(catalog))
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func AddBook(catalog map[string]Book, book Book) {
	catalog[book.BookID] = book
}
