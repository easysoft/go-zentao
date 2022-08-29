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

	"github.com/imroc/req/v3"
)

type TestTasksService struct {
	client *Client
}

type ListTestTasksMsg struct {
	Page      int `json:"page"`
	Total     int `json:"total"`
	Limit     int `json:"limit"`
	Testtasks []struct {
		ID               int         `json:"id"`
		Project          int         `json:"project"`
		Product          int         `json:"product"`
		Name             string      `json:"name"`
		Execution        int         `json:"execution"`
		Build            string      `json:"build"`
		Type             string      `json:"type"`
		Owner            string      `json:"owner"`
		Pri              int         `json:"pri"`
		Begin            string      `json:"begin"`
		End              string      `json:"end"`
		Realfinisheddate interface{} `json:"realFinishedDate"`
		Mailto           string      `json:"mailto"`
		Desc             string      `json:"desc"`
		Report           string      `json:"report"`
		Status           string      `json:"status"`
		Testreport       int         `json:"testreport"`
		Auto             string      `json:"auto"`
		Substatus        string      `json:"subStatus"`
		Deleted          string      `json:"deleted"`
		Productname      string      `json:"productName"`
		Executionname    string      `json:"executionName"`
		Buildname        string      `json:"buildName"`
		Branch           int         `json:"branch"`
	} `json:"testtasks"`
}

type TestTasksGetMsg struct {
	ID               int         `json:"id"`
	Project          int         `json:"project"`
	Product          int         `json:"product"`
	Name             string      `json:"name"`
	Execution        int         `json:"execution"`
	Build            string      `json:"build"`
	Type             string      `json:"type"`
	Owner            string      `json:"owner"`
	Pri              int         `json:"pri"`
	Begin            string      `json:"begin"`
	End              string      `json:"end"`
	Realfinisheddate interface{} `json:"realFinishedDate"`
	Mailto           string      `json:"mailto"`
	Desc             string      `json:"desc"`
	Report           string      `json:"report"`
	Status           string      `json:"status"`
	Testreport       int         `json:"testreport"`
	Auto             string      `json:"auto"`
	Substatus        string      `json:"subStatus"`
	Deleted          string      `json:"deleted"`
	Productname      string      `json:"productName"`
	Executionname    string      `json:"executionName"`
	Buildname        string      `json:"buildName"`
	Branch           int         `json:"branch"`
}

// List 获取测试单列表
func (s *TestTasksService) List(page, limit, order, product, branch string) (*ListTestTasksMsg, *req.Response, error) {
	var et ListTestTasksMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"page":    page,
			"limit":   limit,
			"order":   order,
			"product": product,
			"branch":  branch,
		}).
		SetResult(&et).
		Get(s.client.RequestURL("/testtasks"))
	return &et, resp, err
}

func (s *TestTasksService) ListByProjects(id int64) (*ListTestTasksMsg, *req.Response, error) {
	var et ListTestTasksMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/testtasks", id)))
	return &et, resp, err
}

func (s *TestTasksService) GetByID(id int) (*TestTasksGetMsg, *req.Response, error) {
	var u TestTasksGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/testtasks/%d", id)))
	return &u, resp, err
}
