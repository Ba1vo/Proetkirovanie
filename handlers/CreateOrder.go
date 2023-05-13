package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var d decoder.ServerOrder
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userid := crypt.CookieIsValid(r.Cookies(), "Token")
	if userid == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := queries.AddOrder(userid, d)
	if err == ErrEmpty {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.Write(output)
}
