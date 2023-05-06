package factory

import (
	"fmt"
	"strings"
)

type writtenDoc interface {
	getAuthor() string
	getName() string
	getPrice() string
	getPhoto() string
	setAuthor(string)
	setName(string)
	setPrice(string)
	setDiscount(int)
	setPhoto(string)
}

type book struct {
	name     string
	authors  []string
	price    string
	discount int
	photo    string
}

func NewBook() *book {
	return &book{}
}

func (b *book) getAuthor() string {
	authors := strings.Join(b.authors, ", ")
	return authors
}

func (b *book) getName() string { return b.name }

func (b *book) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", b.price, b.discount)
}

func (b *book) getPhoto() string { return b.photo }

func (b *book) setAuthor(authors string) { b.authors = strings.Split(authors, ", ") }

func (b *book) setName(name string) { b.name = name }

func (b *book) setPrice(price string) { b.price = price }

func (b *book) setDiscount(disc int) { b.discount = disc }

func (b *book) setPhoto(photo string) { b.photo = photo }

type magazine struct {
	name      string
	publisher string
	price     string
	discount  int
	photo     string
}

func NewMagazine() *magazine { return &magazine{} }

func (m *magazine) getAuthor() string { return m.publisher }

func (m *magazine) getName() string { return m.name }

func (m *magazine) getPrice() string {
	return fmt.Sprintf("Price: %s with %d%% discount", m.price, m.discount)
}

func (m *magazine) getPhoto() string { return m.photo }

func (m *magazine) setAuthor(author string) { m.publisher = author }

func (m *magazine) setName(name string) { m.name = name }

func (m *magazine) setPrice(price string) { m.price = price }

func (m *magazine) setDiscount(disc int) { m.discount = disc }

func (m *magazine) setPhoto(photo string) { m.photo = photo }
