package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func UpdateOrder(book_ids []int) ([]decoder.CardBook, error) {
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
	query := fmt.Sprintf(`
	SELECT 
		b."id", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo"
	FROM "books" AS b
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
	WHERE b."id" IN (%s)
	GROUP BY b."id";`,
		intToString(book_ids)) //CHECK
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return books, errors.New("querie error")
	}
	defer rows.Close()

	for rows.Next() {
		var book decoder.CardBook
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo)
		book.Authors = strings.Split(array, ", ")
		books = append(books, book)
	}
	return books, nil
}
