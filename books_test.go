package books_test

import (
	"books"
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

	want := "Sea Room by Adam Nicolson (2 copies)"
	got := books.BookToString(input)

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
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

	got := books.GetAllBooks()

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsBookByBookID(t *testing.T) {
	t.Parallel()
	want := books.Book{
		BookID: 1,
		Title:  "In the company of cheerful ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}

	got, ok := books.GetBook(1)

	if !ok {
		t.Fatal("book not found")
	}

	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	_, ok := books.GetBook(999)

	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}
