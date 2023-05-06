package abstract

import (
	"fmt"
	"strings"
)

type ast struct {
}

func NewAST() ast { return ast{} }

func (a *ast) printDoc() writtenDoc { return &astBook{} }

func (a *ast) printPoster() poster { return &astPoster{} }

type astBook struct {
	name     string
	authors  []string
	price    string
	discount int
	photo    string
}

type astPoster struct {
	price    string
	discount int
	photo    string
}

func (b *astBook) getAuthor() string {
	authors := strings.Join(b.authors, ", ")
	return authors
}

func (b *astBook) getName() string { return b.name }

func (b *astBook) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", b.price, b.discount)
}

func (b *astBook) getPhoto() string { return b.photo }

func (b *astBook) getPublisher() string { return "ACT" }

func (b *astBook) setAuthor(authors string) { b.authors = strings.Split(authors, ", ") }

func (b *astBook) setName(name string) { b.name = name }

func (b *astBook) setPrice(price string) { b.price = price }

func (b *astBook) setDiscount(disc int) { b.discount = disc }

func (b *astBook) setPhoto(photo string) { b.photo = photo }

func (m *astPoster) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", m.price, m.discount)
}

func (m *astPoster) getPhoto() string { return m.photo }

func (m *astPoster) getSize() string { return "10x20x10" }

func (m *astPoster) getPublisher() string { return "ACT" }

func (m *astPoster) setPrice(price string) { m.price = price }

func (m *astPoster) setDiscount(disc int) { m.discount = disc }

func (m *astPoster) setPhoto(photo string) { m.photo = photo }
