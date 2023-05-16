package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	var d int
	if decoder.DecodeJSON(&d, r) {
		fmt.Println("decod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := crypt.CookieIsValid(r.Cookies(), "Token")
	if id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("delfav in use")
	err := queries.DeleteFavourite(id, d)
	if err != nil {
		fmt.Println("querie")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.WriteHeader(http.StatusOK)
}
