package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var d decoder.AuthUser
	if decoder.DecodeJSON(&d, r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !(checks.CheckUserAuth(d)) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	d.Email = checks.InjectCheck(d.Email)
	d.Pass = crypt.Hash(d.Pass)
	fullUser, err := queries.GetUser(d)
	if err == nil {
		output, err := json.Marshal(fullUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		crypt.SetCookies(w, fullUser.ID)
		//w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(output)
		return
	}
	if err.Error() == "Empty" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неправильная почта или пароль"))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
