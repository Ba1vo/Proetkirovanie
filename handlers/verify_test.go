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

func TestVerify(t *testing.T) {
	d := decoder.FavBook{ //?? array or struct?
		User: 5,
		Book: 30,
	}
	//fmt.Println(checks.CheckBook(d))
	res, _ := json.Marshal(d)
	r, _ := http.NewRequest("POST", "http://localhost:2406/api/book", bytes.NewBuffer(res))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteOrder)
	r.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, r)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	t.Logf("Handler returned following status code: %v, book id: %s", rr.Result().Status, string(body))
}
