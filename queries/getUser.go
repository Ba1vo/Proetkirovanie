package queries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

const (
	host     = "localhost"
	port     = 5446
	user     = "postgres"
	password = "159753"
	dbname   = "postgres"
)

var PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

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

	query := fmt.Sprintf(`SELECT u."id", u."name", u."email", u."pass", u."reg_Date"
	FROM "users" AS u
	WHERE u."email" = '%s' AND u."pass" = '%s';`, User.Email, User.Pass)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return Info, errors.New("Querie error")
	}
	defer row.Close()
	if row.Next() {
		row.Scan(&Info.ID, &Info.Name, &Info.Email, &Info.Pass, &Info.Date)
		return Info, nil
	} else {
		return Info, errors.New("Empty")
	}
}
