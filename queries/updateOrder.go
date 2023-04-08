package queries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func UpdateOrder(order decoder.ServerOrder) error {
	db, err := sql.Open("postgres", PsqlInfo)
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
	query := fmt.Sprintf(`
	with o_id as (UPDATE public."orders" ("adress", "status") VALUES ('%s', '%s') WHERE id = %d AND user_id = %d RETURNING "id")
	UPDATE public."orders_books"  ("book_id", "amount")
	SELECT o_id, unnest({%s}::int[]), unnest({%s}::int[]))
	ON CONFLICT DO NOTHING;`, order.Adress, order.Status, order.ID, order.UserID, intToString(order.Book_IDs), intToString(order.Amounts)) //CHECK
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Querie error")
	}
	defer row.Close()
	return nil
}
