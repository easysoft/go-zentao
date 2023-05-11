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
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Model      string    `json:"model"`
	Type       string    `json:"type"`
	Budget     string    `json:"budget"`
	Budgetunit string    `json:"budgetUnit"`
	Parent     int       `json:"parent"`
	Begin      string    `json:"begin"`
	End        string    `json:"end"`
	Status     string    `json:"status"`
	Openedby   UserMeta  `json:"openedBy"`
	Openeddate time.Time `json:"openedDate"`
	Pm         UserMeta  `json:"PM"`
	Progress   int       `json:"progress,omitempty"`
}

type ProjectsCreateMsg struct {
	ProjectsBody
	Lifetime       string      `json:"lifetime"`
	Attribute      string      `json:"attribute"`
	Percent        int         `json:"percent"`
	Milestone      string      `json:"milestone"`
	Output         string      `json:"output"`
	Auth           string      `json:"auth"`
	Path           string      `json:"path"`
	Grade          int         `json:"grade"`
	Realbegan      string      `json:"realBegan"`
	Realend        string      `json:"realEnd"`
	Days           int         `json:"days"`
	Substatus      string      `json:"subStatus"`
	Pri            string      `json:"pri"`
	Desc           string      `json:"desc"`
	Version        int         `json:"version"`
	Parentversion  int         `json:"parentVersion"`
	Planduration   int         `json:"planDuration"`
	Realduration   int         `json:"realDuration"`
	Openedversion  string      `json:"openedVersion"`
	Lasteditedby   UserMeta    `json:"lastEditedBy"`
	Lastediteddate time.Time   `json:"lastEditedDate"`
	Closedby       UserMeta    `json:"closedBy"`
	Closeddate     time.Time   `json:"closedDate"`
	Canceledby     UserMeta    `json:"canceledBy"`
	Canceleddate   time.Time   `json:"canceledDate"`
	Po             interface{} `json:"PO"`
	Qd             interface{} `json:"QD"`
	Rd             interface{} `json:"RD"`
	Team           string      `json:"team"`
	ACL            string      `json:"acl"`
	Whitelist      []UserMeta  `json:"whitelist"`
	Order          int         `json:"order"`
	Deleted        bool        `json:"deleted"`
}

type ProjectsListMsg struct {
	Page     int            `json:"page"`
	Total    int            `json:"total"`
	Limit    int            `json:"limit"`
	Projects []ProjectsBody `json:"projects"`
}

type ProjectsUpdateMeta struct {
	Name      string   `json:"name,omitempty"`
	Code      string   `json:"code,omitempty"`
	Parent    int      `json:"parent,omitempty"`
	Days      int      `json:"days,omitempty"`
	Desc      string   `json:"desc,omitempty"`
	ACL       string   `json:"acl,omitempty"`
	Auth      string   `json:"auth,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
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
		SetResult(&u).
		Get(s.client.RequestURL("/projects"))
	return &u, resp, err
}

// Create 创建项目
func (s *ProjectsService) Create(program ProjectsCreateMeta) (*ProjectsCreateMsg, *req.Response, error) {
	var u ProjectsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetResult(&u).
		Post(s.client.RequestURL("/projects"))
	return &u, resp, err
}

// DeleteByID 删除项目
func (s *ProjectsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改项目
func (s *ProjectsService) UpdateByID(id int, pum ProjectsUpdateMeta) (*ProjectsCreateMsg, *req.Response, error) {
	var u ProjectsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&pum).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}

// GetByID 获取项目详情
func (s *ProjectsService) GetByID(id int) (*ProjectsCreateMsg, *req.Response, error) {
	var u ProjectsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d", id)))
	return &u, resp, err
}
