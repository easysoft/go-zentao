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

type BuildsService struct {
	client *Client
}

type BuildsListMsg struct {
	Total  int          `json:"total"`
	Builds []BuildsBody `json:"builds"`
}

type BuildsCreateMeta struct {
	Execution int    `json:"execution"`
	Product   int    `json:"product"`
	Name      string `json:"name"`
	Builder   string `json:"builder"`
	Date      string `json:"date"`

	Branch   int    `json:"branch,omitempty"`
	ScmPath  string `json:"scmPath,omitempty"`
	FilePath string `json:"filePath,omitempty"`
	Desc     string `json:"desc,omitempty"`
}

type BuildsBody struct {
	ID            int           `json:"id"`
	Project       int           `json:"project"`
	Product       int           `json:"product"`
	Branch        int           `json:"branch"`
	Execution     int           `json:"execution"`
	Name          string        `json:"name"`
	Scmpath       string        `json:"scmPath"`
	Filepath      string        `json:"filePath"`
	Date          string        `json:"date"`
	Stories       []StoriesBody `json:"stories"`
	Bugs          []BugBody     `json:"bugs"`
	Builder       string        `json:"builder"`
	Desc          string        `json:"desc"`
	Deleted       string        `json:"deleted"`
	Executionname string        `json:"executionName"`
	Productname   string        `json:"productName"`
	Executionid   int           `json:"executionID,omitempty"`
	Branchname    string        `json:"branchName,omitempty"`
}

type BuildsCreateMsg struct {
	BuildsBody
	Producttype string        `json:"productType,omitempty"`
	Files       []interface{} `json:"files,omitempty"`
}

// ProjectsList 获取项目版本列表
func (s *BuildsService) ProjectsList(id int) (*BuildsListMsg, *req.Response, error) {
	var u BuildsListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/builds", id)))
	return &u, resp, err
}

// ExecutionsList 获取执行版本列表
func (s *BuildsService) ExecutionsList(id int) (*BuildsListMsg, *req.Response, error) {
	var u BuildsListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/builds", id)))
	return &u, resp, err
}

// Create 创建版本
func (s *BuildsService) Create(id int, build BuildsCreateMeta) (*BuildsCreateMsg, *req.Response, error) {
	var u BuildsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/projects/%d/builds", id)))
	return &u, resp, err
}

// DeleteByID 删除版本
func (s *BuildsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/builds/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改版本
func (s *BuildsService) UpdateByID(id int, build BuildsCreateMeta) (*BuildsCreateMsg, *req.Response, error) {
	var u BuildsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/builds/%d", id)))
	return &u, resp, err
}

// GetByID 获取版本详情
func (s *BuildsService) GetByID(id int) (*BuildsCreateMsg, *req.Response, error) {
	var u BuildsCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/builds/%d", id)))
	return &u, resp, err
}
