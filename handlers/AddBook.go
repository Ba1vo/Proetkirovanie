package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	var d decoder.FullBook
	if decoder.DecodeJSON(&d, r) {
		fmt.Println("decod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !(checks.CheckBook(d)) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("check")
		return
	}
	//d.Email = checks.InjectCheck(d.Email)
	//d.Pass = crypt.Hash(d.Pass)
	//fullUser, err := queries.GetUser(d)
	id, err := queries.AddBook(d)
	if err != nil {
		fmt.Println("querie")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(id)
	if err != nil {
		fmt.Println("marshal")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.Write(output)
}
