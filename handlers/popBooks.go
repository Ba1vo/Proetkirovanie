package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func PopBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pop in use")
	id := crypt.CookieIsValid(r.Cookies(), "Token")
	books, err := queries.GetPopBooks(id)
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
