package services

import (
	"fmt"

	"github.com/abeni-al7/library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
}

func (library *Library) AddBook(book models.Book) {
	id := book.ID
	library.Books[id] = book
}

func (library *Library) RemoveBook(bookID int) {
	delete(library.Books, bookID)
}

func (library *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := library.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id %v is not in the library", bookID)
	}

	if book.Status == "Borrowed" {
		return fmt.Errorf("this book is already borrowed out")
	}
	_, exists := library.Members[memberID]
	if !exists {
		return fmt.Errorf("member with id %v does not exist", memberID)
	}

	book.Status = "Borrowed"
	member := library.Members[memberID]
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	library.Members[memberID] = member
	return nil
}

func (library *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := library.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id %v is not in the library", bookID)
	}

	if book.Status == "Available" {
		return fmt.Errorf("this book was not borrowed out in the first place")
	}
	_, exists := library.Members[memberID]
	if !exists {
		return fmt.Errorf("member with id %v does not exist", memberID)
	}

	book.Status = "Available"
	member := library.Members[memberID]
	for i, v := range member.BorrowedBooks {
		if v == book {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
		}
	}
	library.Members[memberID] = member
	return nil
}

func (library *Library) ListAvailableBooks() []models.Book {
	var books []models.Book
	for _, value := range library.Books {
		if value.Status == "Available" {
			books = append(books, value)
		}
	}
	return books
}

func (library *Library) ListBorrowedBooks(memberID int) []models.Book {
	_, exists := library.Members[memberID]
	if !exists {
		return make([]models.Book, 0)
	}

	return library.Members[memberID].BorrowedBooks
}