package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetOrders(userId int) ([]decoder.Order, error) {
	var Orders []decoder.Order
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return Orders, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return Orders, errors.New("connection error")
	}

	query := fmt.Sprintf(
		`SELECT 
		o."id", 
		o."status",
		o."adress",
		o."creation_date",
		ob."book_id",
		b."name",
		ob."price", 
		ob."discount", 
		ob."amount",
		b."photo",
		string_agg(au."disp_name", ', ')
	FROM "orders" AS o
	LEFT JOIN "orders_books" AS ob ON ob."order_id" = o."id"
	LEFT JOIN "books" as b ON b."id" = ob."book_id"
	LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"
	WHERE o."user_id" = %d
	GROUP BY o."id", ob."book_id", ob."book_id",b."name", b."photo",  ob."price", ob."discount", ob."amount"
	ORDER BY o."id" DESC, ob."book_id" DESC;`, userId)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Orders, errors.New("querie error")
	}
	defer rows.Close()
	order := decoder.Order{ID: 0}
	for rows.Next() {
		var book decoder.CardBook
		var array, status, adress, date string
		var id int
		rows.Scan(&id, &status, &adress, &date, &book.ID, &book.Name, &book.Price, &book.Discount, &book.Amount, &book.Photo, &array)
		if order.ID == 0 || order.ID != id {
			Orders = append(Orders, order)
			order = decoder.Order{
				ID:     id,
				Status: status,
				Adress: adress,
				Date:   date,
			}
		}
		book.Authors = strings.Split(array, ", ")
		fmt.Printf("\n Book val %v \n", book)
		order.Books = append(order.Books, book)
	} // if 0 books found,
	Orders = append(Orders, order)
	Orders = Orders[1:]
	return Orders, nil
}
