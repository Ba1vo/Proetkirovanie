package test

import "fmt"

type user struct {
	mediator *mediator
	id       int
	name     string
	email    string
	pass     string
	date     string
	role     int
}

func NewUser(mediator *mediator, id int, name string, email string, pass string, date string, role int) user {
	u := user{mediator, id, name, email, pass, date, role}
	return u
}

func (u *user) AddBook(book CardBook) {
	u.mediator.AddBook(book)
}

func (u *user) ReplaceBooks(books []CardBook) {
	u.mediator.ReplaceBooks(books)
}

func (u *user) PrintBooks() { fmt.Println(u.mediator.OrderInfo()) }
