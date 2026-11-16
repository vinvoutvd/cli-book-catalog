package main

import (
	"book-catalog/catalog"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//Вывод действий для каталога
	for {
		fmt.Println("=== Книжный каталог ===")
		fmt.Println("1. Показать все книги")
		fmt.Println("2. Добавить книгу")
		fmt.Println("3. Удалить книгу")
		fmt.Println("4. Найти книгу")
		fmt.Println("5. Выйти")
		fmt.Println("Выберите действие:")
		if !scanner.Scan() {
			log.Fatal("Не удалось считать ввод пользователя", scanner.Err())
		}
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Println("Функция для показа каталога")
			showCatalog()
		case "2":
			fmt.Println("Функция для добавления книги")
			fmt.Println("Введите название книги:")
			newBook := catalog.Book{}
			var err error
			if !scanner.Scan() {
				log.Fatal("Не удалось считать название книги ", scanner.Err())
			}
			if strings.TrimSpace(scanner.Text()) == "" {
				log.Fatal("Вы не ввели название книги !")
			}
			newBook.Name = strings.TrimSpace(scanner.Text())
			fmt.Println("Введите автора книги:")
			if !scanner.Scan() {
				log.Fatal("Не удалось считать автора книги", scanner.Err())
			}
			if strings.TrimSpace(scanner.Text()) == "" {
				log.Fatal("Вы не ввели автора!")
			}
			newBook.Author = strings.TrimSpace(scanner.Text())
			fmt.Println("Введите год издания книги:")
			if !scanner.Scan() {
				log.Fatal("Не удалось считать год издания книги", scanner.Err())
			}
			newBook.Year, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				log.Fatal(err)
			}
			if newBook.Year <= 0 {
				log.Fatal("Год должен быть целым числом отличным от 0")
			}

			books, err := catalog.GetCatalog()
			if err != nil {
				fmt.Println(err)
				continue
			}
			books = catalog.AddBook(books, newBook)
			if err := catalog.SaveCatalog(books); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Сохранение произошло успешно!")
		case "3":
			fmt.Println("Функция для удаления книги")
			fmt.Println("Введите номер книги")
			var number int
			if !scanner.Scan() {
				log.Fatal("Не удалось считать ввод пользователя", scanner.Err())
			}
			number, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Ошибка: Нужно ввести число больше 0")
				continue
			}
			books, err := catalog.GetCatalog()
			if err != nil {
				fmt.Println(err)
				continue
			}
			books, ok := catalog.DeleteBook(books, number)
			if !ok {
				fmt.Println("Не удалось удалить книгу")
				continue
			}
			if err := catalog.SaveCatalog(books); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Книга удалена")
		case "4":
			fmt.Println("Функция для поиска книги")
			fmt.Println("Введите имя книги :")
			var name string
			if !scanner.Scan() {
				log.Fatal("Не удалось считать ввод пользователя", scanner.Err())
			}
			name = scanner.Text()
			books, err := catalog.GetCatalog()
			if err != nil {
				fmt.Println(err)
				continue
			}
			book, found := catalog.SearchBook(books, name)
			if !found {
				fmt.Println("Книга не найдена")
				continue
			}
			fmt.Printf("Ваша книга: %s, %s, %d\n", book.Name, book.Author, book.Year)
		case "5":
			fmt.Println("Функция для выхода из каталога")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func showCatalog() {
	books, err := catalog.GetCatalog()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(books) == 0 {
		fmt.Println("Каталог пуст")
		return
	}
	for index, book := range books {
		fmt.Printf("%d - Title: %s , Author: %s , Publishing Year: %d\n", index+1,
			book.Name, book.Author, book.Year)
	}
}
