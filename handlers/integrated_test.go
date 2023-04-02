package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ba1vo/Proektirovanie/crypt"
	"github.com/Ba1vo/Proektirovanie/decoder"
	"github.com/Ba1vo/Proektirovanie/queries"
)

func TestIntegrated(t *testing.T) {

	d := decoder.RegUser{
		Name:  "Joka",
		Email: "2002egor@gmail.com",
		Pass:  "15925555",
	}

	res, _ := json.Marshal(d)
	r, _ := http.NewRequest("POST", "http://localhost:2406/api/register", bytes.NewBuffer(res))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	r.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(rr, r)
	t.Logf("Handler returned following status code: %v", rr.Result().Status)
	tempUser := decoder.AuthUser{
		Email: d.Email,
		Pass:  crypt.Hash(d.Pass),
	}
	user, _ := queries.GetUser(tempUser)
	queries.DeleteUser(user.ID)
}
