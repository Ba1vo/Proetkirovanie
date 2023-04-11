package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

func TestAddBook(t *testing.T) {
	d := decoder.FullBook{
		ID:         0,
		Name:       "Конец вечности",
		Price:      "555.55",
		Discount:   5,
		ISBN:       "978-5-04-106735-9",
		Photo:      "default",
		Desc:       "бла бла",
		Dimensions: [3]int{10, 20, 10},
		Authors:    []string{"Айзек Азимов"},
		Publishers: []string{"АСТ"},
		Genres:     []string{"Фантастика"},
		Amount:     5,
		Date:       "2020-10-5",
	}
	//fmt.Println(checks.CheckBook(d))
	res, _ := json.Marshal(d)
	r, _ := http.NewRequest("POST", "http://localhost:2406/api/book", bytes.NewBuffer(res))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddBook)
	r.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, r)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	t.Logf("Handler returned following status code: %v, book id: %s", rr.Result().Status, string(body))
}
