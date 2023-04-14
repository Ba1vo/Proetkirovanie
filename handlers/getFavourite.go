package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetFavourites(w http.ResponseWriter, r *http.Request) {
	var d int
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//if !(checks.CheckUserAuth(d)) {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	books, err := queries.GetFavourites(d)
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
