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

package buildin

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/imroc/req/v3"
)

// @title Zentao Go SDK
// @version v21.0
// @description This is a Go SDK for Zentao API.

// @schemes https
// @host zentao.demo.qucheng.cc
// @BasePath /

const (
	defaultBaseURL = "https://zentao.demo.qucheng.cc"
	userAgent      = "go-zentao"
	apiVersionPath = "api.php/v1"
)

type ZentaoMode string

func (z ZentaoMode) String() string {
	return string(z)
}

const (
	ZentaoDefultMode ZentaoMode = "PATH_INFO"
	ZentaoGetMode    ZentaoMode = "GET"
)

type Client struct {
	client  *req.Client
	baseURL *url.URL

	// Username and password used for basic authentication.
	username, password string

	// zentao session id
	zentaosid string

	zentaomode ZentaoMode

	// build-in
	Login *LoginService
}

func NewBasicAuthClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	sessionID, _, err := client.Login.GetSessionID()
	if err != nil {
		return nil, err
	}
	if sessionID == nil {
		return nil, fmt.Errorf("token not valid")
	}
	client.zentaosid = *sessionID
	// auth password
	client.username = username
	client.password = password
	return client, nil
}

func newClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{}
	c.client = req.C().SetLogger(nil)
	c.setBaseURL(defaultBaseURL)
	c.setReqUserAgent(userAgent)
	c.setZenTaoMode(ZentaoDefultMode)
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}
	c.Login = &LoginService{client: c}
	return c, nil
}

func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash

	urlStr = strings.TrimSuffix(urlStr, "/")

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	baseURL.Path = strings.TrimSuffix(baseURL.Path, apiVersionPath)

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

func (c *Client) setDebug() error {
	c.client.EnableDebugLog()
	return nil
}

func (c *Client) setZenTaoMode(m ZentaoMode) error {
	c.zentaomode = m
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

func (c *Client) RequestURLFmt(f string, v ...any) string {
	u := *c.baseURL
	u.Path = c.baseURL.Path + fmt.Sprintf(f, v...)
	return u.String()
}
