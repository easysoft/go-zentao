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

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/imroc/req/v3"
)

const (
	defaultBaseURL = "https://demo16.zentao.net/"
	userAgent      = "go-zentao"
	apiVersionPath = "api.php/v1"
)

type Client struct {
	client  *req.Client
	baseURL *url.URL

	// Username and password used for basic authentication.
	username, password string

	// Token used to make authenticated API calls.
	token string

	UserAgent string

	Token        *TokenService
	Users        *UsersService
	Executions   *ExecutionsService
	Tasks        *TasksService
	Programs     *ProgramsService
	Products     *ProductsService
	ProductPlans *ProductPlansService
	Releases     *ReleasesService
	Stories      *StoriesService
	Projects     *ProjectsService
}

func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	client.token = token
	return client, nil
}

func NewBasicAuthClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	client.username = username
	client.password = password
	accesstoken, _, err := client.Token.GetAccessToken()
	if err != nil {
		return nil, err
	}
	if accesstoken != nil {
		client.token = accesstoken.Token
		return client, nil
	}
	return nil, fmt.Errorf("token not valid")
}

func newClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{UserAgent: userAgent}
	c.client = req.C()
	c.setBaseURL(defaultBaseURL)
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}
	c.Token = &TokenService{client: c}
	c.Users = &UsersService{client: c}
	c.Executions = &ExecutionsService{client: c}
	c.Tasks = &TasksService{client: c}
	c.Programs = &ProgramsService{client: c}
	c.Products = &ProductsService{client: c}
	c.ProductPlans = &ProductPlansService{client: c}
	c.Releases = &ReleasesService{client: c}
	c.Stories = &StoriesService{client: c}
	c.Projects = &ProjectsService{client: c}
	return c, nil
}

func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

func (c *Client) setDebug() error {
	c.client.EnableDebugLog()
	return nil
}

func (c *Client) setDumpAll() error {
	c.client.EnableDumpAll()
	return nil
}

func (c *Client) setDisableProxy() error {
	c.client.SetProxy(nil)
	return nil
}

func (c *Client) RequestURL(path string) string {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path
	return u.String()
}
