package test

import (
	"fmt"
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := mediator{}
	user := NewUser(&mediator, 1, "Ivan", "2222eeg@gmail.com", "2222", "2022-12-1", 0)
	order := NewOrd(&mediator, 1, "2022-12-21", "Pushkina 12", "??")
	mediator.AddUser(&user)
	mediator.AddOrder(&order)
	book1 := CardBook{
		ID:       1,
		Name:     "Конец вечности",
		Price:    "555.55",
		Discount: 5,
		Photo:    "default",
		Authors:  []string{"Айзек Азимов"},
		Amount:   5,
	}
	book2 := CardBook{
		ID:       2,
		Name:     "Город и звёзды",
		Price:    "777.55",
		Discount: 0,
		Photo:    "default",
		Authors:  []string{"Артур Кларк"},
		Amount:   100,
	}
	user.AddBook(book1)
	user.PrintBooks()
	fmt.Printf("Order contents: %v \n", order.books)
	user.ReplaceBooks([]CardBook{book2})
	user.PrintBooks()
	fmt.Printf("Order contents: %v\n", order.books)
}
