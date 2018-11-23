package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPingHandler(t *testing.T) {
	// Create a request for our handler,
	// pass nil as the third parameter because we don't require any query parameters.
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Used to record the response of the HTTP call.
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	// Satisfy http.Handler interface, so we can call http's ServeHTTP method.
	handler.ServeHTTP(res, req)

	wantStatus := http.StatusOK // 200
	if got := res.Code; got != wantStatus {
		t.Errorf("expecting status: got %v, want %v", got, wantStatus)
	}

	wantBody := "pong"
	if got := res.Body.String(); got != wantBody {
		t.Errorf("expecting body: got %v, want %v", got, wantBody)
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	req.Form = url.Values{"person": {"john"}}
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(res, req)

	wantStatus := http.StatusOK
	if got := res.Code; got != wantStatus {
		t.Errorf("expecting status: got %v, want %v", got, wantStatus)
	}

	wantBody := "Hello, john!"
	if got := res.Body.String(); got != wantBody {
		t.Errorf("expecting body: got %v, want %v", got, wantBody)
	}
}
