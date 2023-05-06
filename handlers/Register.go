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
		fmt.Println("Decoding error")
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
		code := generator.RandStringBytes()
		if err := queries.SetUserCode(d.Email, code); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := email.Email(d.Email, code); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Такая почта занята"))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
