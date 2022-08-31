package zentao

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)
	mux.HandleFunc("/api.php/v1/tokens", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"token": "cuejkiesahl9k1j8be5bv5lndo"}`)
	})
	expected := &AccessToken{
		Token: "cuejkiesahl9k1j8be5bv5lndo",
	}
	unexpected := &AccessToken{
		Token: "cuejkiesahl9k1j8be5bv5lnd0",
	}
	requests, resp, err := client.Token.GetAccessToken()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, requests)
	assert.NotEqual(t, unexpected, requests)
}
