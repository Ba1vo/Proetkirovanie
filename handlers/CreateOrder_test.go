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

func TestCreateOrd(t *testing.T) {
	d := decoder.ServerOrder{
		UserID:   23,
		Adress:   "Пушкина",
		Status:   "Ожидает",
		Book_IDs: []int{38, 52},
		Amounts:  []int{4, 10},
	}
	//fmt.Println(checks.CheckBook(d))
	res, _ := json.Marshal(d)
	r, _ := http.NewRequest("POST", "http://localhost:2406/api/createOrder", bytes.NewBuffer(res))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrder)
	r.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, r)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	t.Logf("Handler returned following status code: %v, book id: %s", rr.Result().Status, string(body))
}
