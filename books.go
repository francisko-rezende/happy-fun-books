package books

import (
	"fmt"
)

type Book struct {
	BookID int
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		BookID: 1,
		Title:  "In the company of cheerful ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		BookID: 2,
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func GetBook(bookID int) (Book, bool) {
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
