package zentao

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// setup copy form gitlab https://github.com/xanzy/go-gitlab/blob/4256ad62c1044a01b0eab2fd863f53807d562876/gitlab_test.go#L57
func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, err := NewClient("", WithBaseURL(server.URL))
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create client: %v", err)
	}
	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}
