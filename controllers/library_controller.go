package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abeni-al7/library_management/models"
	"github.com/abeni-al7/library_management/services"
)

var (
	library services.Library
)

func AddNewBook() {
	var title string
	var author string

	id := len(library.Books)
	status := "Available"

	fmt.Println()
	fmt.Println("-----------------")
	fmt.Println("Fill the necessary information for the book: ")

	fmt.Println("What is the title of the book?")
	titleReader := bufio.NewReader(os.Stdin)
	title, _ = titleReader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Who is the author of the book?")
	authorReader := bufio.NewReader(os.Stdin)
	author, _ = authorReader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{ID: id, Title: title, Author: author, Status: status}
	library.AddBook(book)
}

func RemoveExistingBook() {
	var id int

	fmt.Println()
	fmt.Println("-----------------")
	fmt.Println("Enter the ID of the book to remove: ")

	fmt.Scanln(&id)
	library.RemoveBook(id)
}

func BorrowBook() {
	var bookId int
	var memberId int

	fmt.Println("Enter the id of the borrower: ")
	fmt.Scanln(&memberId)
	fmt.Println("Enter the id of the book they want to borrow: ")
	fmt.Scanln(&bookId)
	err := library.BorrowBook(bookId, memberId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The book with id %v was borrowed by the member with id %v", bookId, memberId)
		fmt.Println()
	}
}

func ReturnBook() {
	var bookId int
	var memberId int

	fmt.Println("Enter the id of the borrower: ")
	fmt.Scanln(&memberId)
	fmt.Println("Enter the id of the book to be returned: ")
	fmt.Scanln(&bookId)
	err := library.ReturnBook(bookId, memberId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The book with id %v has been returned by the member with id %v", bookId, memberId)
		fmt.Println()
	}
}

func ListAvailableBooks() {
	books := library.ListAvailableBooks()
	fmt.Println("Here are the available books to borrow: ")
	for _, value := range books {
		fmt.Printf("%v - %v BY %v", value.ID, value.Title, value.Author)
		fmt.Println()
	}
}

func ListBorrowedBooks() {
	var memberId int
	fmt.Println("Enter the Id of the member whose borrowed books you want to see: ")
	fmt.Scanln(&memberId)

	books := library.ListBorrowedBooks(memberId)
	fmt.Println("Here are the books borrowed: ")
	for _, value := range books {
		fmt.Printf("%v - %v BY %v", value.ID, value.Title, value.Author)
		fmt.Println()
	}
}