package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetFavourites(id int, pageSize int, page int) ([]decoder.CardBook, error) {
	var Books []decoder.CardBook
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("connection error")
	}

	query := fmt.Sprintf(
		`SELECT 
		b."id", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo"
	FROM "favourites" AS f
    LEFT JOIN "books" AS b ON b."id" = f."book_id"
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
    
	WHERE f."user_id" = %d
	GROUP BY b."id"
	LIMIT %d OFFSET %d;`, id, pageSize, (page-1)*pageSize)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("querie error")
	}
	defer rows.Close()

	for rows.Next() {
		var book decoder.CardBook
		//var amount int
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo)
		//book.Amount = amount
		book.Authors = strings.Split(array, ", ")
		book.Favourite = true
		Books = append(Books, book)
	} // if 0 books found,

	return Books, nil
}
