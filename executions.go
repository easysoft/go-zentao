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
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Project    int       `json:"project"`
	Code       string    `json:"code"`
	Type       string    `json:"type"`
	Parent     int       `json:"parent"`
	Begin      string    `json:"begin"`
	End        string    `json:"end"`
	Status     string    `json:"status"`
	OpenedBy   string    `json:"openedBy"`
	OpenedDate time.Time `json:"openedDate"`
	Progress   int       `json:"progress"`
}

type ExecutionsCreateMeta struct {
	Project int    `json:"project"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Days    int    `json:"days"`

	Lifetime string `json:"lifetime,omitempty"`
	P0       string `json:"PO,omitempty"`
	PM       string `json:"PM,omitempty"`
	QD       string `json:"QD,omitempty"`
	RD       string `json:"RD,omitempty"`
}

type ExecutionsCreateMsg struct {
	ID             int         `json:"id"`
	Project        int         `json:"project"`
	Model          string      `json:"model"`
	Type           string      `json:"type"`
	Lifetime       string      `json:"lifetime"`
	Budget         string      `json:"budget"`
	Budgetunit     string      `json:"budgetUnit"`
	Attribute      string      `json:"attribute"`
	Percent        int         `json:"percent"`
	Milestone      string      `json:"milestone"`
	Output         string      `json:"output"`
	Auth           string      `json:"auth"`
	Parent         int         `json:"parent"`
	Path           string      `json:"path"`
	Grade          int         `json:"grade"`
	Name           string      `json:"name"`
	Code           string      `json:"code"`
	Begin          string      `json:"begin"`
	End            string      `json:"end"`
	Realbegan      string      `json:"realBegan"`
	Realend        string      `json:"realEnd"`
	Days           int         `json:"days"`
	Status         string      `json:"status"`
	Substatus      string      `json:"subStatus"`
	Pri            string      `json:"pri"`
	Desc           string      `json:"desc"`
	Version        int         `json:"version"`
	Parentversion  int         `json:"parentVersion"`
	Planduration   int         `json:"planDuration"`
	Realduration   int         `json:"realDuration"`
	Openedby       string      `json:"openedBy"`
	Openeddate     time.Time   `json:"openedDate"`
	Openedversion  string      `json:"openedVersion"`
	Lasteditedby   string      `json:"lastEditedBy"`
	Lastediteddate time.Time   `json:"lastEditedDate"`
	Closedby       string      `json:"closedBy"`
	Closeddate     interface{} `json:"closedDate"`
	Canceledby     string      `json:"canceledBy"`
	Canceleddate   interface{} `json:"canceledDate"`
	Po             string      `json:"PO"`
	Pm             string      `json:"PM"`
	Qd             string      `json:"QD"`
	Rd             string      `json:"RD"`
	Team           string      `json:"team"`
	ACL            string      `json:"acl"`
	Whitelist      string      `json:"whitelist"`
	Order          int         `json:"order"`
	Deleted        string      `json:"deleted"`
	Totalhours     int         `json:"totalHours"`
	Totalestimate  int         `json:"totalEstimate"`
	Totalconsumed  int         `json:"totalConsumed"`
	Totalleft      int         `json:"totalLeft"`
}

type ExecutionsUpdateMeta struct {
	ExecutionsCreateMeta
	TeamMembers []string `json:"teamMembers,omitempty"`
	Desc        string   `json:"desc,omitempty"`
	ACL         string   `json:"acl,omitempty"`
	Whitelist   []string `json:"whitelist,omitempty"`
}

func (s *ExecutionsService) ListByProject(id int64) (*ProjectExecutions, *req.Response, error) {
	var pe ProjectExecutions
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&pe).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/executions", id)))
	return &pe, resp, err
}

// Create 创建执行
func (s *ExecutionsService) Create(id int, build ExecutionsCreateMeta) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/projects/%d/executions", id)))
	return &u, resp, err
}

// DeleteByID 删除执行
func (s *ExecutionsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改执行
func (s *ExecutionsService) UpdateByID(id int, exec ExecutionsUpdateMeta) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}

// GetByID 获取执行详情
func (s *ExecutionsService) GetByID(id int) (*ExecutionsCreateMsg, *req.Response, error) {
	var u ExecutionsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d", id)))
	return &u, resp, err
}
