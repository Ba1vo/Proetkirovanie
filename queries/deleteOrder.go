package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func DeleteOrder(id int) error {
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
	query := fmt.Sprintf(`DELETE FROM "orders" AS ord WHERE ord."id" = %d`, id) //CHECK
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Querie error")
	}
	return nil
}
