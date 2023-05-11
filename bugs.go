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

type BugsService struct {
	client *Client
}

type ListProductsBugsMsg struct {
	Page  int       `json:"page"`
	Total int       `json:"total"`
	Limit int       `json:"limit"`
	Bugs  []BugBody `json:"bugs"`
}

type BugMeta struct {
	Branch    int    `json:"branch,omitempty"`
	Module    int    `json:"module,omitempty"`
	Execution int    `json:"execution,omitempty"`
	Os        string `json:"os,omitempty"`
	Browser   string `json:"browser,omitempty"`
	Task      int    `json:"task,omitempty"`
	Story     int    `json:"story,omitempty"`
	Deadline  string `json:"deadline,omitempty"`
	Title     string `json:"title"`
	Severity  int    `json:"severity"`
	Pri       int    `json:"pri"`
	Steps     string `json:"steps"`
	Type      string `json:"type"`
}

type BugBody struct {
	BugMeta
	ID             int           `json:"id"`
	Project        int           `json:"project"`
	Product        int           `json:"product"`
	Plan           int           `json:"plan"`
	Storyversion   int           `json:"storyVersion"`
	Totask         int           `json:"toTask"`
	Tostory        int           `json:"toStory"`
	Keywords       string        `json:"keywords"`
	Hardware       string        `json:"hardware"`
	Found          string        `json:"found"`
	Status         interface{}   `json:"status,omitempty"` // 列表返回结构体, 详情返回字符串
	Substatus      string        `json:"subStatus"`
	Color          string        `json:"color"`
	Confirmed      int           `json:"confirmed"`
	Activatedcount int           `json:"activatedCount"`
	Entry          string        `json:"entry"`
	Lines          string        `json:"lines"`
	V1             string        `json:"v1"`
	V2             string        `json:"v2"`
	Duplicatebug   int           `json:"duplicateBug"`
	Linkbug        string        `json:"linkBug"`
	Case           int           `json:"case"`
	Caseversion    int           `json:"caseVersion"`
	Result         int           `json:"result"`
	Repo           int           `json:"repo"`
	Repotype       string        `json:"repoType"`
	Testtask       int           `json:"testtask"`
	Deleted        bool          `json:"deleted"`
	Activateddate  time.Time     `json:"activatedDate"`
	Feedbackby     string        `json:"feedbackBy,omitempty"`  // 仅bug接口
	Notifyemail    string        `json:"notifyEmail,omitempty"` // 仅bug接口
	Mailto         []interface{} `json:"mailto"`
	Openedby       UserMeta      `json:"openedBy"`
	Openeddate     time.Time     `json:"openedDate"`
	Openedbuild    string        `json:"openedBuild"`
	Assignedto     UserMeta      `json:"assignedTo"`
	Assigneddate   time.Time     `json:"assignedDate"`
	Resolvedby     interface{}   `json:"resolvedBy"`
	Resolution     string        `json:"resolution"`
	Resolvedbuild  string        `json:"resolvedBuild"`
	Resolveddate   interface{}   `json:"resolvedDate"`
	Closedby       interface{}   `json:"closedBy"`
	Closeddate     time.Time     `json:"closedDate"`
	Lasteditedby   UserMeta      `json:"lastEditedBy"`
	Lastediteddate time.Time     `json:"lastEditedDate"`
	Needconfirm    bool          `json:"needconfirm,omitempty"` // 仅列表返回
}

type BugCreateMeta struct {
	BugMeta
	Openedbuild []string `json:"openedBuild"`
}

type BugGetMsg struct {
	BugBody
	Executionname      string        `json:"executionName"`
	Storytitle         string        `json:"storyTitle"`
	Storystatus        string        `json:"storyStatus"`
	Lateststoryversion int           `json:"latestStoryVersion"`
	Taskname           interface{}   `json:"taskName"`
	Planname           interface{}   `json:"planName"`
	Projectname        string        `json:"projectName"`
	Tocases            []interface{} `json:"toCases"`
	Files              []interface{} `json:"files"`
}

func (s *BugsService) ListByProducts(id int64) (*ListProductsBugsMsg, *req.Response, error) {
	var et ListProductsBugsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/bugs", id)))
	return &et, resp, err
}

// Create 创建Bug
func (s *BugsService) Create(id int, build BugCreateMeta) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/bugs", id)))
	return &u, resp, err
}

// DeleteByID 删除Bug
func (s *BugsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改Bug
func (s *BugsService) UpdateByID(id int, exec BugCreateMeta) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}

// GetByID 获取Bug详情
func (s *BugsService) GetByID(id int) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}
