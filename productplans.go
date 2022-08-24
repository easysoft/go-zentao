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

type ProductPlansService struct {
	client *Client
}

type ProductPlanMeta struct {
	Branch string `json:"branch,omitempty"`
	Title  string `json:"title"`
	Parent int    `json:"parent,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Begin  string `json:"begin,omitempty"`
	End    string `json:"end,omitempty"`
}

type ProductPlanCreateMsg struct {
	ID           int           `json:"id"`
	Product      int           `json:"product"`
	Branch       int           `json:"branch"`
	Parent       int           `json:"parent"`
	Title        string        `json:"title"`
	Status       string        `json:"status"`
	Desc         string        `json:"desc"`
	Begin        string        `json:"begin"`
	End          string        `json:"end"`
	Order        string        `json:"order"`
	ClosedReason string        `json:"closedReason"`
	Deleted      bool          `json:"deleted"`
	Stories      []interface{} `json:"stories"`
	Bugs         []interface{} `json:"bugs"`
}

type ProductPlanListMsg struct {
	Page  int `json:"page"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Plans []struct {
		ID           int    `json:"id"`
		Product      int    `json:"product"`
		Branch       int    `json:"branch"`
		Parent       int    `json:"parent"`
		Title        string `json:"title"`
		Status       string `json:"status"`
		Desc         string `json:"desc"`
		Begin        string `json:"begin"`
		End          string `json:"end"`
		Order        string `json:"order"`
		ClosedReason string `json:"closedReason"`
		Deleted      bool   `json:"deleted"`
		Stories      int    `json:"stories"`
		Bugs         int    `json:"bugs"`
		Hour         int    `json:"hour"`
		Project      int    `json:"project"`
		ProjectID    string `json:"projectID"`
		Expired      bool   `json:"expired"`
	} `json:"plans"`
}

// Create 创建产品计划
func (s *ProductPlansService) Create(id int, plan ProductPlanMeta) (*ProductPlanCreateMsg, *req.Response, error) {
	var u ProductPlanCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&plan).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/plans", id)))
	return &u, resp, err
}

// GetByID 获取产品计划列表
func (s *ProductPlansService) List(id int) (*ProductPlanListMsg, *req.Response, error) {
	var u ProductPlanListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/plans", id)))
	return &u, resp, err
}
