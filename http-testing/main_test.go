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
	tests := []struct {
		desc       string
		in         url.Values
		wantErr    bool
		wantBody   string
		wantStatus int
	}{
		{
			desc:       "returns a status OK and the correct body when person params is supplied",
			in:         url.Values{"person": {"john"}},
			wantStatus: http.StatusOK,
			wantBody:   "Hello, john!",
		},
		{
			desc:       "returns a status BadRequest when no query parameter is supplied",
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/hello", nil)
			req.Form = test.in
			if err != nil {
				t.Fatal(err)
			}

			res := httptest.NewRecorder()
			handler := http.HandlerFunc(helloHandler)

			handler.ServeHTTP(res, req)

			if got := res.Code; got != test.wantStatus {
				t.Errorf("expecting status: got %d, want %d", got, test.wantStatus)
			}

			got := res.Body.String()

			// Successful HTTP requests.
			if !test.wantErr && got != test.wantBody {
				t.Errorf("expecting body: got %q, want %q", got, test.wantBody)
			}
		})
	}
}
