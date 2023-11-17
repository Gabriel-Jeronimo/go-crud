# Gin CRUD API for Books

This API provides basic CRUD functionality for managing books using the Gin web framework in Go.

## Description

The API allows users to perform essential operations on a collection of books. It supports creating new books, retrieving all books, and fetching individual books by their unique identifiers.

## Features

- **Get All Books:** Retrieve a list of all available books.
- **Get Book by ID:** Retrieve a specific book by its unique identifier.
- **Create Book:** Add a new book to the collection.

## Getting Started

### Prerequisites

- Go installed on your machine ([Go Installation Guide](https://golang.org/doc/install))
- Gin framework installed (`go get -u github.com/gin-gonic/gin`)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Gabriel-Jeronimo/go-crud.git
```

2. Navigate to the project directory:

```bash
cd go-crud
```

### Usage

1. Run the application:

```bash
go run main.go
```

Access the API endpoints using tools like cURL, Postman, or any web browser at http://localhost:8080

API Endpoints

    GET /books: Retrieve all books.
    GET /books/{id}: Retrieve a book by ID.
    POST /books: Create a new book.
