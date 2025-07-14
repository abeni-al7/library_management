package services

import (
	"fmt"

	"github.com/abeni-al7/library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookId int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books map[int]models.Book
	members map[int]models.Member
}

func (library *Library) AddBook(book models.Book) {
	id := book.ID
	library.books[id] = book
}

func (library *Library) RemoveBook(bookID int) {
	delete(library.books, bookID)
}

func (library *Library) BorrowBook(bookId int, memberID int) error {
	book, ok := library.books[bookId]
	if !ok {
		return fmt.Errorf("book with id %v is not in the library", bookId)
	}
	_, exists := library.members[memberID]
	if !exists {
		return fmt.Errorf("member with id %v does not exist", memberID)
	}

	if book.Status == "Borrowed" {
		return fmt.Errorf("this book is already borrowed out")
	}

	book.Status = "Borrowed"
	member := library.members[memberID]
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	library.members[memberID] = member
	return nil
}

func (library *Library) ListAvailableBooks() []models.Book {
	var books []models.Book
	for _, value := range library.books {
		if value.Status == "Available" {
			books = append(books, value)
		}
	}
	return books
}

func (library *Library) ListBorrowedBooks(memberID int) []models.Book {
	_, exists := library.members[memberID]
	if !exists {
		return make([]models.Book, 0)
	}

	return library.members[memberID].BorrowedBooks
}