package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func DiscBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Disc in use")
	id := crypt.CookieIsValid(r.Cookies(), "Token")
	books, err := queries.GetDiscBooks(id)
	if err == nil {

		output, err := json.Marshal(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
