package decoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegUser struct {
	Name  string
	Email string
	Pass  string
}

type InfoUser struct {
	ID    int
	Name  string
	Email string
	Pass  string
	Date  string
}

type AuthUser struct {
	Email string
	Pass  string
}

type FullBook struct {
	ID         int
	Name       string
	Price      string
	Discount   int
	ISBN       string
	Photo      string
	Desc       string
	Dimensions [3]int
	Authors    []string
	Publishers []string
	Genres     []string
	Amount     int
	Date       string
}

type CardBook struct {
	ID       int
	Name     string
	Authors  []string
	Price    string
	Discount int
	Photo    string
	Amount   int
}

type MainBooks struct {
	Popular []CardBook
	New     []CardBook
}

type Order struct {
	ID     int
	Date   string
	Adress string
	Status string
	Books  []CardBook
}

type ServerOrder struct {
	ID       int
	UserID   int
	Adress   string
	Status   string
	Book_IDs []int
	Amounts  []int
}

type SearchOptions struct {
	Str        string
	Genre      string
	Publishers string
	StartDate  string
	EndDate    string
}

type BookId struct {
	Id int
}

type Decodable interface {
	*RegUser | *AuthUser | *Order | *FullBook | *BookId | *ServerOrder | *SearchOptions
}

type Writable interface {
}

func DecodeJSON[JSON Decodable](obj JSON, r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
