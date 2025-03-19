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

func TestClient_setBaseURL(t *testing.T) {
	tests := []struct {
		name    string
		urlStr  string
		want    string
		wantErr bool
	}{
		{
			name:    "基础 URL",
			urlStr:  "https://zentao.demo.qucheng.cc",
			want:    "https://zentao.demo.qucheng.cc/api.php/v1",
			wantErr: false,
		},
		{
			name:    "已包含 api.php/v1",
			urlStr:  "https://zentao.demo.qucheng.cc/api.php/v1",
			want:    "https://zentao.demo.qucheng.cc/api.php/v1",
			wantErr: false,
		},
		{
			name:    "已包含斜杠结尾",
			urlStr:  "https://zentao.demo.qucheng.cc/",
			want:    "https://zentao.demo.qucheng.cc/api.php/v1",
			wantErr: false,
		},
		{
			name:    "完整 URL 带斜杠",
			urlStr:  "https://zentao.demo.qucheng.cc/api.php/v1/",
			want:    "https://zentao.demo.qucheng.cc/api.php/v1",
			wantErr: false,
		},
		{
			name:    "无效 URL",
			urlStr:  "://invalid-url",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{}
			err := c.setBaseURL(tt.urlStr)

			// 检查错误情况
			if (err != nil) != tt.wantErr {
				t.Errorf("setBaseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 如果期望发生错误，不需要检查后续结果
			if tt.wantErr {
				return
			}

			// 检查生成的 URL 是否符合预期
			if got := c.baseURL.String(); got != tt.want {
				t.Errorf("setBaseURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
