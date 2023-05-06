package lightweight

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/queries"
)

type Item interface {
	Display()
	SetPostion(string)
}

type FullBook struct {
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
	Position   string
}

func (c *FullBook) SetPostion(pos string) {
	c.Position = pos
}

func (c *FullBook) Display() {
	fmt.Printf("Book %s written by %s displayed at %s\n", c.Name, c.Authors[0], c.Position)
}

type lightweight struct {
	books map[int]*FullBook
}

func NewLight() lightweight {
	return lightweight{books: map[int]*FullBook{}}
}

func (l *lightweight) GetBook(id int) Item {
	_, ans := l.books[id]
	if !ans {
		fmt.Printf("Adding Book %d\n", id)
		book, err := GetDBBook(id)
		if err != nil {
			return nil
		}
		l.books[id] = &book
	}
	return l.books[id]
}

func GetDBBook(id int) (FullBook, error) {
	var book FullBook
	db, err := sql.Open("postgres", queries.PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return book, errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return book, errors.New("Connection error")
	}

	query := fmt.Sprintf(`
	SELECT 
		b."id", 
		b."isbn",
		b."photo",
		b."pub_date",
		b."desc",
		b."price",
		b."discount", 
		b."width",  
		b."length", 
		b."height", 
		b."amount", 
		string_agg(au."disp_name", ', '), 
		string_agg(g."name", ', '), 
		string_agg(p."name", ', ') 
	FROM "books" AS b
	JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	JOIN "authors" AS au ON au."id" = ba."author_id"
	
	JOIN "books_genres" AS bg ON bg."book_id" = b."id"
	JOIN "genres" AS g ON g."id" = bg."genre_id"

	JOIN "books_publishers" AS bp ON bp."book_id" = b."id"
	JOIN "publishers" AS p ON p."id" = bp."publisher_id"

	WHERE b."id" = %d
	GROUP BY b."id";`, id)
	row, err := db.Query(query)
	if err != nil {
		//fmt.Println(err.Error())
		return book, errors.New("auerie error")
	}
	defer row.Close()
	if row.Next() {
		var authors, genres, publishers string
		row.Scan(&book.ID, &book.ISBN, &book.Photo, &book.Date, &book.Desc, &book.Price, &book.Discount,
			&book.Dimensions[0], &book.Dimensions[1], &book.Dimensions[2], &book.Amount, &authors, &genres, &publishers)
		//fmt.Println(book.Amount)
		//fmt.Println(book.Name)
		book.Authors = strings.Split(authors, ", ")
		//fmt.Println(book.Authors)
		book.Genres = strings.Split(genres, ", ")
		//fmt.Println(book.Genres)
		book.Publishers = strings.Split(publishers, ", ")
		//fmt.Println(book.Publishers)
		return book, nil
	} else {
		return book, errors.New("Empty")
	}
}
