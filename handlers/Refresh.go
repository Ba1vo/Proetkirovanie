package handlers

import (
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	id := crypt.CookieIsValid(r.Cookies(), "Refresh")
	if id > 0 {
		crypt.SetCookies(w, id)
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "Token", MaxAge: -1})
	http.SetCookie(w, &http.Cookie{Name: "Refresh", MaxAge: -1})
	w.WriteHeader(http.StatusBadRequest)
}
