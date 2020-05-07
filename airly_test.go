package airly

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client, err := NewClient(nil, "apiKey")
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}
	client.baseURL, _ = url.Parse(server.URL)

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("request method: %v, want %v", got, want)
	}
}

func testFormValues(t *testing.T, r *http.Request, values url.Values) {
	t.Helper()
	if err := r.ParseForm(); err != nil {
		t.Errorf("ParseForm: %v", err)
	}
	if got := r.Form; !reflect.DeepEqual(got, values) {
		t.Errorf("request parameters: %v, want %v", got, values)
	}
}
