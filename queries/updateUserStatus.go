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
		return errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Connection error")
	}
	query := fmt.Sprintf(`UPDATE public."users"
		SET code = NULL, role = %d 
		WHERE email = '%s' AND code = '%s';`, role, email, code) //CHECK
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Querie error")
	}
	defer row.Close()
	return nil
}
