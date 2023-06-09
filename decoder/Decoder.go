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
	Role  int
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
	Favourite  bool
}

type CardBook struct {
	ID        int
	Name      string
	Authors   []string
	Price     string
	Discount  int
	Photo     string
	Amount    int
	Favourite bool
}

type Order struct {
	ID     int
	Date   string
	Adress string
	Status string
	Books  []CardBook
}

type ServerOrder struct {
	Adress   string
	Book_IDs []int
	Amounts  []int
}

type SearchOptions struct {
	Order    string
	Str      string
	Genres   []string
	MaxDate  string
	MinDate  string
	MaxPrice string
	MinPrice string
	PageSize int
	Page     int
}

type Data struct {
	Email string
	Code  string
}

type Decodable interface {
	*RegUser | *AuthUser | *Order | *FullBook | *ServerOrder | *SearchOptions | *int | *[]int | *Data
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
