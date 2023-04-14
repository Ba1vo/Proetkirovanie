package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func AddFavourite(w http.ResponseWriter, r *http.Request) {
	var d decoder.FavBook
	if decoder.DecodeJSON(&d, r) {
		fmt.Println("decod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//d.Email = checks.InjectCheck(d.Email)
	//d.Pass = crypt.Hash(d.Pass)
	//fullUser, err := queries.GetUser(d)
	err := queries.AddFavourite(d.User, d.Book)
	if err != nil {
		fmt.Println("querie")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.WriteHeader(http.StatusOK)
}
