package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	var d decoder.FavBook
	if decoder.DecodeJSON(&d, r) {
		fmt.Println("decod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("delfav in use")
	err := queries.DeleteFavourite(d.User, d.Book)
	if err != nil {
		fmt.Println("querie")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.WriteHeader(http.StatusOK)
}
