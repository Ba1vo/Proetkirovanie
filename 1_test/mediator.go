package test

import "fmt"

type mediator struct {
	order *order
	user  *user
}

func (m *mediator) AddBook(book CardBook) {
	books := m.order.GetBooks()
	books = append(books, book)
	m.order.SetBooks(books)
}

func (m *mediator) ReplaceBooks(books []CardBook) {
	m.order.SetBooks(books)
}

func (m *mediator) AddUser(user *user) {
	m.user = user
}

func (m *mediator) AddOrder(order *order) {
	m.order = order
}

func (m *mediator) OrderInfo() string {
	return fmt.Sprintf("User %s ordered following books: %v", m.user.name, m.order.books)
}
