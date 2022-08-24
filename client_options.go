//
//  Copyright 2022, ysicing
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

type ClientOptionFunc func(*Client) error

// WithBaseURL sets the base URL for API requests to a custom endpoint.
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

func WithDevMode() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDebug()
	}
}

func WithDumpAll() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDumpAll()
	}
}

// WithoutProxy 禁用代理, 默认情况下会读取HTTP_PROXY/HTTPS_PROXY/http_proxy/https_proxy变量
func WithoutProxy() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDisableProxy()
	}
}
