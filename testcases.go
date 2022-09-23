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

type TestCasesService struct {
	client *Client
}

type ListProductsTestCasesMsg struct {
	Page      int `json:"page"`
	Total     int `json:"total"`
	Limit     int `json:"limit"`
	Testcases []struct {
		ID             int         `json:"id"`
		Project        int         `json:"project"`
		Product        int         `json:"product"`
		Execution      int         `json:"execution"`
		Branch         int         `json:"branch"`
		Lib            int         `json:"lib"`
		Module         int         `json:"module"`
		Path           int         `json:"path"`
		Story          int         `json:"story"`
		Storyversion   int         `json:"storyVersion"`
		Title          string      `json:"title"`
		Precondition   string      `json:"precondition"`
		Keywords       string      `json:"keywords"`
		Pri            int         `json:"pri"`
		Type           string      `json:"type"`
		Auto           string      `json:"auto"`
		Frame          string      `json:"frame"`
		Stage          string      `json:"stage"`
		Howrun         string      `json:"howRun"`
		Scriptedby     string      `json:"scriptedBy"`
		Scripteddate   interface{} `json:"scriptedDate"`
		Scriptstatus   string      `json:"scriptStatus"`
		Scriptlocation string      `json:"scriptLocation"`
		Status         string      `json:"status"`
		Substatus      string      `json:"subStatus"`
		Color          string      `json:"color"`
		Frequency      string      `json:"frequency"`
		Order          int         `json:"order"`
		Openedby       struct {
			ID       int    `json:"id"`
			Account  string `json:"account"`
			Avatar   string `json:"avatar"`
			Realname string `json:"realname"`
		} `json:"openedBy"`
		Openeddate      time.Time   `json:"openedDate"`
		Reviewedby      interface{} `json:"reviewedBy"`
		Revieweddate    interface{} `json:"reviewedDate"`
		Lasteditedby    interface{} `json:"lastEditedBy"`
		Lastediteddate  interface{} `json:"lastEditedDate"`
		Version         int         `json:"version"`
		Linkcase        string      `json:"linkCase"`
		Frombug         int         `json:"fromBug"`
		Fromcaseid      int         `json:"fromCaseID"`
		Fromcaseversion int         `json:"fromCaseVersion"`
		Deleted         bool        `json:"deleted"`
		Lastrunner      string      `json:"lastRunner"`
		Lastrundate     interface{} `json:"lastRunDate"`
		Lastrunresult   string      `json:"lastRunResult"`
		Storytitle      interface{} `json:"storyTitle"`
		Needconfirm     bool        `json:"needconfirm"`
		Bugs            int         `json:"bugs"`
		Results         int         `json:"results"`
		Casefails       int         `json:"caseFails"`
		Stepnumber      int         `json:"stepNumber"`
		Statusname      string      `json:"statusName"`
	} `json:"testcases"`
}

type TestCasesCreateMeta struct {
	Branch       int             `json:"branch,omitempty"`
	Module       int             `json:"module,omitempty"`
	Story        int             `json:"story,omitempty"`
	Stage        string          `json:"stage,omitempty"`
	Precondition string          `json:"precondition,omitempty"`
	Title        string          `json:"title"`
	Keywords     string          `json:"keywords,omitempty"`
	Pri          int             `json:"pri,omitempty"`
	Steps        []TestCasesStep `json:"steps"`
	Type         string          `json:"type"`
}

type TestCasesStep struct {
	Desc   string `json:"json"`
	Expect string `json:"expect"`
}

type TestCasesCreateMsg struct {
	ID              int           `json:"id"`
	Project         int           `json:"project"`
	Product         int           `json:"product"`
	Execution       int           `json:"execution"`
	Branch          int           `json:"branch"`
	Lib             int           `json:"lib"`
	Module          int           `json:"module"`
	Path            int           `json:"path"`
	Story           int           `json:"story"`
	Storyversion    int           `json:"storyVersion"`
	Title           string        `json:"title"`
	Precondition    string        `json:"precondition"`
	Keywords        string        `json:"keywords"`
	Pri             int           `json:"pri"`
	Type            string        `json:"type"`
	Auto            string        `json:"auto"`
	Frame           string        `json:"frame"`
	Stage           string        `json:"stage"`
	Howrun          string        `json:"howRun"`
	Scriptedby      string        `json:"scriptedBy"`
	Scripteddate    interface{}   `json:"scriptedDate"`
	Scriptstatus    string        `json:"scriptStatus"`
	Scriptlocation  string        `json:"scriptLocation"`
	Status          string        `json:"status"`
	Substatus       string        `json:"subStatus"`
	Color           string        `json:"color"`
	Frequency       string        `json:"frequency"`
	Order           int           `json:"order"`
	Openedby        string        `json:"openedBy"`
	Openeddate      time.Time     `json:"openedDate"`
	Reviewedby      string        `json:"reviewedBy"`
	Revieweddate    interface{}   `json:"reviewedDate"`
	Lasteditedby    string        `json:"lastEditedBy"`
	Lastediteddate  interface{}   `json:"lastEditedDate"`
	Version         int           `json:"version"`
	Linkcase        string        `json:"linkCase"`
	Frombug         int           `json:"fromBug"`
	Fromcaseid      int           `json:"fromCaseID"`
	Fromcaseversion int           `json:"fromCaseVersion"`
	Deleted         bool          `json:"deleted"`
	Lastrunner      string        `json:"lastRunner"`
	Lastrundate     interface{}   `json:"lastRunDate"`
	Lastrunresult   string        `json:"lastRunResult"`
	Tobugs          []interface{} `json:"toBugs"`
	Steps           []struct {
		ID      int    `json:"id"`
		Parent  int    `json:"parent"`
		Case    int    `json:"case"`
		Version int    `json:"version"`
		Type    string `json:"type"`
		Desc    string `json:"desc"`
		Expect  string `json:"expect"`
	} `json:"steps"`
	Files          []interface{} `json:"files"`
	Currentversion int           `json:"currentVersion"`
}

type TestCasesGetMsg struct {
	ID                 int           `json:"id"`
	Project            int           `json:"project"`
	Product            int           `json:"product"`
	Branch             int           `json:"branch"`
	Module             int           `json:"module"`
	Execution          int           `json:"execution"`
	Plan               int           `json:"plan"`
	Story              int           `json:"story"`
	Storyversion       int           `json:"storyVersion"`
	Task               int           `json:"task"`
	Totask             int           `json:"toTask"`
	Tostory            int           `json:"toStory"`
	Title              string        `json:"title"`
	Keywords           string        `json:"keywords"`
	Severity           int           `json:"severity"`
	Pri                int           `json:"pri"`
	Type               string        `json:"type"`
	Os                 string        `json:"os"`
	Browser            string        `json:"browser"`
	Hardware           string        `json:"hardware"`
	Found              string        `json:"found"`
	Steps              string        `json:"steps"`
	Status             string        `json:"status"`
	Substatus          string        `json:"subStatus"`
	Color              string        `json:"color"`
	Confirmed          int           `json:"confirmed"`
	Activatedcount     int           `json:"activatedCount"`
	Activateddate      time.Time     `json:"activatedDate"`
	Feedbackby         string        `json:"feedbackBy"`
	Notifyemail        string        `json:"notifyEmail"`
	Mailto             []interface{} `json:"mailto"`
	Openedby           UserMeta      `json:"openedBy"`
	Openeddate         time.Time     `json:"openedDate"`
	Openedbuild        string        `json:"openedBuild"`
	Assignedto         UserMeta      `json:"assignedTo"`
	Assigneddate       time.Time     `json:"assignedDate"`
	Deadline           interface{}   `json:"deadline"`
	Resolvedby         UserMeta      `json:"resolvedBy"`
	Resolution         string        `json:"resolution"`
	Resolvedbuild      string        `json:"resolvedBuild"`
	Resolveddate       interface{}   `json:"resolvedDate"`
	Closedby           UserMeta      `json:"closedBy"`
	Closeddate         time.Time     `json:"closedDate"`
	Duplicatebug       int           `json:"duplicateBug"`
	Linkbug            string        `json:"linkBug"`
	Case               int           `json:"case"`
	Caseversion        int           `json:"caseVersion"`
	Result             int           `json:"result"`
	Repo               int           `json:"repo"`
	Entry              string        `json:"entry"`
	Lines              string        `json:"lines"`
	V1                 string        `json:"v1"`
	V2                 string        `json:"v2"`
	Repotype           string        `json:"repoType"`
	Testtask           int           `json:"testtask"`
	Lasteditedby       UserMeta      `json:"lastEditedBy"`
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

func (s *TestCasesService) ListByProducts(id int64) (*ListProductsTestCasesMsg, *req.Response, error) {
	var et ListProductsTestCasesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/testcases", id)))
	return &et, resp, err
}

// Create 创建用例
func (s *TestCasesService) Create(id int, build TestCasesCreateMeta) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/testcases", id)))
	return &u, resp, err
}

// DeleteByID 删除用例
func (s *TestCasesService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改用例
func (s *TestCasesService) UpdateByID(id int, tc TestCasesCreateMeta) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&tc).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}

// GetByID 获取用例详情
func (s *TestCasesService) GetByID(id int) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}
