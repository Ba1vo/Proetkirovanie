package strategy

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/queries"
)

type Books interface {
	info()
}

type strategy interface {
	search(SearchOptions) Books
}

type UserBook struct {
	ID       int
	Name     string
	Authors  []string
	Price    string
	Discount int
	Photo    string
	Amount   int
}

type UserBooks struct {
	books []UserBook
}

func (c *UserBooks) info() {
	fmt.Printf("Found follownig books %v", c)
}

type StaffBook struct {
	ID       int
	isbn     string
	Name     string
	Authors  []string
	Price    string
	Discount int
	Photo    string
	Amount   int
}

type StaffBooks struct {
	books []StaffBook
}

func (c *StaffBooks) info() {
	fmt.Printf("Found following books %v", c)
}

type search struct {
	opt       SearchOptions
	algorithm strategy
}

func (s *search) SearchBooks() Books {
	return s.algorithm.search(s.opt)
}

type SearchOptions struct {
	Str        string
	Genre      string
	Publishers string
	StartDate  string
	EndDate    string
}

type StaffSearch struct{}

func (s *StaffSearch) search(options SearchOptions) Books {
	var Books StaffBooks
	db, err := sql.Open("postgres", queries.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}

	searchString := strings.Replace(options.Str, " ", " & ", -1)
	searchString = strings.ToLower(searchString)
	query := fmt.Sprintf(
		`SELECT 
		b."id",
		b."isbn", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo"
	FROM "books" AS b
    LEFT JOIN "search_vectors" AS srch ON srch."book_id" = b."id"
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
    
	WHERE srch."vector" @@ to_tsquery('%s')
	GROUP BY b."id", "vector"
    ORDER BY ts_rank("vector", to_tsquery('%s')) DESC;`, searchString, searchString)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}
	defer rows.Close()

	for rows.Next() {
		var book StaffBook
		//var amount int
		var array string
		rows.Scan(&book.ID, &book.isbn, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo)
		//book.Amount = amount
		book.Authors = strings.Split(array, ", ")
		Books.books = append(Books.books, book)
	} // if 0 books found,

	return &Books
}

type UserSearch struct{}

func (s *UserSearch) search(options SearchOptions) Books {
	var Books UserBooks
	db, err := sql.Open("postgres", queries.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}

	searchString := strings.Replace(options.Str, " ", " & ", -1)
	searchString = strings.ToLower(searchString)
	query := fmt.Sprintf(
		`SELECT 
		b."id", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo"
	FROM "books" AS b
    LEFT JOIN "search_vectors" AS srch ON srch."book_id" = b."id"
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
    
	WHERE srch."vector" @@ to_tsquery('%s')
	GROUP BY b."id", "vector"
    ORDER BY ts_rank("vector", to_tsquery('%s')) DESC;`, searchString, searchString)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return &Books
	}
	defer rows.Close()

	for rows.Next() {
		var book UserBook
		//var amount int
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo)
		//book.Amount = amount
		book.Authors = strings.Split(array, ", ")
		Books.books = append(Books.books, book)
	} // if 0 books found,

	return &Books
}
