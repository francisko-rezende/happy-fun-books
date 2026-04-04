package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() books.Catalog {
	return books.Catalog{
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

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := input.String()

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
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
	catalog := getTestCatalog()

	got := catalog.GetAllBooks()
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsBookByBookID(t *testing.T) {
	t.Parallel()
	want := books.Book{
		BookID: "a",
		Title:  "In the company of cheerful ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}

	catalog := getTestCatalog()
	got, ok := catalog.GetBook("a")

	if !ok {
		t.Fatal("book not found")
	}

	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()
	_, ok := catalog.GetBook("xyz")

	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsBookToCatalog(t *testing.T) {
	t.Parallel()
	newBookId := "c"
	catalog := getTestCatalog()
	_, ok := catalog.GetBook(newBookId)

	if ok {
		t.Fatal("book already in catalog")
	}

	catalog.AddBook(books.Book{
		BookID: newBookId,
		Title:  "Test book title",
		Author: "Test book author",
		Copies: 3,
	})

	_, ok = catalog.GetBook(newBookId)

	if !ok {
		t.Fatal("Added book not found")
	}
}

func TestAddCopies_AddsTheReceivedNumberToABook(t *testing.T) {
	t.Parallel()

	catalog := getTestCatalog()
	newCopiesValue := 42
	book, _ := catalog.GetBook("a")

	err := book.SetCopies(newCopiesValue)

	want := newCopiesValue
	got := book.Copies

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v copies, got %v", want, got)
	}
}

func TestAddCopies_MethodReturnsErrorIfItCopiesNegative(t *testing.T) {
	t.Parallel()
	book := books.Book{}
	err := book.SetCopies(-10)
	if err == nil {
		t.Fatal("want error for negative copies, got nil")
	}
}
