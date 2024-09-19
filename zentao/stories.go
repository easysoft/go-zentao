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

type StoriesService struct {
	client *Client
}

type StoriesListMeta struct {
	Page    int           `json:"page"`
	Total   int           `json:"total"`
	Limit   int           `json:"limit"`
	Stories []StoriesBody `json:"stories"`
}

type StoriesMeta struct {
	Title  string `json:"title,omitempty"`
	Spec   string `json:"spec,omitempty"`
	Verify string `json:"verify,omitempty"`
}

type StoriesExtMeta struct {
	Source     string  `json:"source,omitempty"`
	SourceNote string  `json:"sourceNote,omitempty"`
	Pri        int     `json:"pri,omitempty"`
	Category   string  `json:"category,omitempty"`
	Estimate   float64 `json:"estimate,omitempty"`
	Keywords   string  `json:"keywords,omitempty"`
}

type StoriesBody struct {
	StoriesExtMeta
	ID             int    `json:"id"`
	Parent         any    `json:"parent"` // 可能是数组, 产品计划时是数组
	Product        int    `json:"product"`
	Branch         int    `json:"branch"`
	Module         int    `json:"module"`
	Plan           string `json:"plan"`
	Frombug        int    `json:"fromBug"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	Substatus      string `json:"subStatus"`
	Color          string `json:"color"`
	Stage          string `json:"stage"`
	Stagedby       string `json:"stagedBy"`
	Mailto         string `json:"mailto"`
	Openedby       any    `json:"openedBy"` // 产品计划是string, 其他可能是UserMeta
	Openeddate     string `json:"openedDate"`
	Assignedto     string `json:"assignedTo"`
	Assigneddate   any    `json:"assignedDate"`
	Lasteditedby   string `json:"lastEditedBy"`
	Lastediteddate string `json:"lastEditedDate"`
	Reviewedby     string `json:"reviewedBy"`
	Revieweddate   any    `json:"reviewedDate"`
	Closedby       string `json:"closedBy"`
	Closeddate     any    `json:"closedDate"`
	Closedreason   string `json:"closedReason"`
	Tobug          int    `json:"toBug"`
	Childstories   string `json:"childStories"`
	Linkstories    string `json:"linkStories"`
	Duplicatestory int    `json:"duplicateStory"`
	Version        int    `json:"version"`
	Urchanged      string `json:"URChanged"`
	Deleted        string `json:"deleted"`
	Plantitle      string `json:"planTitle,omitempty"`
}

type StoriesCreateMeta struct {
	StoriesMeta
	StoriesExtMeta
	Product int `json:"product"`
}

type StoriesUpdateFieldMeta struct {
	StoriesExtMeta
	Module int `json:"module,omitempty"`
}

type StoriesMsg struct {
	StoriesBody
	Spec       string        `json:"spec"`
	Verify     string        `json:"verify"`
	Executions []interface{} `json:"executions"`
	Tasks      []interface{} `json:"tasks"`
	Stages     []interface{} `json:"stages"`
	Children   []interface{} `json:"children"`
}

// ProjectsList 获取项目需求列表
func (s *StoriesService) ProjectsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/stories", id)))
	return &u, resp, err
}

// ProductsList 获取产品需求列表
func (s *StoriesService) ProductsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/stories", id)))
	return &u, resp, err
}

// Create 创建需求
func (s *StoriesService) Create(story StoriesCreateMeta) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL("/stories"))
	return &u, resp, err
}

// DeleteByID 删除需求
func (s *StoriesService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}

// UpdateByID 变更需求
func (s *StoriesService) UpdateByID(id int, story StoriesMeta) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/change", id)))
	return &u, resp, err
}

// GetByID 获取需求详情
func (s *StoriesService) GetByID(id int) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}

// ExecutionsList 获取执行需求列表
func (s *StoriesService) ExecutionsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/stories", id)))
	return &u, resp, err
}

// UpdateFieldByID 修改需求其他字段
func (s *StoriesService) UpdateFieldByID(id int, uf StoriesUpdateFieldMeta) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&uf).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}
