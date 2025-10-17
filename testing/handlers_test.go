package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/timrxd/exhibit-backend/proto/players"
)

func TestAPI(t *testing.T) {

	cases := []struct {
		method string
		path   string
		body   pb.PlayerReq

		outCode int
		outBody string
	}{
		// Add first user to address book
		{"POST", "/users", pb.PlayerReq{Name: "Federer"}, 200, `{"player":{"name":"Federer", "country":"", "age":0, "points":0}`},
	}

	for _, c := range cases {
		response := httptest.NewRecorder()
		body, _ := json.Marshal(c.body)
		request, _ := http.NewRequest(c.method, c.path, bytes.NewBuffer(body))
		AddressRouter().ServeHTTP(response, request)

		b := response.Body.String()
		if b[:len(b)-1] != c.outBody {
			t.Errorf("Body Mismatch: %s %s\nResponse\t%s\nExpected\t%s", c.method, c.path, b[:len(b)-1], c.outBody)
		} else if response.Code != c.outCode {
			t.Errorf("Response Code Mismatch: %s %s\nResponse\t%s\nExpected\t%s", c.method, c.path, response.Code, c.outCode)
		}
	}
}
