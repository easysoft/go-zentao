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
	"time"

	"github.com/imroc/req/v3"
)

type ProductsService struct {
	client *Client
}

type ProductsList struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	Products []struct {
		ID             int        `json:"id"`
		Program        int        `json:"program"`
		Name           string     `json:"name"`
		Code           string     `json:"code"`
		Bind           string     `json:"bind"`
		Line           int        `json:"line"`
		Type           string     `json:"type"`
		Status         string     `json:"status"`
		Substatus      string     `json:"subStatus"`
		Desc           string     `json:"desc"`
		Po             UserMeta   `json:"PO"`
		Qd             UserMeta   `json:"QD"`
		Rd             UserMeta   `json:"RD"`
		ACL            string     `json:"acl"`
		Whitelist      []UserMeta `json:"whitelist"`
		Reviewer       string     `json:"reviewer"`
		Createdby      UserMeta   `json:"createdBy"`
		Createddate    time.Time  `json:"createdDate"`
		Createdversion string     `json:"createdVersion"`
		Order          int        `json:"order"`
		Deleted        string     `json:"deleted"`
		Linename       string     `json:"lineName"`
		Programname    string     `json:"programName"`
		Stories        struct {
			Num0    string `json:"0"`
			Num1    string `json:"1"`
			Num2    string `json:"2"`
			Num3    string `json:"3"`
			Num4    string `json:"4"`
			int     `json:""`
			Draft   int `json:"draft"`
			Active  int `json:"active"`
			Closed  int `json:"closed"`
			Changed int `json:"changed"`
		} `json:"stories"`
		Requirements struct {
			Num0    string `json:"0"`
			Num1    string `json:"1"`
			Num2    string `json:"2"`
			Num3    string `json:"3"`
			Num4    string `json:"4"`
			int     `json:""`
			Draft   int `json:"draft"`
			Active  int `json:"active"`
			Closed  int `json:"closed"`
			Changed int `json:"changed"`
		} `json:"requirements"`
		Plans        int `json:"plans"`
		Releases     int `json:"releases"`
		Bugs         int `json:"bugs"`
		Unresolved   int `json:"unResolved"`
		Closedbugs   int `json:"closedBugs"`
		Fixedbugs    int `json:"fixedBugs"`
		Thisweekbugs int `json:"thisWeekBugs"`
		Assigntonull int `json:"assignToNull"`
		Progress     int `json:"progress"`
	} `json:"products"`
}

type ProductsCreateMeta struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ProductsCreateMsg struct {
	ID             int           `json:"id"`
	Program        int           `json:"program"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Bind           string        `json:"bind"`
	Line           int           `json:"line"`
	Type           string        `json:"type"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Desc           string        `json:"desc"`
	Po             interface{}   `json:"PO"`
	Qd             interface{}   `json:"QD"`
	Rd             interface{}   `json:"RD"`
	Feedback       string        `json:"feedback"`
	ACL            string        `json:"acl"`
	Whitelist      []interface{} `json:"whitelist"`
	Reviewer       string        `json:"reviewer"`
	Createdby      UserMeta      `json:"createdBy"`
	Createddate    time.Time     `json:"createdDate"`
	Createdversion string        `json:"createdVersion"`
	Order          int           `json:"order"`
	Vision         string        `json:"vision"`
	Deleted        string        `json:"deleted"`
}

type ProductsGetMsg struct {
	ID             int           `json:"id"`
	Program        int           `json:"program"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Bind           string        `json:"bind"`
	Line           int           `json:"line"`
	Type           string        `json:"type"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Desc           string        `json:"desc"`
	Po             interface{}   `json:"PO"`
	Qd             interface{}   `json:"QD"`
	Rd             interface{}   `json:"RD"`
	Feedback       interface{}   `json:"feedback"`
	ACL            string        `json:"acl"`
	Whitelist      []interface{} `json:"whitelist"`
	Reviewer       string        `json:"reviewer"`
	Createdby      UserMeta      `json:"createdBy"`
	Createddate    time.Time     `json:"createdDate"`
	Createdversion string        `json:"createdVersion"`
	Order          int           `json:"order"`
	Vision         string        `json:"vision"`
	Deleted        string        `json:"deleted"`
	Stories        struct {
		int     `json:""`
		Draft   int `json:"draft"`
		Active  int `json:"active"`
		Closed  int `json:"closed"`
		Changed int `json:"changed"`
	} `json:"stories"`
	Plans      int  `json:"plans"`
	Releases   int  `json:"releases"`
	Builds     int  `json:"builds"`
	Cases      int  `json:"cases"`
	Projects   int  `json:"projects"`
	Executions int  `json:"executions"`
	Bugs       int  `json:"bugs"`
	Docs       int  `json:"docs"`
	Progress   int  `json:"progress"`
	Casereview bool `json:"caseReview"`
}

type ProductsUpdateMeta struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ProductsUpdateMsg struct {
	ID             int           `json:"id"`
	Program        int           `json:"program"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Bind           string        `json:"bind"`
	Line           int           `json:"line"`
	Type           string        `json:"type"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Desc           string        `json:"desc"`
	Po             UserMeta      `json:"PO"`
	Qd             UserMeta      `json:"QD"`
	Rd             UserMeta      `json:"RD"`
	ACL            string        `json:"acl"`
	Whitelist      []interface{} `json:"whitelist"`
	Reviewer       string        `json:"reviewer"`
	Createdby      UserMeta      `json:"createdBy"`
	Createddate    time.Time     `json:"createdDate"`
	Createdversion string        `json:"createdVersion"`
	Order          int           `json:"order"`
	Deleted        string        `json:"deleted"`
	Stories        struct {
		Active  int `json:"active"`
		Draft   int `json:"draft"`
		int     `json:""`
		Closed  int `json:"closed"`
		Changed int `json:"changed"`
	} `json:"stories"`
	Plans      int  `json:"plans"`
	Releases   int  `json:"releases"`
	Builds     int  `json:"builds"`
	Cases      int  `json:"cases"`
	Projects   int  `json:"projects"`
	Executions int  `json:"executions"`
	Bugs       int  `json:"bugs"`
	Docs       int  `json:"docs"`
	Progress   int  `json:"progress"`
	Casereview bool `json:"caseReview"`
}

// List 获取产品列表
func (s *ProductsService) List() (*ProductsList, *req.Response, error) {
	var u ProductsList
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL("/products"))
	return &u, resp, err
}

// Create 创建产品
func (s *ProductsService) Create(program ProductsCreateMeta) (*ProductsCreateMsg, *req.Response, error) {
	var u ProductsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetResult(&u).
		Post(s.client.RequestURL("/products"))
	return &u, resp, err
}

// DeleteByID 删除产品
func (s *ProductsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}

// UpdateByID 更新产品
func (s *ProductsService) UpdateByID(id int, user ProductsUpdateMeta) (*ProductsUpdateMsg, *req.Response, error) {
	var u ProductsUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}

// GetByID 获取产品详情
func (s *ProductsService) GetByID(id int) (*ProductsGetMsg, *req.Response, error) {
	var u ProductsGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}
