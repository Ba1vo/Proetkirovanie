package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	d.Order = r.URL.Query().Get("order_by")
	PageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d.PageSize = PageSize
	if r.URL.Query().Has("page") {
		d.Page, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		d.Page = 1
	}
	dates := strings.Split(r.URL.Query().Get("dates"), ";")
	if dates[0] != "" {
		d.MinDate = dates[0]
		d.MaxDate = dates[1]
	}
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
