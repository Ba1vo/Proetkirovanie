package queries

import (
	"database/sql"
	"errors"
	"fmt"
)

func SetUserCode(email string, code string) error {
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
		SET code = '%s'
		WHERE email = '%s';`, code, email) //CHECK
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Querie error")
	}
	defer row.Close()
	return nil
}
