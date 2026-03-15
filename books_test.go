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

	got := books.GetAllBooks()

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

	got, ok := books.GetBook("a")

	if !ok {
		t.Fatal("book not found")
	}

	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()

	_, ok := books.GetBook("xyz")

	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}
