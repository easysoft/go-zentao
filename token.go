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
	"github.com/imroc/req/v3"
)

type TokenService struct {
	client *Client
}

type AccessToken struct {
	Token string `json:"token"`
}

type BasicAuth struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (s *TokenService) GetAccessToken() (*AccessToken, *req.Response, error) {
	var result AccessToken
	resp, err := s.client.client.R().
		SetBody(&BasicAuth{Account: s.client.username, Password: s.client.password}).
		SetSuccessResult(&result).
		Post(s.client.RequestURL("/tokens"))

	return &result, resp, err
}
