package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetFavourites(w http.ResponseWriter, r *http.Request) {
	var pageSize, page, id int
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id = crypt.CookieIsValid(r.Cookies(), "Token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	books, err := queries.GetFavourites(id, pageSize, page)
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
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
