package abstract

import (
	"fmt"
	"strings"
)

type bnj struct {
}

func NewBNJ() bnj {
	return bnj{}
}

func (a *bnj) printDoc() writtenDoc {
	return &bnjBook{}
}

func (a *bnj) printPoster() poster {
	return &bnjPoster{}
}

type bnjBook struct {
	name     string
	authors  []string
	price    string
	discount int
	photo    string
}

type bnjPoster struct {
	price    string
	discount int
	photo    string
}

func (b *bnjBook) getAuthor() string {
	authors := strings.Join(b.authors, ", ")
	return authors
}

func (b *bnjBook) getName() string {
	return b.name
}

func (b *bnjBook) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", b.price, b.discount)
}

func (b *bnjBook) getPhoto() string {
	return b.photo
}

func (b *bnjBook) getPublisher() string {
	return "Boka i Joka"
}

func (b *bnjBook) setAuthor(authors string) {
	b.authors = strings.Split(authors, ", ")
}

func (b *bnjBook) setName(name string) {
	b.name = name
}

func (b *bnjBook) setPrice(price string) {
	b.price = price
}

func (b *bnjBook) setDiscount(disc int) {
	b.discount = disc
}

func (b *bnjBook) setPhoto(photo string) {
	b.photo = photo
}

func (m *bnjPoster) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", m.price, m.discount)
}

func (m *bnjPoster) getPhoto() string {
	return m.photo
}

func (m *bnjPoster) getSize() string {
	return "7x7x7"
}

func (m *bnjPoster) getPublisher() string {
	return "Boka i Joka"
}

func (m *bnjPoster) setPrice(price string) {
	m.price = price
}
func (m *bnjPoster) setDiscount(disc int) {
	m.discount = disc
}

func (m *bnjPoster) setPhoto(photo string) {
	m.photo = photo
}
