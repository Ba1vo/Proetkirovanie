package queries

import (
	"fmt"
	"testing"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func TestSearch(t *testing.T) {
	options := decoder.SearchOptions{
		Str: "Айзек",
	}
	testBooks := []decoder.FullBook{
		/*{
			ID:         0,
			Name:       "2001: Космическая Одиссея",
			Price:      "600",
			Discount:   5,
			ISBN:       "978-5-04-095784-2",
			Photo:      "static/2664811.webp",
			Desc:       "bla-bla",
			Dimensions: [3]int{10, 20, 13},
			Authors:    []string{"Артур Чарльз Кларк"},
			Publishers: []string{"Эксмо"},
			Genres:     []string{"Фантастика"},
			Amount:     10,
			Date:       "1968-1-1",
		},*/
		/*{
			ID:         0,
			Name:       "Город и звезды",
			Price:      "656.55",
			Discount:   0,
			ISBN:       "978-5-17-105787-9",
			Photo:      "default",
			Desc:       "Артура Кларка, Айзека Азимова и Роберта Хайнлайна называют «большой тройкой»,",
			Dimensions: [3]int{18, 11, 16},
			Authors:    []string{"Артур Чарльз Кларк"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фантастика"},
			Amount:     5,
			Date:       "2021-10-5",
		},*/
		{
			ID:         0,
			Name:       "Конец вечности",
			Price:      "566.55",
			Discount:   5,
			ISBN:       "978-5-04-106735-9",
			Photo:      "default",
			Desc:       "бла бла",
			Dimensions: [3]int{10, 20, 10},
			Authors:    []string{"Айзек Азимов"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фантастика"},
			Amount:     5,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Хоббит",
			Price:      "408",
			Discount:   18,
			ISBN:       "978-5-17-114531-6",
			Photo:      "static/2715497.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{20, 13, 2},
			Authors:    []string{"Джон Толкиен"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фэнтези"},
			Amount:     10,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Сильмариллион",
			Price:      "1999",
			Discount:   12,
			ISBN:       "978-5-17-088588-6",
			Photo:      "static/2715497.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{24, 16, 3},
			Authors:    []string{"Джон Толкиен"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фэнтези"},
			Amount:     10,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Властелин колец",
			Price:      "2299",
			Discount:   12,
			ISBN:       "978-5-17-092791-3",
			Photo:      "static/2503148.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{22, 14, 7},
			Authors:    []string{"Джон Толкиен"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фэнтези"},
			Amount:     10,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Чужак в стране чужой",
			Price:      "999",
			Discount:   19,
			ISBN:       "978-5-389-22044-7",
			Photo:      "static/2960597.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{22, 14, 3},
			Authors:    []string{"Роберт Хайнлайн"},
			Publishers: []string{"Азбука"},
			Genres:     []string{"Фантастика"},
			Amount:     10,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Луна - суровая госпожа",
			Price:      "780",
			Discount:   0,
			ISBN:       "978-5-389-13691-5",
			Photo:      "static/2619838.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{22, 14, 3},
			Authors:    []string{"Роберт Хайнлайн"},
			Publishers: []string{"Азбука"},
			Genres:     []string{"Фантастика"},
			Amount:     10,
			Date:       "2020-10-5",
		},
		{
			ID:         0,
			Name:       "Свидание с Рамой",
			Price:      "599",
			Discount:   12,
			ISBN:       "978-5-17-109871-1",
			Photo:      "static/2662287.webp",
			Desc:       "бла бла",
			Dimensions: [3]int{18, 12, 2},
			Authors:    []string{"Артур Чарльз Кларк"},
			Publishers: []string{"АСТ"},
			Genres:     []string{"Фантастика"},
			Amount:     10,
			Date:       "2020-10-5",
		},
	}
	for _, book := range testBooks {
		AddBook(book)
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
		//DeleteBook(book.ID)
	}
}
