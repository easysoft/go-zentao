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

type ProgramsService struct {
	client *Client
}

type ProgramsList struct {
	Programs []struct {
		ID            int    `json:"id"`
		Project       int    `json:"project"`
		Model         string `json:"model"`
		Type          string `json:"type"`
		Lifetime      string `json:"lifetime"`
		Budget        string `json:"budget"`
		Budgetunit    string `json:"budgetUnit"`
		Attribute     string `json:"attribute"`
		Percent       int    `json:"percent"`
		Milestone     string `json:"milestone"`
		Output        string `json:"output"`
		Auth          string `json:"auth"`
		Parent        int    `json:"parent"`
		Path          string `json:"path"`
		Grade         int    `json:"grade"`
		Name          string `json:"name"`
		Code          string `json:"code"`
		Begin         string `json:"begin"`
		End           string `json:"end"`
		Realbegan     string `json:"realBegan"`
		Realend       string `json:"realEnd"`
		Days          int    `json:"days"`
		Status        string `json:"status"`
		Substatus     string `json:"subStatus"`
		Pri           string `json:"pri"`
		Desc          string `json:"desc"`
		Version       int    `json:"version"`
		Parentversion int    `json:"parentVersion"`
		Planduration  int    `json:"planDuration"`
		Realduration  int    `json:"realDuration"`
		Openedby      struct {
			ID       int    `json:"id"`
			Account  string `json:"account"`
			Avatar   string `json:"avatar"`
			Realname string `json:"realname"`
		} `json:"openedBy"`
		Openeddate     time.Time   `json:"openedDate"`
		Openedversion  string      `json:"openedVersion"`
		Lasteditedby   string      `json:"lastEditedBy"`
		Lastediteddate time.Time   `json:"lastEditedDate"`
		Closedby       interface{} `json:"closedBy"`
		Closeddate     time.Time   `json:"closedDate"`
		Canceledby     interface{} `json:"canceledBy"`
		Canceleddate   time.Time   `json:"canceledDate"`
		Po             interface{} `json:"PO"`
		Pm             struct {
			ID       int    `json:"id"`
			Account  string `json:"account"`
			Avatar   string `json:"avatar"`
			Realname string `json:"realname"`
		} `json:"PM"`
		Qd        interface{} `json:"QD"`
		Rd        interface{} `json:"RD"`
		Team      string      `json:"team"`
		ACL       string      `json:"acl"`
		Whitelist interface{} `json:"whitelist"`
		Order     int         `json:"order"`
		Deleted   bool        `json:"deleted"`
		Progress  int         `json:"progress"`
	} `json:"programs"`
}

type ProgramsCreateMsg struct {
	ProgramsMeta
	ID            int         `json:"id"`
	Project       int         `json:"project"`
	Model         string      `json:"model"`
	Type          string      `json:"type"`
	Lifetime      string      `json:"lifetime"`
	Budget        string      `json:"budget"`
	Budgetunit    string      `json:"budgetUnit"`
	Attribute     string      `json:"attribute"`
	Percent       int         `json:"percent"`
	Milestone     string      `json:"milestone"`
	Output        string      `json:"output"`
	Auth          string      `json:"auth"`
	Path          string      `json:"path"`
	Grade         int         `json:"grade"`
	Code          string      `json:"code"`
	Realbegan     interface{} `json:"realBegan"`
	Realend       interface{} `json:"realEnd"`
	Days          int         `json:"days"`
	Status        string      `json:"status"`
	Substatus     string      `json:"subStatus"`
	Pri           string      `json:"pri"`
	Version       int         `json:"version"`
	Parentversion int         `json:"parentVersion"`
	Planduration  int         `json:"planDuration"`
	Realduration  int         `json:"realDuration"`
	Openedby      struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"openedBy"`
	Openeddate     time.Time   `json:"openedDate"`
	Openedversion  string      `json:"openedVersion"`
	Lasteditedby   string      `json:"lastEditedBy"`
	Lastediteddate interface{} `json:"lastEditedDate"`
	Closedby       interface{} `json:"closedBy"`
	Closeddate     interface{} `json:"closedDate"`
	Canceledby     interface{} `json:"canceledBy"`
	Canceleddate   interface{} `json:"canceledDate"`
	Po             interface{} `json:"PO"`
	Pm             interface{} `json:"PM"`
	Qd             interface{} `json:"QD"`
	Rd             interface{} `json:"RD"`
	Team           string      `json:"team"`
	ACL            string      `json:"acl"`
	Whitelist      []struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"whitelist"`
	Order   int  `json:"order"`
	Deleted bool `json:"deleted"`
}

type ProgramsMeta struct {
	Name   string `json:"name"`
	Parent int    `json:"parent"`
	Desc   string `json:"desc"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
}

type ProgramsUpdateMsg struct {
	ProgramsMeta
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
	Path           string      `json:"path"`
	Grade          int         `json:"grade"`
	Code           string      `json:"code"`
	Realbegan      string      `json:"realBegan"`
	Realend        string      `json:"realEnd"`
	Days           int         `json:"days"`
	Status         string      `json:"status"`
	Substatus      string      `json:"subStatus"`
	Pri            string      `json:"pri"`
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
}

type ProgramsGetMsg struct {
	ProgramsMeta
	ID            int         `json:"id"`
	Project       int         `json:"project"`
	Model         string      `json:"model"`
	Type          string      `json:"type"`
	Lifetime      string      `json:"lifetime"`
	Budget        string      `json:"budget"`
	Budgetunit    string      `json:"budgetUnit"`
	Attribute     string      `json:"attribute"`
	Milestone     string      `json:"milestone"`
	Output        string      `json:"output"`
	Auth          string      `json:"auth"`
	Parent        int         `json:"parent"`
	Path          string      `json:"path"`
	Grade         int         `json:"grade"`
	Code          string      `json:"code"`
	Realbegan     interface{} `json:"realBegan"`
	Realend       interface{} `json:"realEnd"`
	Days          int         `json:"days"`
	Status        string      `json:"status"`
	Substatus     string      `json:"subStatus"`
	Pri           string      `json:"pri"`
	Version       int         `json:"version"`
	Parentversion int         `json:"parentVersion"`
	Planduration  int         `json:"planDuration"`
	Realduration  int         `json:"realDuration"`
	Openedby      struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"openedBy"`
	Openeddate     time.Time   `json:"openedDate"`
	Openedversion  string      `json:"openedVersion"`
	Lasteditedby   string      `json:"lastEditedBy"`
	Lastediteddate interface{} `json:"lastEditedDate"`
	Closedby       interface{} `json:"closedBy"`
	Closeddate     interface{} `json:"closedDate"`
	Canceledby     interface{} `json:"canceledBy"`
	Canceleddate   interface{} `json:"canceledDate"`
	Po             interface{} `json:"PO"`
	Pm             struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"PM"`
	Qd        interface{}   `json:"QD"`
	Rd        interface{}   `json:"RD"`
	Team      string        `json:"team"`
	ACL       string        `json:"acl"`
	Whitelist []interface{} `json:"whitelist"`
	Order     int           `json:"order"`
	Deleted   bool          `json:"deleted"`
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
		}).
		SetResult(&u).
		Get(s.client.RequestURL("/programs"))
	return &u, resp, err
}

// Create 创建项目集
func (s *ProgramsService) Create(program ProgramsMeta) (*ProgramsCreateMsg, *req.Response, error) {
	var u ProgramsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&program).
		SetResult(&u).
		Post(s.client.RequestURL("/programs"))
	return &u, resp, err
}

// DeleteByID 删除项目集
func (s *ProgramsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}

// UpdateByID 更新项目集
func (s *ProgramsService) UpdateByID(id int, user ProgramsMeta) (*ProgramsUpdateMsg, *req.Response, error) {
	var u ProgramsUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}

// GetByID 获取项目集详情
func (s *ProgramsService) GetByID(id int) (*ProgramsGetMsg, *req.Response, error) {
	var u ProgramsGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/programs/%d", id)))
	return &u, resp, err
}
