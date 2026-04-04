package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	catalog := books.GetCatalog()

	if len(args) != 2 {
		fmt.Println("Usage: find <book id>")
		return
	}

	id := args[1]
	book, ok := books.GetBook(catalog, id)

	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the catalog")
		return
	}

	fmt.Println(books.BookToString(book))
}
