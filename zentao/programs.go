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

type ProgramsService struct {
	client *Client
}

type ProgramsMeta struct {
	Name       string `json:"name,omitempty"`
	Parent     int    `json:"parent,omitempty"`
	Desc       string `json:"desc,omitempty"`
	Begin      string `json:"begin,omitempty"`
	End        string `json:"end,omitempty"`
	Pm         any    `json:"PM,omitempty"`         // 请求是字符串, 返回是UserMeta
	Budget     string `json:"budget,omitempty"`     // 预算
	Budgetunit string `json:"budgetUnit,omitempty"` // 预算币种(CNY/USD)
	ACL        ACL    `json:"acl,omitempty"`
	Whitelist  any    `json:"whitelist,omitempty"` // 白名单, 请求和返回([]UserMeta)不一样, 请注意
}

type ProgramsBody struct {
	ID             int       `json:"id"`
	Project        int       `json:"project"`
	Model          string    `json:"model,omitempty"`
	Type           string    `json:"type"`
	Lifetime       string    `json:"lifetime,omitempty"`
	Attribute      string    `json:"attribute,omitempty"`
	Percent        int       `json:"percent"`
	Milestone      string    `json:"milestone"`
	Output         string    `json:"output,omitempty"`
	Auth           string    `json:"auth,omitempty"`
	Path           string    `json:"path"`
	Grade          int       `json:"grade"`
	Code           string    `json:"code"`
	HasProduct     int       `json:"hasProduct"`
	FirstEnd       string    `json:"firstEnd,omitempty"`
	Realbegan      string    `json:"realBegan,omitempty"`
	Realend        string    `json:"realEnd,omitempty"`
	Days           int       `json:"days"`
	Status         string    `json:"status"`
	Substatus      string    `json:"subStatus"`
	Pri            string    `json:"pri"`
	Version        int       `json:"version"`
	Parentversion  int       `json:"parentVersion"`
	Planduration   int       `json:"planDuration"`
	Realduration   int       `json:"realDuration"`
	Openedby       UserMeta  `json:"openedBy"`
	Openeddate     time.Time `json:"openedDate"`
	Openedversion  string    `json:"openedVersion"`
	Lasteditedby   string    `json:"lastEditedBy"`
	Lastediteddate time.Time `json:"lastEditedDate"`
	Closedby       UserMeta  `json:"closedBy"`
	Closeddate     time.Time `json:"closedDate"`
	Canceledby     UserMeta  `json:"canceledBy"`
	Canceleddate   time.Time `json:"canceledDate"`
	Po             UserMeta  `json:"PO,omitempty"`
	Qd             UserMeta  `json:"QD,omitempty"`
	Rd             UserMeta  `json:"RD,omitempty"`
	Team           string    `json:"team,omitempty"`
	Order          int       `json:"order"`
	Deleted        bool      `json:"deleted"`
}

type ProgramsListBody struct {
	ProgramsMeta
	ProgramsBody
	Progress string `json:"progress"`
}

type ProgramsList struct {
	Programs []ProgramsListBody `json:"programs"`
}

type ProgramsMsg struct {
	ProgramsMeta
	ProgramsBody
}

// List 获取项目集列表
func (s *ProgramsService) List(order string) (*ProgramsList, *req.Response, error) {
	var u ProgramsList
	if order == "" {
		order = "order_asc"
	}
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"order": order,
			"page":  "1",
			"limit": "500",
		}).
		SetSuccessResult(&u).
		Get(s.client.RequestURL("/programs"))
	return &u, resp, err
}

// Create 创建项目集
func (s *ProgramsService) Create(program ProgramsMeta) (*ProgramsMsg, *req.Response, error) {
	var u ProgramsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetSuccessResult(&u).
		Post(s.client.RequestURL("/programs"))
	return &u, resp, err
}

// DeleteByID 删除项目集
func (s *ProgramsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}

// UpdateByID 更新项目集
func (s *ProgramsService) UpdateByID(id int, user ProgramsMeta) (*ProgramsMsg, *req.Response, error) {
	var u ProgramsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}

// GetByID 获取项目集详情
func (s *ProgramsService) GetByID(id int) (*ProgramsMsg, *req.Response, error) {
	var u ProgramsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}
