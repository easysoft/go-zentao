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
	"time"

	"github.com/imroc/req/v3"
)

type ProductsService struct {
	client *Client
}

type ProductsListBody struct {
	ProductsMeta
	ProductsBody

	Linename     string          `json:"lineName"`
	Programname  string          `json:"programName"`
	Stories      ProductsStories `json:"stories"`
	Requirements ProductsStories `json:"requirements,omitempty"`
	Plans        int             `json:"plans"`
	Releases     int             `json:"releases"`
	Bugs         int             `json:"bugs"`
	Unresolved   int             `json:"unResolved"`
	Closedbugs   int             `json:"closedBugs"`
	Fixedbugs    int             `json:"fixedBugs"`
	Thisweekbugs int             `json:"thisWeekBugs"`
	Assigntonull int             `json:"assignToNull"`
	Progress     float64         `json:"progress"`
}

type ProductsList struct {
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	Limit    int                `json:"limit"`
	Products []ProductsListBody `json:"products"`
}

type ProductsMeta struct {
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Line    int    `json:"line,omitempty"`
	Program int    `json:"program,omitempty"`
	Status  string `json:"status,omitempty"`
	Desc    string `json:"desc,omitempty"`
}

type ProductsBody struct {
	ID             int        `json:"id"`
	Bind           string     `json:"bind"`
	Substatus      string     `json:"subStatus,omitempty"`
	Po             UserMeta   `json:"PO,omitempty"`
	Qd             UserMeta   `json:"QD,omitempty"`
	Rd             UserMeta   `json:"RD,omitempty"`
	ACL            ACL        `json:"acl"`
	Whitelist      []UserMeta `json:"whitelist"`
	Reviewer       string     `json:"reviewer,omitempty"`
	Createdby      UserMeta   `json:"createdBy"`
	Createddate    time.Time  `json:"createdDate"`
	Createdversion string     `json:"createdVersion,omitempty"`
	Order          int        `json:"order"`
	Deleted        string     `json:"deleted"`
}

type ProductsCreateMsg struct {
	ProductsMeta
	ProductsBody
	Feedback string `json:"feedback,omitempty"`
	Vision   string `json:"vision,omitempty"`
}

type ProductsGetMsg struct {
	ProductsMeta
	Feedback string `json:"feedback,omitempty"`
	Vision   string `json:"vision,omitempty"`

	Stories ProductsStories `json:"stories,omitempty"`
	ProductsExtMsg
}

type ProductsUpdateMsg struct {
	ProductsMeta
	ProductsBody

	Stories ProductsStories `json:"stories"`
	ProductsExtMsg
}

type ProductsExtMsg struct {
	Plans      int  `json:"plans,omitempty"`
	Releases   int  `json:"releases,omitempty"`
	Builds     int  `json:"builds,omitempty"`
	Cases      int  `json:"cases,omitempty"`
	Projects   int  `json:"projects,omitempty"`
	Executions int  `json:"executions,omitempty"`
	Bugs       int  `json:"bugs,omitempty"`
	Docs       int  `json:"docs,omitempty"`
	Progress   int  `json:"progress,omitempty"`
	Casereview bool `json:"caseReview,omitempty"`
}

type ProductsStories struct {
	Num0    string `json:"0,omitempty"`
	Num1    string `json:"1,omitempty"`
	Num2    string `json:"2,omitempty"`
	Num3    string `json:"3,omitempty"`
	Num4    string `json:"4,omitempty"`
	Active  int    `json:"active,omitempty"`
	Draft   int    `json:"draft,omitempty"`
	Unknow  int    `json:",omitempty"` // 空 {"": 0} ???
	Closed  int    `json:"closed,omitempty"`
	Changed int    `json:"changed,omitempty"`
}

// List 获取产品列表
func (s *ProductsService) List() (*ProductsList, *req.Response, error) {
	var u ProductsList
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL("/products"))
	return &u, resp, err
}

// Create 创建产品
func (s *ProductsService) Create(program ProductsMeta) (*ProductsCreateMsg, *req.Response, error) {
	var u ProductsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetSuccessResult(&u).
		Post(s.client.RequestURL("/products"))
	return &u, resp, err
}

// DeleteByID 删除产品
func (s *ProductsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}

// UpdateByID 更新产品
func (s *ProductsService) UpdateByID(id int, user ProductsMeta) (*ProductsUpdateMsg, *req.Response, error) {
	var u ProductsUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}

// GetByID 获取产品详情
func (s *ProductsService) GetByID(id int) (*ProductsGetMsg, *req.Response, error) {
	var u ProductsGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d", id)))
	return &u, resp, err
}
