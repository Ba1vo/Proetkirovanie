package handlers

import (
	"net/http"

	"github.com/Ba1vo/Proektirovanie/checks"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("c")
	email := r.URL.Query().Get("email")
	if !(checks.CheckEmail(email) && checks.CheckCode(code)) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err := queries.UpdateUserStatus(email, code, 1)
	if err == nil {
		w.Write(nil)
		return
	}
	if err.Error() == "Querie error" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	return
}
