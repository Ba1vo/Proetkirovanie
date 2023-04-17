package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var d []int
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	books, err := queries.UpdateOrder(d)
	if err == ErrEmpty {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.Write(output)
}
