package main

import (
	"fmt"
	"github.com/abeni-al7/library_management/controllers"
)

func Choices() {
	fmt.Println("Here are the things you can do: ")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow out a book")
	fmt.Println("4. Return a borrowed book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List books borrowed by a member")
	fmt.Println("7. Register new member")
	fmt.Println("8. Quit")
}

func main() {
	fmt.Println("Welcome to the Library Management System")
	fmt.Println("----------------------------------------")

	choice := 0

	for choice != -1 {
		Choices()
		fmt.Scanln(&choice)

		switch {
		case choice == 1:
			println()
			controllers.AddNewBook()
			println()
		case choice == 2:
			println()
			controllers.RemoveExistingBook()
			println()
		case choice == 3:
			println()
			controllers.BorrowBook()
			println()
		case choice == 4:
			println()
			controllers.ReturnBook()
			println()
		case choice == 5:
			println()
			controllers.ListAvailableBooks()
			println()
		case choice == 6:
			println()
			controllers.ListBorrowedBooks()
			println()
		case choice == 7:
			println()
			controllers.RegisterMember()
			println()
		case choice == 8:
			println()
			controllers.ListMembers()
		case choice == 8:
			fmt.Println("Goodbye!")
			choice = -1
		}
	}
}