//
//  Copyright 2024, zentao
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

	"github.com/imroc/req/v3"
)

type ModulesService struct {
	client *Client
}

type ModulesListOptions struct {
	Type ModuleType `json:"type"`
	ID   int        `json:"id"`
}

func (o ModulesListOptions) toURLValues() (string, error) {
	v := url.Values{}
	if len(o.Type) == 0 {
		return "", fmt.Errorf("Type is required")
	}
	if o.ID == 0 {
		return "", fmt.Errorf("ID is required")
	}
	v.Set("type", fmt.Sprintf("%v", o.Type))
	v.Set("id", fmt.Sprintf("%v", o.ID))
	return v.Encode(), nil
}

type ListModules struct {
	Modules []Module `json:"modules"`
}

type Module struct {
	ID       int        `json:"id,omitempty"` // 子模块ID
	Name     string     `json:"name"`
	Root     int        `json:"root"`
	Type     ModuleType `json:"type"`
	Actions  bool       `json:"actions,omitempty"` // 父模块字段
	Grade    int        `json:"grade,omitempty"`   // 子模块字段
	Parent   int        `json:"parent,omitempty"`  // 子模块字段
	URL      string     `json:"url,omitempty"`     // 子模块字段
	Path     string     `json:"path,omitempty"`    // 子模块字段
	Children []Module   `json:"children,omitempty"`
}

func (s *ModulesService) List(op ModulesListOptions) (*ListModules, *req.Response, error) {
	var et ListModules
	opValues, err := op.toURLValues()
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryString(opValues).
		SetSuccessResult(&et).
		Get(s.client.RequestURL("/modules"))
	return &et, resp, err
}
