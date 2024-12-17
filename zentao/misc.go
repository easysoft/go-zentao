//
//  Copyright 2024, easysoft
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

type MiscService struct {
	client *Client
}

type Configurations []Configuration

type Configuration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *MiscService) GetConfigurations() (*Configurations, *req.Response, error) {
	var result Configurations
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&result).
		Get(s.client.RequestURL("/configurations"))
	return &result, resp, err
}

func (s *MiscService) GetVersion() (*Configuration, *req.Response, error) {
	var result Configuration
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&result).
		Get(s.client.RequestURL("/configurations/version"))
	return &result, resp, err
}
