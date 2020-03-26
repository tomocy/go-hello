package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(sayHello))
	defer s.Close()

	resp, err := http.Get(s.URL)
	if err != nil {
		t.Errorf("should have requested with GET: %s", err)
		return
	}
	defer resp.Body.Close()

	expectedStatusCode := 200
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("should have got the expected HTTP status: expected %d, but got %d", expectedStatusCode, resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("should have read all of body: %s", err)
		return
	}

	expectedBody := "Ver 2: Hello!\n"
	if string(body) != expectedBody {
		t.Errorf("should have responsed the expected body: expected %q, but got %q", expectedBody, body)
	}
}
