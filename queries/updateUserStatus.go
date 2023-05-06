package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func UpdateUserStatus(email string, code string, role int) error {
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
	fmt.Println("IM DOING SMTH")
	query := fmt.Sprintf(`UPDATE public."users"
		SET code = NULL, role = %d 
		WHERE email = '%s' AND code = '%s';`, role, email, code) //CHECK
	row, err := db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("querie error")
	}
	if num, _ := row.RowsAffected(); num == 0 {
		fmt.Println("no rows affected")
		return errors.New("empty")
	}
	return nil
}
