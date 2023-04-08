package handlers

import (
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func RedactBook(w http.ResponseWriter, r *http.Request) {
	var d decoder.FullBook
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !(checks.CheckBook(d)) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//d.Email = checks.InjectCheck(d.Email)
	//d.Pass = crypt.Hash(d.Pass)
	//fullUser, err := queries.GetUser(d)
	err := queries.UpdateBook(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
