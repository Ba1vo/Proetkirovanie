package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func UpdateBook(book decoder.FullBook) error {
	db, err := sql.Open("postgres", PsqlInfo) //insert hash code
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Connection error")
	}
	query := fmt.Sprintf(`SELECT * FROM update_book(%d::int,'%s'::varchar, %s, %d::smallint, '%s'::varchar, '%s'::varchar, '%s'::text, '{%d, %d, %d}'::int[], '{%s}'::varchar[], '{%s}'::varchar[], '{%s}'::varchar[], %d, '%s'::date)`,
		book.ID, book.Name, book.Price, book.Discount, book.ISBN, book.Photo, book.Desc, book.Dimensions[0], book.Dimensions[1], book.Dimensions[2],
		strings.Join(book.Authors, ", "), strings.Join(book.Publishers, ", "), strings.Join(book.Genres, ", "), book.Amount, book.Date)
	//CHECK
	row, err := db.Query(query) //Can take book_id if needed
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Querie error")
	}
	defer row.Close()
	return nil
}
