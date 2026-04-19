package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: find <book id>")
		return
	}

	catalog, err := books.OpenCatalog("testdata/catalog")
	if err != nil {
		fmt.Printf("opening catalog %v\n", err)
		return
	}

	ID := args[1]
	book, ok := catalog.GetBook(ID)

	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the catalog")
		return
	}

	fmt.Println(book)
}
