package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	id := crypt.CookieIsValid(r.Cookies(), "Token")
	fmt.Printf("Userd ord id %d", id)
	if id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	books, err := queries.GetOrders(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
