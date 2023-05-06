package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/queries"
)

func DiscBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Disc in use")
	//if misc.CheckSign(r.Cookies(), &id) {
	//	w.WriteHeader(http.StatusTeapot)
	//	return
	//} proverku i ustanovku v odno mesto, esli cookie est obnovit inache chel bez akkaunta
	//misc.SetCookies(w, id)
	//id = queries.Auth(d.Creds)
	books, err := queries.GetDiscBooks()
	if err == nil {

		output, err := json.Marshal(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
