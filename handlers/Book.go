package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

var ErrCreds = errors.New("creds error")
var ErrConnection = errors.New("connection error")
var ErrQuerie = errors.New("querie error")
var ErrEmpty = errors.New("empty")

func Book(w http.ResponseWriter, r *http.Request) {
	var d decoder.BookId
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//if !(checks.CheckUserAuth(d)) {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	book, err := queries.GetBook(d.Id)
	if err == ErrEmpty {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.Write(output)
}
