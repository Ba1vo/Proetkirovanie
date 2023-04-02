package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetMainBooks() (decoder.MainBooks, error) {
	var Books decoder.MainBooks
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("Connection error")
	}

	query := fmt.Sprintf(`SELECT b."id", b."name", string_agg(a.disp_name, ', '), u."price", u."discount", b."amount"
	FROM "books" AS b
	JOIN "books_authors" AS ba ON ba.books_id = b.id
	JOIN "authors" AS a ON a.id = ba.authors_id
	ORDER BY (SELECT coalesce(SELECT SUM(o."amount") FROM "orders_books" as o WHERE o."books_id" = b.id), 0)
	LIMIT 15`)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("Querie error")
	}
	defer rows.Close()
	var pop []decoder.CardBook
	for rows.Next() {
		var book decoder.CardBook
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount)
		book.Authors = strings.Split(array, ", ")
		pop = append(pop, book)
	}

	return Books, nil
	/*if row.Next() {
		row.Scan(&Info.ID, &Info.Name, &Info.Email, &Info.Pass, &Info.Date)
		return Info, nil
	} else {
		return Info, errors.New("Empty")
	}*/
}
