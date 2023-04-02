package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func getOrders(userId int) ([]decoder.Order, error) {
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
			ob."price", 
			ob."discount", 
			ob."amount"
			string_agg(au."disp_name", ', '),
		FROM "orders" AS o
		JOIN "orders_books" AS ob ON ob."order_id" = o."id"
		JOIN "books_authors" AS ba ON ba."book_id" = ob."book_id"
		JOIN "authors" AS au ON au."id" = ba."author_id"
		
		WHERE o."user_id" = %d
		GROUP BY b."id"
		ORDER BY o."id" ASC, ob."book_id" DESC;`, userId)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Orders, errors.New("querie error")
	}
	defer rows.Close()
	var order decoder.Order
	for rows.Next() {
		var book decoder.CardBook
		var array, status, adress, date string
		var id int
		rows.Scan(&id, &status, &adress, &date, &book.ID, &book.Price, &book.Discount, &book.Amount, &array)
		if order.ID == 0 || order.ID != id {
			Orders = append(Orders, order)
			order.ID = id
			order.Status = status
			order.Adress = adress
			order.Date = date
		}
		book.Authors = strings.Split(array, ", ")
		order.Books = append(order.Books, book)
	} // if 0 books found,

	return Orders, nil
}
