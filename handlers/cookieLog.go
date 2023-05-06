package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func CookieLogin(w http.ResponseWriter, r *http.Request) {
	id := crypt.CookieIsValid(r.Cookies(), "Token")
	if id == 0 {
		http.SetCookie(w, &http.Cookie{Name: "Token", MaxAge: -1})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(nil)
		return
	}
	fullUser, err := queries.GetUserByID(id)
	if err == nil {
		output, err := json.Marshal(fullUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
