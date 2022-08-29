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

type BugsService struct {
	client *Client
}

type ListProductsBugsMsg struct {
	Page  int `json:"page"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Bugs  []struct {
		ID           int    `json:"id"`
		Project      int    `json:"project"`
		Product      int    `json:"product"`
		Branch       int    `json:"branch"`
		Module       int    `json:"module"`
		Execution    int    `json:"execution"`
		Plan         int    `json:"plan"`
		Story        int    `json:"story"`
		Storyversion int    `json:"storyVersion"`
		Task         int    `json:"task"`
		Totask       int    `json:"toTask"`
		Tostory      int    `json:"toStory"`
		Title        string `json:"title"`
		Keywords     string `json:"keywords"`
		Severity     int    `json:"severity"`
		Pri          int    `json:"pri"`
		Type         string `json:"type"`
		Os           string `json:"os"`
		Browser      string `json:"browser"`
		Hardware     string `json:"hardware"`
		Found        string `json:"found"`
		Steps        string `json:"steps"`
		Status       struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"status"`
		Substatus      string      `json:"subStatus"`
		Color          string      `json:"color"`
		Confirmed      int         `json:"confirmed"`
		Activatedcount int         `json:"activatedCount"`
		Activateddate  time.Time   `json:"activatedDate"`
		Feedbackby     string      `json:"feedbackBy"`
		Notifyemail    string      `json:"notifyEmail"`
		Mailto         interface{} `json:"mailto"`
		Openedby       struct {
			ID       int    `json:"id"`
			Account  string `json:"account"`
			Avatar   string `json:"avatar"`
			Realname string `json:"realname"`
		} `json:"openedBy"`
		Openeddate     time.Time   `json:"openedDate"`
		Openedbuild    string      `json:"openedBuild"`
		Assignedto     interface{} `json:"assignedTo"`
		Assigneddate   time.Time   `json:"assignedDate"`
		Deadline       string      `json:"deadline"`
		Resolvedby     interface{} `json:"resolvedBy"`
		Resolution     string      `json:"resolution"`
		Resolvedbuild  string      `json:"resolvedBuild"`
		Resolveddate   time.Time   `json:"resolvedDate"`
		Closedby       interface{} `json:"closedBy"`
		Closeddate     time.Time   `json:"closedDate"`
		Duplicatebug   int         `json:"duplicateBug"`
		Linkbug        string      `json:"linkBug"`
		Case           int         `json:"case"`
		Caseversion    int         `json:"caseVersion"`
		Result         int         `json:"result"`
		Repo           int         `json:"repo"`
		Entry          string      `json:"entry"`
		Lines          string      `json:"lines"`
		V1             string      `json:"v1"`
		V2             string      `json:"v2"`
		Repotype       string      `json:"repoType"`
		Testtask       int         `json:"testtask"`
		Lasteditedby   interface{} `json:"lastEditedBy"`
		Lastediteddate time.Time   `json:"lastEditedDate"`
		Deleted        bool        `json:"deleted"`
		Needconfirm    bool        `json:"needconfirm"`
	} `json:"bugs"`
}

type BugCreateMeta struct {
	Branch      int       `json:"branch,omitempty"`
	Module      int       `json:"module,omitempty"`
	Execution   int       `json:"execution,omitempty"`
	Os          string    `json:"os,omitempty"`
	Browser     string    `json:"browser,omitempty"`
	Task        int       `json:"task,omitempty"`
	Story       int       `json:"story,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Title       string    `json:"title"`
	Severity    int       `json:"severity"`
	Pri         int       `json:"pri"`
	Steps       string    `json:"steps"`
	Type        string    `json:"type"`
	Openedbuild []string  `json:"openedBuild"`
}

type BugCreateMsg struct {
	ID             int           `json:"id"`
	Project        int           `json:"project"`
	Product        int           `json:"product"`
	Branch         int           `json:"branch"`
	Module         int           `json:"module"`
	Execution      int           `json:"execution"`
	Plan           int           `json:"plan"`
	Story          int           `json:"story"`
	Storyversion   int           `json:"storyVersion"`
	Task           int           `json:"task"`
	Totask         int           `json:"toTask"`
	Tostory        int           `json:"toStory"`
	Title          string        `json:"title"`
	Keywords       string        `json:"keywords"`
	Severity       int           `json:"severity"`
	Pri            int           `json:"pri"`
	Type           string        `json:"type"`
	Os             string        `json:"os"`
	Browser        string        `json:"browser"`
	Hardware       string        `json:"hardware"`
	Found          string        `json:"found"`
	Steps          string        `json:"steps"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Color          string        `json:"color"`
	Confirmed      int           `json:"confirmed"`
	Activatedcount int           `json:"activatedCount"`
	Activateddate  time.Time     `json:"activatedDate"`
	Feedbackby     string        `json:"feedbackBy"`
	Notifyemail    string        `json:"notifyEmail"`
	Mailto         []interface{} `json:"mailto"`
	Openedby       struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"openedBy"`
	Openeddate  time.Time `json:"openedDate"`
	Openedbuild string    `json:"openedBuild"`
	Assignedto  struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"assignedTo"`
	Assigneddate  time.Time   `json:"assignedDate"`
	Deadline      interface{} `json:"deadline"`
	Resolvedby    interface{} `json:"resolvedBy"`
	Resolution    string      `json:"resolution"`
	Resolvedbuild string      `json:"resolvedBuild"`
	Resolveddate  interface{} `json:"resolvedDate"`
	Closedby      interface{} `json:"closedBy"`
	Closeddate    time.Time   `json:"closedDate"`
	Duplicatebug  int         `json:"duplicateBug"`
	Linkbug       string      `json:"linkBug"`
	Case          int         `json:"case"`
	Caseversion   int         `json:"caseVersion"`
	Result        int         `json:"result"`
	Repo          int         `json:"repo"`
	Entry         string      `json:"entry"`
	Lines         string      `json:"lines"`
	V1            string      `json:"v1"`
	V2            string      `json:"v2"`
	Repotype      string      `json:"repoType"`
	Testtask      int         `json:"testtask"`
	Lasteditedby  struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"lastEditedBy"`
	Lastediteddate     time.Time     `json:"lastEditedDate"`
	Deleted            bool          `json:"deleted"`
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

type BugGetMsg struct {
	ID             int           `json:"id"`
	Project        int           `json:"project"`
	Product        int           `json:"product"`
	Branch         int           `json:"branch"`
	Module         int           `json:"module"`
	Execution      int           `json:"execution"`
	Plan           int           `json:"plan"`
	Story          int           `json:"story"`
	Storyversion   int           `json:"storyVersion"`
	Task           int           `json:"task"`
	Totask         int           `json:"toTask"`
	Tostory        int           `json:"toStory"`
	Title          string        `json:"title"`
	Keywords       string        `json:"keywords"`
	Severity       int           `json:"severity"`
	Pri            int           `json:"pri"`
	Type           string        `json:"type"`
	Os             string        `json:"os"`
	Browser        string        `json:"browser"`
	Hardware       string        `json:"hardware"`
	Found          string        `json:"found"`
	Steps          string        `json:"steps"`
	Status         string        `json:"status"`
	Substatus      string        `json:"subStatus"`
	Color          string        `json:"color"`
	Confirmed      int           `json:"confirmed"`
	Activatedcount int           `json:"activatedCount"`
	Activateddate  time.Time     `json:"activatedDate"`
	Feedbackby     string        `json:"feedbackBy"`
	Notifyemail    string        `json:"notifyEmail"`
	Mailto         []interface{} `json:"mailto"`
	Openedby       struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"openedBy"`
	Openeddate  time.Time `json:"openedDate"`
	Openedbuild string    `json:"openedBuild"`
	Assignedto  struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"assignedTo"`
	Assigneddate  time.Time   `json:"assignedDate"`
	Deadline      interface{} `json:"deadline"`
	Resolvedby    interface{} `json:"resolvedBy"`
	Resolution    string      `json:"resolution"`
	Resolvedbuild string      `json:"resolvedBuild"`
	Resolveddate  interface{} `json:"resolvedDate"`
	Closedby      interface{} `json:"closedBy"`
	Closeddate    time.Time   `json:"closedDate"`
	Duplicatebug  int         `json:"duplicateBug"`
	Linkbug       string      `json:"linkBug"`
	Case          int         `json:"case"`
	Caseversion   int         `json:"caseVersion"`
	Result        int         `json:"result"`
	Repo          int         `json:"repo"`
	Entry         string      `json:"entry"`
	Lines         string      `json:"lines"`
	V1            string      `json:"v1"`
	V2            string      `json:"v2"`
	Repotype      string      `json:"repoType"`
	Testtask      int         `json:"testtask"`
	Lasteditedby  struct {
		ID       int    `json:"id"`
		Account  string `json:"account"`
		Avatar   string `json:"avatar"`
		Realname string `json:"realname"`
	} `json:"lastEditedBy"`
	Lastediteddate     time.Time     `json:"lastEditedDate"`
	Deleted            bool          `json:"deleted"`
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
func (s *BugsService) Create(id int, build BugCreateMeta) (*BugCreateMsg, *req.Response, error) {
	var u BugCreateMsg
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
func (s *BugsService) UpdateByID(id int, exec BugCreateMeta) (*BugCreateMsg, *req.Response, error) {
	var u BugCreateMsg
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
