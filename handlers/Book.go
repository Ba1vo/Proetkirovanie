package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Ba1vo/Proektirovanie/queries"
	"github.com/gorilla/mux"
)

var ErrCreds = errors.New("creds error")
var ErrConnection = errors.New("connection error")
var ErrQuerie = errors.New("querie error")
var ErrEmpty = errors.New("empty")

func Book(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := queries.GetBook(id)
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
