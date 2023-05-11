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

	"github.com/imroc/req/v3"
)

type ProductPlansService struct {
	client *Client
}

type ProductPlanMeta struct {
	Branch int    `json:"branch,omitempty"`
	Title  string `json:"title"`
	Parent int    `json:"parent,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Begin  string `json:"begin,omitempty"`
	End    string `json:"end,omitempty"`
}

type ProductPlanBody struct {
	ID      int    `json:"id"`
	Product int    `json:"product"`
	Order   string `json:"order"`
	Deleted bool   `json:"deleted"`
	ProductPlanMeta
}

type ProductPlanCreateMsg struct {
	ProductPlanBody

	Status       string        `json:"status"`
	ClosedReason string        `json:"closedReason"`
	Stories      []StoriesBody `json:"stories"`
	Bugs         []BugBody     `json:"bugs"`
}

type ProductPlanListBody struct {
	ProductPlanBody
	Status       string `json:"status,omitempty"`
	ClosedReason string `json:"closedReason"`
	Stories      int    `json:"stories"`
	Bugs         int    `json:"bugs"`
	Hour         int    `json:"hour"`
	Project      int    `json:"project"`
	ProjectID    string `json:"projectID"`
	Expired      bool   `json:"expired"`
}

type ProductPlanListMsg struct {
	Page  int                   `json:"page"`
	Total int                   `json:"total"`
	Limit int                   `json:"limit"`
	Plans []ProductPlanListBody `json:"plans"`
}

type ProductPlanGetMsg struct {
	ProductPlanBody
	Stories []StoriesBody `json:"stories"`
	Bugs    []BugBody     `json:"bugs"`
}

type ProductPlanUpdateMsg struct {
	ProductPlanBody
	Stories []StoriesBody `json:"stories"`
	Bugs    []BugBody     `json:"bugs"`
}

type PlansStoriesIDs struct {
	Stories []string `json:"stories"`
}

type PlansBugIDs struct {
	Bugs []string `json:"bugs"`
}

type LinkStoriesMsg struct {
	ProductPlanBody
	Stories []StoriesBody `json:"stories"`
	Bugs    []BugBody     `json:"bugs"`
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

// List 获取产品计划列表
func (s *ProductPlansService) List(id int) (*ProductPlanListMsg, *req.Response, error) {
	var u ProductPlanListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/plans", id)))
	return &u, resp, err
}

// GetByID 获取计划详情
func (s *ProductPlansService) GetByID(id int) (*ProductPlanGetMsg, *req.Response, error) {
	var u ProductPlanGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改计划
func (s *ProductPlansService) UpdateByID(id int, plan ProductPlanMeta) (*ProductPlanUpdateMsg, *req.Response, error) {
	var u ProductPlanUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&plan).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// DeleteByID 删除计划
func (s *ProductPlansService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// LinkStories 产品计划关联需求
func (s *ProductPlansService) LinkStories(id int, story PlansStoriesIDs) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/linkstories", id)))
	return &u, resp, err
}

// LinkStories 产品计划关联需求
func (s *ProductPlansService) UnLinkStories(id int, story PlansStoriesIDs) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/unlinkstories", id)))
	return &u, resp, err
}

// LinkBugs 产品计划关联Bug
func (s *ProductPlansService) LinkBugs(id int, bug PlansBugIDs) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&bug).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/linkBugs", id)))
	return &u, resp, err
}

// UnLinkBugs 产品计划取消关联Bug
func (s *ProductPlansService) UnLinkBugs(id int, bug PlansBugIDs) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&bug).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/unlinkbugs", id)))
	return &u, resp, err
}
