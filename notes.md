# The deeper love of go

## 1. Happy fun books

The chapter starts with this program

```go
package main

import "fmt"

func main() {
 fmt.Println("Hello, world")
}
```

### Packages

Code in Go belongs to a package, which we can see by the first line of code. _Packages_ are a way of grouping related code.
Packages can have any name but _main_ is special: it produces an executable when compiled.

### Imports

We can import packages like so `import "fmt"`.

### Functions

Functions are declared using the `func` keyword.
The _main_ func is special (kinda like the main package): it's the function where the execution of your program will likely start.

### Calling a function

We call functions like so `fmt.Println("Hello, world")`. This is a _statement_, meaning that this tells go to do something (unlike the import _declaration_ we saw previously, which gives descriptive information about the program).

### Function ordering

Function ordering doesn't matter, you don't have to define a function before calling it. Order them as you see fit. The book recommended putting the most important, commonly-used functions first

### Variables

A variable is a piece of data that might change later (as opposed to constants, which shouldn't change). A variable is a way of giving a name to a piece of data so it's easier to refer to it again later, for example. We define variables like so:

```go
var book = "'Master and commander', by Patrick O'Brian"
```

A variable like the one above is called a _literal_ . A literal value is a value that is directly written in the code as opposed to being derived from a function call or another kind of computation.

### Assigning values to variables

After creating a variable we can reassign its value like so

```go
var hi = "hello"
hi = "hi"
```

### Declaring function parameters

We can declare function parameters like so

```go
func printBook(title, author string) {}
```

Since both `title` and `author` are strings, we are allowed to declare their types at the end of the parameter list

## 2. Just my type

### The `int` type

We can use the `int` type to deal with numbers in go

### Variable declarations

We can declare a variable and assign a value to it straight away

```go
var title = "The city"
```

Or we can create an empty variable

```go
var title string
```

### Zero value and default values

Variables declared without a value get a default empty value, like 0 for ints and "" for strings

### Declaring a new struct type

We can declare new structs like so

```go
type book struct {
  Title string
  Author string
  Copies int
}
```

### Struct variables and values

We can declare an empty struct like so

```go
var book Book
```

And we can declare a filled struct like so

```go
var book = Book{
  Title: "Sea Room",
  Author: "Adam Nicolson",
  Copies: 2,
}
```

We can also assign variables straight away using the walrus operator

```go
book := Book{
  Title: "Sea Room",
  Author: "Adam Nicolson",
  Copies: 2,
}
```

### Accessing struct fields

We can access struct fields by using the dot notation `book.Title`

## 3. Test of strength

### The `fmt.Sprintf` function

Go offers different ways to deal with strings. `fmt.Printf` allows you pass parameters to build a string and then print it while `fmt.Sprintf` also allows parameters but returns the final string instead of printing it.

### The `testing` package

Go provides a package with utils for building self testing function (functions for testing the functionality of other functions)

#### Declaring a test

We can write tests in go like so

```go

package main

import "testing"

func TestBookToString_FormatsBookInfoASString(t *testing.T) {
 input := Book{
  Title:  "Sea Room",
  Author: "Adam Nicolson",
  Copies: 2,
 }

 want := "Sea Room, by Adam Nicolson - 2 copies"
 got := BookToString(input)

 if want != got {
  t.Fatal("BookToString: wrong result")
 }
}
```

To run this test, we need to create a module in our app

#### Creating a module

A _module_ is a group of packages that live together.
With the module in place, we can start our tests using `go test`

#### Interpreting the failure output

When printing formatted strings, we can use `%q` to print out strings between double quotes

#### The naming of tests

We can run `go run github.com/bitfield/gotestdox/cmd/gotestdox@latest` in our project to get prettier output for our test results as long as name them as `TestName_ExpectedBehavior`

### Creating a package

#### `main` is not importable

Besides being the only package that can be compiled into an executable binary file, `main` _can't_ be imported

#### Updating the `main` package

You can't have two different packages in the same folder, so we created a new package for our main file in `/cmd/list`. `cmd` is a convention used by go developers to keep packages that are executable.

## 4. Slicing & dicing

### Slices

Collections of a given type of data are called slices

#### Slice variables and literals

We can initialize slices like so

```go
// create an empty slice of book
var books []Book

// assign a value to the variable
books = []Book{}

// assign a book with an item
books = []Book{
  {
    Title: "Chasing the Thrill",
    Author: "Daniel Barbarisi"
  }
}
```

#### Indexing slices

Trying to access an index out of bounds results in a panic

We can modify a given item of a slice like so

```go
books[0] = Book{
  Title: "The Power Broker",
  Author: "Robert A. Caro"
}
```

We can modify a specific field of a struct in a map like so

```go
books[0].Title = "The Power Broker"
```

#### The `len` and `append` functions

We can get the length of a slice using the `len` function

```go
fmt.Println(len(books))
// output: 1
```

We can add new elements to a slice like so:

```go
newBook := Book{
  Title: "Skyjack",
  Author: "Geoffrey Gray"
}

// returns a new slice with the new item
books = append(books, newBook)
```

### A collection of books

#### A global `catalog` variable

The _scope_ of a variable is the part of the program that is allowed to refer to said variable.
