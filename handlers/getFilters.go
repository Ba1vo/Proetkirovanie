package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ba1vo/Proektirovanie/queries"
)

func GetFilters(w http.ResponseWriter, r *http.Request) {
	filters, err := queries.GetFilters()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
