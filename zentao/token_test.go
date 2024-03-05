//
//  Copyright 2022, easysoft
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

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
