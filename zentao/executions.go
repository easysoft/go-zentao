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

	"github.com/davecgh/go-spew/spew"
	"github.com/imroc/req/v3"
)

type ExecutionsService struct {
	client *Client
}

type ProjectExecutions struct {
	Page       int         `json:"page"`
	Total      int         `json:"total"`
	Limit      int         `json:"limit"`
	Executions []Execution `json:"executions"`
}

type Execution struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Project    int    `json:"project"`
	Code       string `json:"code"`
	Type       string `json:"type"`
	Parent     int    `json:"parent"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Status     string `json:"status"`
	OpenedBy   any    `json:"openedBy"`
	OpenedDate string `json:"openedDate"`
	Delay      int    `json:"delay"`
	Progress   any    `json:"progress"`
}

type ExecutionsCreateMeta struct {
	Project int    `json:"project"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Begin   string `json:"begin"`
	End     string `json:"end"`

	Plans    []int `json:"plans,omitempty"`
	Products []int `json:"products,omitempty"`

	Days        int               `json:"days,omitempty"`
	Lifetime    ExecutionLifeTime `json:"lifetime,omitempty"`
	P0          string            `json:"PO,omitempty"`
	PM          string            `json:"PM,omitempty"`
	QD          string            `json:"QD,omitempty"`
	RD          string            `json:"RD,omitempty"`
	TeamMembers []string          `json:"teamMembers,omitempty"`
	Desc        string            `json:"desc,omitempty"`
	Acl         ACL               `json:"acl,omitempty"`
	// Whitelist   []string `json:"whitelist,omitempty"`
}

type ExecutionsCreateMsg struct {
	ID             int               `json:"id"`
	Project        int               `json:"project"`
	Model          string            `json:"model"`
	Type           string            `json:"type"`
	Lifetime       ExecutionLifeTime `json:"lifetime"`
	Budget         string            `json:"budget"`
	Budgetunit     string            `json:"budgetUnit"`
	Attribute      string            `json:"attribute"`
	Percent        int               `json:"percent"`
	Milestone      string            `json:"milestone"`
	Output         string            `json:"output"`
	Auth           string            `json:"auth"`
	Parent         int               `json:"parent"`
	Path           string            `json:"path"`
	Grade          int               `json:"grade"`
	Name           string            `json:"name"`
	Code           string            `json:"code"`
	Begin          string            `json:"begin"`
	End            string            `json:"end"`
	Realbegan      string            `json:"realBegan"`
	Realend        string            `json:"realEnd"`
	Days           int               `json:"days"`
	Status         string            `json:"status"`
	Substatus      string            `json:"subStatus"`
	Pri            string            `json:"pri"`
	Desc           string            `json:"desc"`
	Version        int               `json:"version"`
	Parentversion  int               `json:"parentVersion"`
	Planduration   int               `json:"planDuration"`
	Realduration   int               `json:"realDuration"`
	Openedby       any               `json:"openedBy"`
	Openeddate     string            `json:"openedDate"`
	Openedversion  string            `json:"openedVersion"`
	Lasteditedby   any               `json:"lastEditedBy"`
	Lastediteddate string            `json:"lastEditedDate"`
	Closedby       any               `json:"closedBy"`
	Closeddate     string            `json:"closedDate"`
	Canceledby     any               `json:"canceledBy"`
	Canceleddate   string            `json:"canceledDate"`
	Po             any               `json:"PO"`
	Pm             any               `json:"PM"`
	Qd             any               `json:"QD"`
	Rd             any               `json:"RD"`
	Team           string            `json:"team"`
	ACL            ACL               `json:"acl"`
	Whitelist      any               `json:"whitelist"`
	Order          int               `json:"order"`
	Deleted        bool              `json:"deleted"`
	Totalhours     int               `json:"totalHours"`
	Totalestimate  int               `json:"totalEstimate"`
	Totalconsumed  int               `json:"totalConsumed"`
	Totalleft      int               `json:"totalLeft"`
}

type ExecutionsUpdateMeta struct {
	ExecutionsCreateMeta
	TeamMembers []string `json:"teamMembers,omitempty"`
	Desc        string   `json:"desc,omitempty"`
	ACL         ACL      `json:"acl,omitempty"`
	Whitelist   []string `json:"whitelist,omitempty"`
}

func (s *ExecutionsService) ListByProject(id int64) (*ProjectExecutions, *req.Response, error) {
	var pe ProjectExecutions
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&pe).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/executions", id)))
	return &pe, resp, err
}

// Create 创建执行
func (s *ExecutionsService) Create(id int, build ExecutionsCreateMeta) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	spew.Dump(build)
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/projects/%d/executions", id)))
	return &u, resp, err
}

// DeleteByID 删除执行
func (s *ExecutionsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改执行
func (s *ExecutionsService) UpdateByID(id int, exec ExecutionsUpdateMeta) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}

// GetByID 获取执行详情
func (s *ExecutionsService) GetByID(id int) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}
