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
	"net/url"
	"strings"

	"github.com/imroc/req/v3"
)

// @title Zentao Go SDK
// @version v20.0
// @description This is a Go SDK for Zentao API.

// @schemes https
// @host zentao.demo.qucheng.cc
// @BasePath /api.php/v1

const (
	defaultBaseURL = "https://zentao.demo.qucheng.cc/"
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

	// restful api
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
	Builds       *BuildsService
	Bugs         *BugsService
	TestCases    *TestCasesService
	TestTasks    *TestTasksService
	FeedBacks    *FeedBacksService
	Modules      *ModulesService
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
	c := &Client{}
	c.client = req.C().SetLogger(nil)
	c.setBaseURL(defaultBaseURL)
	c.setReqUserAgent(userAgent)
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
	c.Builds = &BuildsService{client: c}
	c.Bugs = &BugsService{client: c}
	c.TestCases = &TestCasesService{client: c}
	c.TestTasks = &TestTasksService{client: c}
	c.FeedBacks = &FeedBacksService{client: c}
	c.Modules = &ModulesService{client: c}
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

func (c *Client) setReqUserAgent(ua string) error {
	c.client.SetUserAgent(ua)
	return nil
}

// RequestURL restful api request url
func (c *Client) RequestURL(path string) string {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path
	return u.String()
}

// ListOptions specifies the optional parameters to various List methods that
type ListOptions struct {
	Page  int `url:"page,omitempty" json:"page,omitempty"`
	Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	// Status string `url:"status,omitempty" json:"status,omitempty"` // 目前不生效 NeedFixed
}

func (o ListOptions) toURLValues() string {
	v := url.Values{}
	if o.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", o.Page))
	}
	if o.Limit > 0 {
		v.Set("limit", fmt.Sprintf("%d", o.Limit))
	}
	return v.Encode()
}
