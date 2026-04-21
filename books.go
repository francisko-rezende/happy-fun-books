package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
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

func (b Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	catalog := Catalog{}
	err = json.NewDecoder(file).Decode(&catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}

type Catalog map[string]Book

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

func (c Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(c)
	if err != nil {
		return err
	}
	return nil
}

func (c Catalog) SetCopies(id string, copies int) error {
	book, ok := c[id]
	if !ok {
		return fmt.Errorf("ID %q not found", id)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	c[id] = book
	return nil
}
