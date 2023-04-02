package queries

import (
	"fmt"
	"testing"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func TestSearch(t *testing.T) {
	options := decoder.SearchOptions{
		Str: "Азимов",
	}
	testBooks := []decoder.FullBook{
		{ID: 0,
			Name:       "Конец вечности",
			Price:      "555.55",
			Discount:   5,
			ISBN:       "978-5-04-106735-9",
			Photo:      "default",
			Desc:       "bla-bla",
			Dimensions: [3]int{10, 20, 10},
			Authors:    []string{"А. Азимов"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фантастика"},
			Amount:     5,
			Date:       "2020-10-5",
		},
	}
	for _, book := range testBooks {
		addBook(book)
	}
	books, err := GetSearchBooks(options)
	if err != nil {
		fmt.Printf("Options: %#v , result: %s", options, err.Error())
		return
	}
	fmt.Printf("Options: %#v , result: okay, found books:", options)
	var i int
	for _, book := range books {
		fmt.Printf("\nBook %d: %d, %s, %s %#v", i, book.ID, book.Name, book.Price, book.Authors)
		i++
		deleteBook(book.ID)
	}
}
