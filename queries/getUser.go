package queries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetUser(User decoder.AuthUser) (decoder.InfoUser, error) {
	var Info decoder.InfoUser
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return Info, errors.New("Creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return Info, errors.New("Connection error")
	}

	query := fmt.Sprintf(`SELECT u."id", u."name", u."email", u."reg_Date"
	FROM "users" AS u
	WHERE u."email" = '%s' AND u."pass" = '%s';`, User.Email, User.Pass) //Switch between id and email&pass
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Info, errors.New("Querie error")
	}
	defer row.Close()
	if row.Next() {
		row.Scan(&Info.ID, &Info.Name, &Info.Email, &Info.Date)
		return Info, nil
	} else {
		return Info, errors.New("Empty")
	}
}
