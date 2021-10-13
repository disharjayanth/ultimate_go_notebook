package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/disharjayanth/ultimate_go_notebook/chap07/04_webEndpointTest/handlers"
)

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	r := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != statusCode {
		t.Log("exp:", statusCode)
		t.Log("got:", w.Code)
		t.Fatal("status does'nt match")
	}

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		t.Fatal("unable to decode the response")
	}

	name := "Bill"
	if u.Name != name {
		t.Log("exp:", name)
		t.Log("got:", u.Name)
		t.Fatal("user name does not match")
	}

	email := "bill@ardanlabs.com"
	if u.Email != email {
		t.Log("exp:", email)
		t.Log("got:", u.Email)
		t.Fatal("email does not match")
	}
}
