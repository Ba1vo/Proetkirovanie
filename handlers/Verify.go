package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	var d decoder.Data
	if decoder.DecodeJSON(&d, r) {
		fmt.Println("Decoding error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("IM DOING")
	if !(checks.CheckEmail(d.Email) && checks.CheckCode(d.Code)) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err := queries.UpdateUserStatus(d.Email, d.Code, 1)
	if err == nil {
		w.Write(nil)
		return
	}
	if err.Error() == "empty" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
