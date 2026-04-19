package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: copies <book_id> <number_of_copies>")
		return
	}

	catalog, err := books.OpenCatalog("testdata/catalog")
	if err != nil {
		fmt.Printf("opening the catalog %v\n", err)
		return
	}

	bookID := args[1]
	copies, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = catalog.SetCopies(bookID, copies)
	if err != nil {
		fmt.Printf("Updating book: %v\n", err)
		return
	}

	err = catalog.Sync("testdata/catalog")
	if err != nil {
		fmt.Printf("Updating catalog file: %v\n", err)
		return
	}

	fmt.Printf("Updated book %v to %d copies", bookID, copies)
}
