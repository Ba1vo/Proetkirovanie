package proxy

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type book interface {
	info()
	getBookId() int
}

type fullBook struct {
	ID         int
	Name       string
	Price      string
	Discount   int
	ISBN       string
	Photo      string
	Desc       string
	Dimensions [3]int
	Authors    []string
	Publishers []string
	Genres     []string
	Amount     int
	Date       string
}

func (f fullBook) info() {
	fmt.Printf("%s written by %v, costs %s with %d%% discount\n", f.Name, f.Authors, f.Price, f.Discount)
}

func (f fullBook) getBookId() int { return f.ID }

func NewBook(id int) *fullBook {
	book, _ := GetBook(id)
	return &book
}

type proxyBook struct {
	id   int
	Book fullBook
}

func NewProxyBook(id int) *proxyBook {
	return &proxyBook{id: id, Book: fullBook{}}
}

func (p *proxyBook) info() {
	if (cmp.Equal(p.Book, fullBook{})) {
		p.Book = *NewBook(p.id)
	}
	p.Book.info()
}

func (p proxyBook) getBookId() int {
	return p.Book.ID
}
