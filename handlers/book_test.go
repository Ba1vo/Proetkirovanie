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

func TestGetBook(t *testing.T) {
	d := 30
	//fmt.Println(checks.CheckBook(d))
	res, _ := json.Marshal(d) //ds
	r, _ := http.NewRequest("POST", "http://localhost:2406/api/book", bytes.NewBuffer(res))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Book)
	r.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, r)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	var result decoder.FullBook
	json.Unmarshal(body, &result)
	t.Logf("Handler returned following status code: %v, book : %v", rr.Result().Status, result)
}
