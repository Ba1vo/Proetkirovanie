package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	var ids []int
	if decoder.DecodeJSON(&ids, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	books, err := queries.GetCart(ids)
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
