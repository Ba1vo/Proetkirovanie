package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/email"
	"github.com/Ba1vo/Proektirovanie/generator"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var d decoder.RegUser
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !(checks.CheckUserReg(d)) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(d)
	//d.Email = checks.InjectCheck(d.Email)
	//d.Name = checks.InjectCheck(d.Name)
	d.Pass = crypt.Hash(d.Pass)
	err := queries.AddUser(d)
	if err == nil {
		if Error := email.Email(d.Email, generator.RandStringBytes()); Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(nil)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
