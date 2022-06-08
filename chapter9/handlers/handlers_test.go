package handlers_test

import (
	"encoding/json"
	"gia/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the sendJSON endpoint")
	{
		req, err := http.NewRequest("GET", "/sendJSON", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request", ballotX, err)
		}

		t.Log("\tShould be able to create a request", checkMark)

		rec := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatal("\tShould receive status code \"200\"", ballotX, rec.Code)
		} else {
			t.Log("\tShould receive status code \"200\"", checkMark)
		}

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rec.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response", ballotX)
		}

		t.Log("\tShould decode the response", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name", checkMark)
		} else {
			t.Error("\tShould have a Name", ballotX, u.Name)
		}

		if u.Email == "bill@clifton.org" {
			t.Log("\tShould have an Email", checkMark)
		} else {
			t.Error("\tShould have an Email", ballotX, u.Email)
		}
	}
}