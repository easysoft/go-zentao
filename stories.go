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

type StoriesService struct {
	client *Client
}

type StoriesListMeta struct {
	Page    int `json:"page"`
	Total   int `json:"total"`
	Limit   int `json:"limit"`
	Stories []struct {
		ID             int         `json:"id"`
		Parent         int         `json:"parent"`
		Product        int         `json:"product"`
		Branch         int         `json:"branch"`
		Module         int         `json:"module"`
		Plan           string      `json:"plan"`
		Source         string      `json:"source"`
		Sourcenote     string      `json:"sourceNote"`
		Frombug        int         `json:"fromBug"`
		Title          string      `json:"title"`
		Keywords       string      `json:"keywords"`
		Type           string      `json:"type"`
		Category       string      `json:"category"`
		Pri            int         `json:"pri"`
		Estimate       int         `json:"estimate"`
		Status         string      `json:"status"`
		Substatus      string      `json:"subStatus"`
		Color          string      `json:"color"`
		Stage          string      `json:"stage"`
		Stagedby       string      `json:"stagedBy"`
		Mailto         string      `json:"mailto"`
		Openedby       string      `json:"openedBy"`
		Openeddate     time.Time   `json:"openedDate"`
		Assignedto     string      `json:"assignedTo"`
		Assigneddate   time.Time   `json:"assignedDate"`
		Lasteditedby   string      `json:"lastEditedBy"`
		Lastediteddate interface{} `json:"lastEditedDate"`
		Reviewedby     string      `json:"reviewedBy"`
		Revieweddate   interface{} `json:"reviewedDate"`
		Closedby       string      `json:"closedBy"`
		Closeddate     interface{} `json:"closedDate"`
		Closedreason   string      `json:"closedReason"`
		Tobug          int         `json:"toBug"`
		Childstories   string      `json:"childStories"`
		Linkstories    string      `json:"linkStories"`
		Duplicatestory int         `json:"duplicateStory"`
		Version        int         `json:"version"`
		Urchanged      string      `json:"URChanged"`
		Deleted        bool        `json:"deleted"`
		Plantitle      string      `json:"planTitle"`
	} `json:"stories"`
}

type StoriesCreateMeta struct {
	Title      string  `json:"title"`
	Product    int     `json:"product"`
	Pri        int     `json:"pri"`
	Category   string  `json:"category"`
	Spec       string  `json:"spec,omitempty"`
	Verify     string  `json:"verify,omitempty"`
	Source     string  `json:"source,omitempty"`
	SourceNote string  `json:"sourceNote,omitempty"`
	Estimate   float64 `json:"estimate,omitempty"`
	Keywords   string  `json:"keywords,omitempty"`
}

type StoriesCreateMsg struct {
	ID             int           `json:"id"`
	Parent         int           `json:"parent"`
	Product        int           `json:"product"`
	Branch         int           `json:"branch"`
	Module         int           `json:"module"`
	Plan           string        `json:"plan"`
	Source         string        `json:"source"`
	Sourcenote     string        `json:"sourceNote"`
	Frombug        int           `json:"fromBug"`
	Title          string        `json:"title"`
	Keywords       string        `json:"keywords"`
	Type           string        `json:"type"`
	Category       string        `json:"category"`
	Pri            int           `json:"pri"`
	Estimate       int           `json:"estimate"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Color          string        `json:"color"`
	Stage          string        `json:"stage"`
	Stagedby       string        `json:"stagedBy"`
	Mailto         interface{}   `json:"mailto"`
	Openedby       string        `json:"openedBy"`
	Openeddate     time.Time     `json:"openedDate"`
	Assignedto     string        `json:"assignedTo"`
	Assigneddate   interface{}   `json:"assignedDate"`
	Lasteditedby   string        `json:"lastEditedBy"`
	Lastediteddate interface{}   `json:"lastEditedDate"`
	Reviewedby     string        `json:"reviewedBy"`
	Revieweddate   interface{}   `json:"reviewedDate"`
	Closedby       string        `json:"closedBy"`
	Closeddate     interface{}   `json:"closedDate"`
	Closedreason   string        `json:"closedReason"`
	Tobug          int           `json:"toBug"`
	Childstories   string        `json:"childStories"`
	Linkstories    string        `json:"linkStories"`
	Duplicatestory int           `json:"duplicateStory"`
	Version        int           `json:"version"`
	Urchanged      string        `json:"URChanged"`
	Deleted        bool          `json:"deleted"`
	Spec           string        `json:"spec"`
	Verify         string        `json:"verify"`
	Executions     []interface{} `json:"executions"`
	Tasks          []interface{} `json:"tasks"`
	Stages         []interface{} `json:"stages"`
	Children       []interface{} `json:"children"`
}

type StoriesUpdateMeta struct {
	Title  string `json:"title,omitempty"`
	Spec   string `json:"spec,omitempty"`
	Verify string `json:"verify,omitempty"`
}

type StoriesUpdateFieldMeta struct {
	Module     int     `json:"module,omitempty"`
	Source     string  `json:"source,omitempty"`
	SourceNote string  `json:"sourceNote,omitempty"`
	Pri        int     `json:"pri,omitempty"`
	Category   string  `json:"category,omitempty"`
	Estimate   float64 `json:"estimate,omitempty"`
	Keywords   string  `json:"keywords,omitempty"`
}

type StoriesUpdateMsg struct {
	ID             int         `json:"id"`
	Parent         int         `json:"parent"`
	Product        int         `json:"product"`
	Branch         int         `json:"branch"`
	Module         int         `json:"module"`
	Plan           string      `json:"plan"`
	Source         string      `json:"source"`
	Sourcenote     string      `json:"sourceNote"`
	Frombug        int         `json:"fromBug"`
	Title          string      `json:"title"`
	Keywords       string      `json:"keywords"`
	Type           string      `json:"type"`
	Category       string      `json:"category"`
	Pri            int         `json:"pri"`
	Estimate       int         `json:"estimate"`
	Status         string      `json:"status"`
	Substatus      string      `json:"subStatus"`
	Color          string      `json:"color"`
	Stage          string      `json:"stage"`
	Stagedby       string      `json:"stagedBy"`
	Mailto         string      `json:"mailto"`
	Openedby       string      `json:"openedBy"`
	Openeddate     time.Time   `json:"openedDate"`
	Assignedto     string      `json:"assignedTo"`
	Assigneddate   interface{} `json:"assignedDate"`
	Lasteditedby   string      `json:"lastEditedBy"`
	Lastediteddate time.Time   `json:"lastEditedDate"`
	Reviewedby     string      `json:"reviewedBy"`
	Revieweddate   interface{} `json:"reviewedDate"`
	Closedby       string      `json:"closedBy"`
	Closeddate     interface{} `json:"closedDate"`
	Closedreason   string      `json:"closedReason"`
	Tobug          int         `json:"toBug"`
	Childstories   string      `json:"childStories"`
	Linkstories    string      `json:"linkStories"`
	Duplicatestory int         `json:"duplicateStory"`
	Version        int         `json:"version"`
	Urchanged      string      `json:"URChanged"`
	Deleted        string      `json:"deleted"`
	Spec           string      `json:"spec"`
	Verify         string      `json:"verify"`
	Executions     struct {
		Num1 struct {
			Project int    `json:"project"`
			Name    string `json:"name"`
			Status  string `json:"status"`
		} `json:"1"`
	} `json:"executions"`
	Tasks struct {
		Num1 []struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Assignedto string `json:"assignedTo"`
			Execution  int    `json:"execution"`
			Status     string `json:"status"`
			Consumed   int    `json:"consumed"`
			Left       int    `json:"left"`
			Type       string `json:"type"`
		} `json:"1"`
	} `json:"tasks"`
	Stages    []interface{} `json:"stages"`
	Plantitle struct {
		Num1 string `json:"1"`
	} `json:"planTitle"`
	Children []interface{} `json:"children"`
}

// ProjectsList 获取项目需求列表
func (s *StoriesService) ProjectsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/stories", id)))
	return &u, resp, err
}

// ProductsList 获取产品需求列表
func (s *StoriesService) ProductsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/stories", id)))
	return &u, resp, err
}

// Create 创建需求
func (s *StoriesService) Create(story StoriesCreateMeta) (*StoriesCreateMsg, *req.Response, error) {
	var u StoriesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL("/stories"))
	return &u, resp, err
}

// DeleteByID 删除需求
func (s *StoriesService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}

// UpdateByID 变更需求
func (s *StoriesService) UpdateByID(id int, story StoriesUpdateMeta) (*StoriesUpdateMsg, *req.Response, error) {
	var u StoriesUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/change", id)))
	return &u, resp, err
}

// GetByID 获取需求详情
func (s *StoriesService) GetByID(id int) (*StoriesCreateMsg, *req.Response, error) {
	var u StoriesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}

// ExecutionsList 获取执行需求列表
func (s *StoriesService) ExecutionsList(id int) (*StoriesListMeta, *req.Response, error) {
	var u StoriesListMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/stories", id)))
	return &u, resp, err
}

// UpdateFieldByID 修改需求其他字段
func (s *StoriesService) UpdateFieldByID(id int, uf StoriesUpdateFieldMeta) (*StoriesCreateMsg, *req.Response, error) {
	var u StoriesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&uf).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/stories/%d", id)))
	return &u, resp, err
}
