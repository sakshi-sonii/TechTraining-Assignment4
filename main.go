package main

import (
    "fmt"
)

type Book struct {
    ID     int
    Title  string
    Author string
    Status string // "Available" or "CheckedOut"
}

type Library struct {
    Books []Book
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book Book) {
    l.Books = append(l.Books, book)
}

// BorrowBook marks a book as checked out
func (l *Library) BorrowBook(bookID int) error {
    for i, book := range l.Books {
        if book.ID == bookID && book.Status == "Available" {
            l.Books[i].Status = "CheckedOut"
            return nil
        }
    }
    return fmt.Errorf("book not available")
}

// ReturnBook marks a book as available
func (l *Library) ReturnBook(bookID int) error {
    for i, book := range l.Books {
        if book.ID == bookID && book.Status == "CheckedOut" {
            l.Books[i].Status = "Available"
            return nil
        }
    }
    return fmt.Errorf("book not found or not checked out")
}

type LibraryManager interface {
    AddBook(book Book)
    BorrowBook(bookID int) error
    ReturnBook(bookID int) error
}

func main() {
    library := &Library{}
    
    library.AddBook(Book{ID: 1, Title: "1984", Author: "George Orwell", Status: "Available"})
    library.AddBook(Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Available"})

    fmt.Println("Books in library:", library.Books)
    
    err := library.BorrowBook(1)
    if err != nil {
        fmt.Println("Error borrowing book:", err)
    }
    
    fmt.Println("Books in library after borrowing:", library.Books)
    
    err = library.ReturnBook(1)
    if err != nil {
        fmt.Println("Error returning book:", err)
    }
    
    fmt.Println("Books in library after returning:", library.Books)
}
