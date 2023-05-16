package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func GetSearchBooks(id int, options decoder.SearchOptions) ([]decoder.CardBook, error) {
	var sb strings.Builder
	var clauses []string
	var Books []decoder.CardBook
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("creds error")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("connection error")
	}

	searchString := strings.Replace(options.Str, " ", " & ", -1)
	searchString = strings.ToLower(searchString)
	sb.WriteString(`SELECT 
		b."id", 
		b."name", 
		string_agg(au."disp_name", ', '), 
		b."price", 
		b."discount", 
		b."amount",
		b."photo",
		(f."user_id" IS NOT NULL)
	FROM "books" AS b
	LEFT JOIN "search_vectors" AS srch ON srch."book_id" = b."id"
	LEFT JOIN "favourites" AS f ON f."book_id" = b."id" AND f."user_id" = `)
	sb.WriteString(strconv.Itoa(id))
	sb.WriteString(`LEFT JOIN "books_authors" AS ba ON ba."book_id" = b."id"
	LEFT JOIN "authors" AS au ON au."id" = ba."author_id"`)
	fmt.Printf("Search Object: %v\n", options)
	if options.Genres[0] != "" {
		sb.WriteString(`
		LEFT JOIN "books_genres" AS bg ON b."id" = bg."book_id"
		LEFT JOIN "genres" AS g ON bg."genre_id" = g."id"
		`)
		clauses = append(clauses, fmt.Sprintf(`g."name" in ('%s')`, strings.Join(options.Genres, "', '")))
	}
	if searchString != "" {
		clauses = append(clauses, fmt.Sprintf(`srch."vector" @@ to_tsquery('%s')`, searchString))
	}
	if options.MaxDate != "" {
		clauses = append(clauses, fmt.Sprintf(`b."pub_date" <= '%s-12-31'`, options.MaxDate))
	}
	if options.MinDate != "" {
		clauses = append(clauses, fmt.Sprintf(`b."pub_date" >= '%s-1-1'`, options.MinDate))
	}
	if options.MaxPrice != "" {
		clauses = append(clauses, fmt.Sprintf(`b."price" < '%s'`, options.MaxPrice))
	}
	if options.MinPrice != "" {
		clauses = append(clauses, fmt.Sprintf(`b."price" > '%s'`, options.MinPrice))
	}
	if len(clauses) > 0 {
		sb.WriteString("WHERE ")
		sb.WriteString(strings.Join(clauses, " AND "))
	}
	sb.WriteString(`GROUP BY b."id", f."user_id" `)
	if searchString != "" {
		sb.WriteString(`, "vector" `)
	} else {
		//sb.WriteString(`ORDER BY ts_rank("vector", to_tsquery('%s')) DESC;`) ???
	}
	sb.WriteString("ORDER BY ")
	if options.Order != "" {
		sb.WriteString(options.Order)
		sb.WriteString(", ")
	}
	if searchString != "" {
		sb.WriteString(fmt.Sprintf(`ts_rank("vector", to_tsquery('%s')) DESC`, searchString))
		sb.WriteString(", ")
	}
	sb.WriteString("1")
	sb.WriteString(" LIMIT ")
	sb.WriteString(strconv.Itoa(options.PageSize))
	sb.WriteString(" OFFSET ")
	sb.WriteString(strconv.Itoa(options.PageSize * (options.Page - 1)))
	sb.WriteString(";")
	fmt.Println(sb.String())
	rows, err := db.Query(sb.String())
	if err != nil {
		fmt.Println(err.Error())
		return Books, errors.New("querie error")
	}
	defer rows.Close()

	for rows.Next() {
		var book decoder.CardBook
		//var amount int
		var array string
		rows.Scan(&book.ID, &book.Name, &array, &book.Price, &book.Discount, &book.Amount, &book.Photo, &book.Favourite)
		//book.Amount = amount
		book.Authors = strings.Split(array, ", ")
		Books = append(Books, book)
	} // if 0 books found,

	return Books, nil
}
