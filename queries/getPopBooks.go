package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetPopBooks(id int) ([]decoder.CardBook, error) {
	var books []decoder.CardBook
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return books, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return books, errors.New("connection error")
	}
	query := fmt.Sprintf(`SELECT 
		b."id", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo",
		(f."user_id" IS NOT NULL)
	FROM "books" AS b
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "favourites" AS f ON f."book_id" = b."id" AND f."user_id" = %d
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
	LEFT JOIN "orders_books" AS ob ON ob."book_id" = b."id"
	GROUP BY b."id", f."user_id"
	ORDER BY COALESCE(SUM(ob."amount"), 0) DESC
	LIMIT 15;`, id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return books, errors.New("querie error")
	}
	defer rows.Close()
	for rows.Next() {
		var book decoder.CardBook
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo, &book.Favourite)
		book.Authors = strings.Split(array, ", ")
		books = append(books, book)
	}

	return books, nil
	/*if row.Next() {
		row.Scan(&Info.ID, &Info.Name, &Info.Email, &Info.Pass, &Info.Date)
		return Info, nil
	} else {
		return Info, errors.New("Empty")
	}*/
}
