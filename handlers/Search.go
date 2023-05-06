package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func Search(w http.ResponseWriter, r *http.Request) {
	var d decoder.SearchOptions
	d.Str = r.URL.Query().Get("str")
	d.Genres = strings.Split(r.URL.Query().Get("genres"), ";")
	d.MaxPrice = r.URL.Query().Get("max_price")
	d.MinPrice = r.URL.Query().Get("min_price")
	dates := strings.Split(r.URL.Query().Get("dates"), ";")
	d.MaxDate = dates[1]
	d.MinDate = dates[0]
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
