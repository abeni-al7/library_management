# Library Management System
## By: Abenezer Alebachew Endalew, July 14, 2025

## Table of Contents

## Introduction
This app is useful to manage library workflows such as keeping track of books, members and borrowing information. The scope of this project is to allow the Library manager to Add, remove, borrow and return books, register members and see all the information related to books and members on the console. This app is going to be used by a Librarian. It is built with Go 1.24.5 using the standard library of Go exclusively without any third party packages.

## System Requirements
- OS: any Operating system can be used
- Go version 1.24.5 or higher
- No external libraries required
- Run with a single command

## Installation and Setup

1. Clone the repository from Github
```
git clone https://github.com/abeni-al7/library_management.git
```

2. Get into the application directory
```
cd library_management
```

3. Run the app
```
go run main.go
```

4. Use the command Line Interface to guide you throughout the application.


## Features
- Store author, title and status of a Book
- Store Name and Borrow list of a member
- Add new books to the library system
- Remove books from the library system
- Register new members
- List Books available for borrow
- List Books that a member has borrowed
- List members of the library
- Robust error handling without crashing the app

## Key Structs
```
type Book struct {
	ID int
	Title string
	Author string
	Status string
}
```

```
type Member struct {
	ID int
	Name string
	BorrowedBooks []Book
}
```

```
type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
}
```

## File structure
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod

## Limitations
- No database support (in-memory) and volatile storage
- Only supports a single user