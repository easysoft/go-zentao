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

type ProjectsService struct {
	client *Client
}

type ProjectsCreateMeta struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Begin    string `json:"begin"`            // 计划开始时间
	End      string `json:"end"`              // 计划结束时间
	Products []int  `json:"products"`         // 关联产品
	Model    string `json:"model,omitempty"`  // 项目模型, 默认为scrum
	Parent   int    `json:"parent,omitempty"` // 所属项目集, 默认为0
}

type ProjectsBody struct {
	ID             int         `json:"id"`
	Project        int         `json:"project"`
	Charter        int         `json:"charter"`
	Name           string      `json:"name,omitempty"`
	Code           string      `json:"code,omitempty"`
	Model          string      `json:"model"`
	Type           string      `json:"type"`
	Category       string      `json:"category"`
	Budget         string      `json:"budget"`
	Budgetunit     string      `json:"budgetUnit"`
	Parent         int         `json:"parent"`
	Begin          string      `json:"begin"`
	End            string      `json:"end"`
	Status         string      `json:"status"`
	SubStatus      string      `json:"subStatus"`
	Openedby       any         `json:"openedBy"`
	Openeddate     string      `json:"openedDate"`
	Pm             any         `json:"PM"`
	Po             any         `json:"PO"`
	Qd             any         `json:"QD"`
	Rd             any         `json:"RD"`
	Progress       any         `json:"progress,omitempty"`
	Auth           ProjectAuth `json:"auth"`
	StoryType      string      `json:"storyType"`
	HasProduct     int         `json:"hasProduct"`
	Days           int         `json:"days"`
	Pri            string      `json:"pri"`
	Desc           string      `json:"desc"`
	Estimate       any         `json:"estimate"` // 列表是字符串, 其他是float64
	Left           int         `json:"left"`
	Consumed       int         `json:"consumed"`
	TeamCount      int         `json:"teamCount"`
	LastEditedBy   any         `json:"lastEditedBy"`
	LastEditedDate string      `json:"lastEditedDate"`
	ClosedBy       any         `json:"closedBy"`
	ClosedDate     string      `json:"closedDate"`
	ClosedReason   string      `json:"closedReason"`
	CanceledBy     any         `json:"canceledBy"`
	CanceledDate   string      `json:"canceledDate"`
	SuspendedDate  string      `json:"suspendedDate"`
	Team           string      `json:"team"`
	Acl            ACL         `json:"acl"`
	Whitelist      any         `json:"whitelist"`
	Vision         string      `json:"vision"`
	Deleted        bool        `json:"deleted"`
	TeamMembers    any         `json:"teamMembers"`
	LeftTasks      string      `json:"leftTasks"`
	StatusTitle    string      `json:"statusTitle"`
	Consume        string      `json:"consume"`
	Surplus        string      `json:"surplus"`
	Invested       float64     `json:"invested"`
	StoryCount     int         `json:"storyCount"`
	StoryPoints    string      `json:"storyPoints"`
	ExecutionCount int         `json:"executionCount"`
	From           string      `json:"from"`
	Actions        []string    `json:"actions"`
	Version        int         `json:"version"`
	Parentversion  int         `json:"parentVersion"`
	Order          int         `json:"order"`
	Lifetime       string      `json:"lifetime"`
	Attribute      string      `json:"attribute"`
	Percent        int         `json:"percent"`
	Milestone      string      `json:"milestone"`
	Output         string      `json:"output"`
	Grade          int         `json:"grade"`
	Realbegan      string      `json:"realBegan"`
	Realend        string      `json:"realEnd"`
	Planduration   int         `json:"planDuration"`
	Realduration   int         `json:"realDuration"`
	Openedversion  string      `json:"openedVersion"`
}

type ProjectsListMsg struct {
	Page     int            `json:"page"`
	Total    int            `json:"total"`
	Limit    int            `json:"limit"`
	Projects []ProjectsBody `json:"projects"`
}

type ProjectsUpdateMeta struct {
	Name       string      `json:"name,omitempty"`
	Code       string      `json:"code,omitempty"`
	Parent     int         `json:"parent,omitempty"`
	Budget     int         `json:"budget,omitempty"`
	BudgetUnit string      `json:"budgetUnit,omitempty"`
	PM         int         `json:"PM,omitempty"`
	Days       int         `json:"days,omitempty"`
	Desc       string      `json:"desc,omitempty"`
	ACL        ACL         `json:"acl,omitempty"`
	Auth       ProjectAuth `json:"auth,omitempty"`
	Whitelist  any         `json:"whitelist,omitempty"`
}

// List 获取项目列表
func (s *ProjectsService) List(limit, page string) (*ProjectsListMsg, *req.Response, error) {
	var u ProjectsListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"page":  page,
			"limit": limit,
		}).
		SetSuccessResult(&u).
		Get(s.client.RequestURL("/projects"))
	return &u, resp, err
}

// Create 创建项目
func (s *ProjectsService) Create(program ProjectsCreateMeta) (*ProjectsBody, *req.Response, error) {
	var u ProjectsBody
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetSuccessResult(&u).
		Post(s.client.RequestURL("/projects"))
	return &u, resp, err
}

// DeleteByID 删除项目
func (s *ProjectsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改项目
func (s *ProjectsService) UpdateByID(id int, pum ProjectsUpdateMeta) (*ProjectsBody, *req.Response, error) {
	var u ProjectsBody
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&pum).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}

// GetByID 获取项目详情
func (s *ProjectsService) GetByID(id int) (*ProjectsBody, *req.Response, error) {
	var u ProjectsBody
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}
