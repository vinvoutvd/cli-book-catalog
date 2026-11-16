// Package catalog предоставляет функции для работы с книжным каталогом.
package catalog

import (
	"strings"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func SearchBook(books []Book, name string) (Book, bool) {
	name = strings.TrimSpace(name)
	for _, book := range books {
		if strings.EqualFold(book.Name, name) {
			return book, true
		}
	}
	return Book{}, false
}

func DeleteBook(books []Book, number int) ([]Book, bool) {
	if len(books) == 0 || number < 1 || number > len(books) {
		return books, false
	}
	books = append(books[:number-1], books[number:]...)
	return books, true
}

func AddBook(books []Book, book Book) []Book {
	books = append(books, book)
	return books
}
