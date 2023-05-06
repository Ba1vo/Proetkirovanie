package proxy

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetBook(id int) (fullBook, error) {
	var book fullBook
	db, err := sql.Open("postgres", queries.PsqlInfo)
	if err != nil {
		return book, errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return book, errors.New("Connection error")
	}

	query := fmt.Sprintf(`
	SELECT 
		b."id",
		b."name", 
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
		return book, errors.New("auerie error")
	}
	defer row.Close()
	if row.Next() {
		var authors, genres, publishers string
		row.Scan(&book.ID, &book.Name, &book.ISBN, &book.Photo, &book.Date, &book.Desc, &book.Price, &book.Discount,
			&book.Dimensions[0], &book.Dimensions[1], &book.Dimensions[2], &book.Amount, &authors, &genres, &publishers)
		book.Authors = strings.Split(authors, ", ")
		book.Genres = strings.Split(genres, ", ")
		book.Publishers = strings.Split(publishers, ", ")
		return book, nil
	} else {
		return book, errors.New("Empty")
	}
}
