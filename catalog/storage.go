package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func GetCatalog() ([]Book, error) {
	fileName := "books.json"
	file, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return []Book{}, nil
	}
	if err != nil {
		return []Book{}, fmt.Errorf("Не удалось прочитать файл каталога: %w", err)
	}
	var books []Book
	if err = json.Unmarshal(file, &books); err != nil {
		return []Book{}, fmt.Errorf("Не удалось считать JSON: %w", err)
	}
	return books, nil
}

func SaveCatalog(books []Book) error {
	jsonBooks, err := json.Marshal(books)
	if err != nil {
		return fmt.Errorf("не удалось сериализовать каталог. Ошибка: %w", err)
	}
	if err := os.WriteFile("books.json", jsonBooks, 0644); err != nil {
		return fmt.Errorf("не удалось записать каталог в файл. Ошибка: %w", err)
	}
	return nil
}
