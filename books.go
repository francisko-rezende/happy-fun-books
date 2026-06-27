package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"sync"
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

func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	catalog := NewCatalog()

	err = json.NewDecoder(file).Decode(&catalog.data)
	if err != nil {
		return nil, err
	}
	catalog.Path = path

	return catalog, nil
}

type Catalog struct {
	mu   *sync.RWMutex
	data map[string]Book
	Path string
}

func (c *Catalog) GetAllBooks() []Book {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return slices.Collect(maps.Values(c.data))
}

func (c *Catalog) GetBook(bookID string) (Book, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	book, ok := c.data[bookID]

	return book, ok
}

func (c *Catalog) AddBook(book Book) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.data[book.BookID]
	if ok {
		return fmt.Errorf("id %q already exists", book.BookID)
	}

	c.data[book.BookID] = book
	return nil
}

func (c *Catalog) Sync() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	file, err := os.Create(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(c.data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Catalog) SetCopies(id string, copies int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	book, ok := c.data[id]
	if !ok {
		return fmt.Errorf("ID %q not found", id)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	c.data[id] = book
	return nil
}

func (c *Catalog) GetCopies(id string) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	book, ok := c.data[id]
	if !ok {
		return 0, fmt.Errorf("id %q not found", id)
	}

	return book.Copies, nil
}

func NewCatalog() *Catalog {
	return &Catalog{
		mu:   &sync.RWMutex{},
		data: map[string]Book{},
	}
}
