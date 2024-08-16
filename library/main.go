package main

import (
	"books/books"
	"fmt"
)

func main() {
	fmt.Println("library books")

	var book1 books.Book = books.Book{
		Name:   "Maincraft",
		Author: "Arnold Hitman",
		Year:   1946,
	}
	var book2 books.Book = books.Book{
		Name:   "Capital",
		Author: "Joe Marx",
		Year:   1999,
	}
	var book2delete string = "Capital"
	books.AddBook(book1)
	books.AddBook(book2)
	books.DeleteBook(book2delete)
	fmt.Println(books.BookArr)
}
