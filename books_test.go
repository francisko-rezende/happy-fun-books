package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)

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
	catalog := books.GetCatalog()

	got := books.GetAllBooks(catalog)
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

	catalog := books.GetCatalog()
	got, ok := books.GetBook(catalog, "a")

	if !ok {
		t.Fatal("book not found")
	}

	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	catalog := books.GetCatalog()
	_, ok := books.GetBook(catalog, "xyz")

	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsBookToCatalog(t *testing.T) {
	t.Parallel()
	newBookId := "c"
	catalog := books.GetCatalog()
	_, ok := books.GetBook(catalog, newBookId)

	if ok {
		t.Fatal("book already in catalog")
	}

	books.AddBook(catalog, books.Book{
		BookID: newBookId,
		Title:  "Test book title",
		Author: "Test book author",
		Copies: 3,
	})

	_, ok = books.GetBook(catalog, newBookId)

	if !ok {
		t.Fatal("Added book not found")
	}
}
