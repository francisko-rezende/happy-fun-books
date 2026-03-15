package books

import (
	"fmt"
)

type Book struct {
	BookID string
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		BookID: "a",
		Title:  "In the company of cheerful ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		BookID: "b",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func GetBook(bookID string) (Book, bool) {
	for _, book := range catalog {
		if book.BookID == bookID {
			return book, true
		}
	}

	return Book{}, false
}

func GetAllBooks() []Book {
	return catalog
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (%v copies)", book.Title, book.Author, book.Copies)
}
