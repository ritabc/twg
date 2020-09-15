package stripe

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Takes testing.T. Returns stripe.Client, serveMux used when hitting API endpoints, and teardown function
func TestClient(t *testing.T) (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	c := &Client{
		baseURL: server.URL,
	}
	return c, mux, func() {
		server.Close()
	}
}
