package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetFilters() (decoder.SearchOptions, error) {
	var options decoder.SearchOptions
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return options, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return options, errors.New("connection error")
	}

	query := fmt.Sprintf(
		`SELECT 
			MAX(b."pub_date"), 
			MIN(b."pub_date"),
			MAX(b."price"), 
			MIN(b."price"), 
			(SELECT string_agg(g."name", ', ') AS genres FROM public.genres AS g )
		FROM public.books as b;`)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return options, errors.New("querie error")
	}
	defer rows.Close()

	for rows.Next() {
		var genres string
		//var amount int
		rows.Scan(&options.MaxDate, &options.MinDate, &options.MaxPrice, &options.MinPrice, &genres)
		options.Genres = strings.Split(genres, ", ")
		//book.Amount = amount
	} // if 0 books found,

	return options, nil
}
