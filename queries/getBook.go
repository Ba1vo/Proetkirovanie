package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetBook(id int) (decoder.FullBook, error) {
	var book decoder.FullBook
	db, err := sql.Open("postgres", PsqlInfo)
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
		b."name",
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
		fmt.Println(err.Error())
		return book, errors.New("auerie error")
	}
	defer row.Close()
	if row.Next() {
		var authors, genres, publishers string
		row.Scan(&book.ID, &book.ISBN, &book.Name, &book.Photo, &book.Date, &book.Desc, &book.Price, &book.Discount,
			&book.Dimensions[0], &book.Dimensions[1], &book.Dimensions[2], &book.Amount, &authors, &genres, &publishers)
		fmt.Println(book.Amount)
		fmt.Println(book.Name)
		book.Authors = strings.Split(authors, ", ")
		fmt.Println(book.Authors)
		book.Genres = strings.Split(genres, ", ")
		fmt.Println(book.Genres)
		book.Publishers = strings.Split(publishers, ", ")
		fmt.Println(book.Publishers)
		return book, nil
	} else {
		return book, errors.New("Empty")
	}
}
