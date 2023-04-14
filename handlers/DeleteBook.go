package handlers

import (
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var d int
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := queries.DeleteBook(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.WriteHeader(http.StatusOK)
}
