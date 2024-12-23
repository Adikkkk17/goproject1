package Library

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) {
	l.books[book.ID] = book
	fmt.Println("Книга добавлена:", book.Title)
}
func (l *Library) BorrowBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Книга с таким айдишником не найдена")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Книга уже взята")
		return
	}
	book.IsBorrowed = true
	l.books[id] = book // важно перезаписать обновлённую структуру обратно в map
	fmt.Println("Книга взята:", book.Title)
}
func (l *Library) ReturnBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Книга с таким айдишником не найдена")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Книга и так не взята возвращать не нужно")
		return
	}
	book.IsBorrowed = false
	l.books[id] = book
	fmt.Println("Книга возвращена:", book.Title)
}

func (l *Library) ListBooks() {
	fmt.Println("Список даступных книг(IsBorrowed = false):")
	for _, book := range l.books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Название: %s, Автор: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
