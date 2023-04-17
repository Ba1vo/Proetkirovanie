package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func AddOrder(order decoder.ServerOrder) (int, error) {
	var id int
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return id, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return id, errors.New("connection error")
	}
	query := fmt.Sprintf(`
	with o_id as (INSERT INTO public."orders" ("user_id", "adress") VALUES ('%d', '%s') RETURNING "id")
	INSERT INTO public."orders_books" ("order_id", "book_id", "amount")
	SELECT "id", unnest('{%s}'::int[]), unnest('{%s}'::int[]) FROM o_id;`, order.UserID, order.Adress, intToString(order.Book_IDs), intToString(order.Amounts)) //CHECK
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return id, errors.New("querie error")
	}
	defer row.Close()
	return id, nil
}

func intToString(array []int) string {
	strArray := []string{}

	for _, v := range array {
		str := strconv.Itoa(v)
		strArray = append(strArray, str)
	}
	return strings.Join(strArray, ", ")
}
