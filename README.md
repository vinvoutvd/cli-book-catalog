# CLI Book Catalog

A simple command-line application written in Go for managing a personal book collection.

## Features

- **List books** — display all books in the catalog.
- **Add books** — add a new book with its title, author, and publication year.
- **Search books** — find a book by its title.
- **Delete books** — remove a book from the catalog.
- **JSON storage** — save and load the catalog using a local JSON file.

## Project Structure

```text
CLI_Book_Catalog/
├── catalog/
│   ├── book.go
│   └── storage.go
├── books.json
├── go.mod
├── main.go
└── .gitignore
```

## Requirements

- Go 1.27.3 or later

## Getting Started

Clone the repository:

```bash
git clone https://github.com/vinvoutvd/ClI_Book_Catalog.git
```

Navigate to the project directory:

```bash
cd ClI_Book_Catalog
```

Run the application:

```bash
go run .
```

## Data Storage

The application stores book information in a local `books.json` file.

Each book contains:

- Name
- Author
- Year

## Technologies

- Go
- JSON
- Go standard library

## Purpose

This project was created as a practical exercise in Go programming.

It demonstrates working with:
- Structs and slices
- Packages and modules
- File operations
- JSON encoding and decoding
- Error handling
- Command-line input and output
