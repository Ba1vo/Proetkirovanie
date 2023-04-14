package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddFavourite(userId int, book_id int) error {
	db, err := sql.Open("postgres", PsqlInfo) //insert hash code
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("connection error")
	}
	query := fmt.Sprintf(`INSERT INTO public."favourites" (
		"user_id", "book_id")
		VALUES (%d, %d);`, userId, book_id) //CHECK
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("querie error")
	}
	defer row.Close()
	return nil
}
