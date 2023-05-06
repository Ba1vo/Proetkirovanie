package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	opt := SearchOptions{
		Str: "Азимов",
	}
	Search := search{
		opt: opt,
	}
	Search.algorithm = &UserSearch{}
	books := Search.SearchBooks()
	fmt.Printf("Type %T. Value: %v\n", books, books)
	Search.algorithm = &StaffSearch{}
	books = Search.SearchBooks()
	fmt.Printf("Type %T. Value: %v\n", books, books)
}
