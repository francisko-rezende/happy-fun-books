package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() *books.Catalog {
	catalog := books.NewCatalog()
	err := catalog.AddBook(books.Book{
		BookID: "a",
		Title:  "In the company of cheerful ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	})
	if err != nil {
		panic(err)
	}

	err = catalog.AddBook(books.Book{
		BookID: "b",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	})
	if err != nil {
		panic(err)
	}

	return catalog
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
	catalog := getTestCatalog()
	got := catalog.GetAllBooks()
	assertTestBooks(t, got)
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

	err := catalog.AddBook(books.Book{
		BookID: newBookId,
		Title:  "Test book title",
		Author: "Test book author",
		Copies: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, ok = catalog.GetBook(newBookId)

	if !ok {
		t.Fatal("Added book not found")
	}
}

func TestAddBook_DoesntAllowAddingABookWithAnExtantId(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("a")
	if !ok {
		t.Fatalf("book with id %q not found", "a")
	}
	input := books.Book{
		BookID: "a",
		Title:  "New book",
		Author: "New author",
		Copies: 5,
	}

	err := catalog.AddBook(input)
	if err == nil {
		t.Error("expect error when trying to add a book with an id that is already present in the catalog")
	}
}

func TestSetCopies_IsRaceFree(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	go func() {
		for range 100 {
			err := catalog.SetCopies("a", 0)
			if err != nil {
				panic(err)
			}
		}
	}()

	for range 100 {
		_, err := catalog.GetCopies("a")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSetCopies_AddsTheReceivedNumberToABook(t *testing.T) {
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

func TestSetCopies_MethodReturnsErrorIfItCopiesNegative(t *testing.T) {
	t.Parallel()
	book := books.Book{}
	err := book.SetCopies(-10)
	if err == nil {
		t.Fatal("want error for negative copies, got nil")
	}
}

func TestOpenCatalog_MethodReturnsErrorIfEmptyAddressIsProvided(t *testing.T) {
	t.Parallel()
	_, err := books.OpenCatalog("")
	if err == nil {
		t.Fatal("want error for empty string provided to catalog, got nil")
	}
}

func TestOpenCatalog_ReadsSameDataWrittenBySync(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	catalog.Path = t.TempDir() + "/catalog"
	err := catalog.Sync()
	if err != nil {
		t.Fatal(err)
	}
	newCatalog, err := books.OpenCatalog(catalog.Path)
	if err != nil {
		t.Fatal(err)
	}
	got := newCatalog.GetAllBooks()
	assertTestBooks(t, got)
}

func TestCatalogSetCopies_SetsTheNumberOfCopiesTheBookWithTheReceivedID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	book, ok := catalog.GetBook("a")
	if !ok {
		t.Fatalf("book with id %v not found", "a")
	}
	if book.Copies != 1 {
		t.Fatalf("want book to have 1 copy before set copies call, got %v", book.Copies)
	}
	err := catalog.SetCopies("a", 51)
	if err != nil {
		t.Fatal(err)
	}
	book, ok = catalog.GetBook("a")
	if !ok {
		t.Fatalf("book with id %v not found", "a")
	}
	if 51 != book.Copies {
		t.Fatalf("want %v copies after change, got %v", 51, book.Copies)
	}
}

func TestNewCatalog_ReturnsANewCatalog(t *testing.T) {
	t.Parallel()
	catalog := books.NewCatalog()
	got := len(catalog.GetAllBooks())
	want := 0

	if got != want {
		t.Fatalf("want catalog to be empty and thus have zero length, got length %v", got)
	}
}

func assertTestBooks(t *testing.T, got []books.Book) {
	t.Helper()

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

	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}
