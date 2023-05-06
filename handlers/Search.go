package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Search(w http.ResponseWriter, r *http.Request) {
	var d decoder.SearchOptions
	d.Str = r.URL.Query().Get("str")
	//if !(checks.CheckUserAuth(d)) {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	fmt.Println("search in use")
	books, err := queries.GetSearchBooks(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("BKS: %v \n", books)
	output, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//crypt.SetCookies(w, fullUser.ID)
	w.Write(output)
}
