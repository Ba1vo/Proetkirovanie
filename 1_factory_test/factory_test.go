package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	book := NewItem("book")
	magazine := NewItem("magazine")
	booklet := NewItem("booklet")
	fmt.Printf("Type: %T\n", book)
	fmt.Printf("Type: %T\n", magazine)
	fmt.Printf("Type: %T\n", booklet)
}
