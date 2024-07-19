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
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type TestCasesService struct {
	client *Client
}

type TestCasesMeta struct {
	ID             any         `json:"id"` // 老版本int, 新版本string
	Project        int         `json:"project"`
	Product        int         `json:"product"`
	Execution      int         `json:"execution"`
	Branch         int         `json:"branch"`
	Module         int         `json:"module"`
	Story          int         `json:"story"`
	Storyversion   int         `json:"storyVersion"`
	Title          string      `json:"title"`
	Keywords       string      `json:"keywords"`
	Pri            int         `json:"pri"`
	Type           string      `json:"type"`
	Status         string      `json:"status"`
	Substatus      string      `json:"subStatus"`
	Color          string      `json:"color"`
	Openedby       UserMeta    `json:"openedBy"`
	Openeddate     time.Time   `json:"openedDate"`
	Lasteditedby   interface{} `json:"lastEditedBy"`
	Lastediteddate interface{} `json:"lastEditedDate"`
	Deleted        bool        `json:"deleted"`
	Storytitle     interface{} `json:"storyTitle,omitempty"`
}

type TestcasesListBody struct {
	TestCasesMeta
	Lib             int    `json:"lib"`
	Path            int    `json:"path"`
	Precondition    string `json:"precondition"`
	Auto            string `json:"auto"`
	Frame           string `json:"frame"`
	Stage           string `json:"stage"`
	Howrun          string `json:"howRun"`
	Scriptedby      string `json:"scriptedBy"`
	Scripteddate    any    `json:"scriptedDate"`
	Scriptstatus    string `json:"scriptStatus"`
	Scriptlocation  string `json:"scriptLocation"`
	Frequency       string `json:"frequency"`
	Order           int    `json:"order"`
	Reviewedby      any    `json:"reviewedBy"`
	Revieweddate    any    `json:"reviewedDate"`
	Version         int    `json:"version"`
	Linkcase        string `json:"linkCase"`
	Frombug         int    `json:"fromBug"`
	Fromcaseid      int    `json:"fromCaseID"`
	Fromcaseversion int    `json:"fromCaseVersion"`
	Lastrunner      string `json:"lastRunner"`
	Lastrundate     any    `json:"lastRunDate"`
	Lastrunresult   string `json:"lastRunResult"`
	Needconfirm     bool   `json:"needconfirm"`
	Bugs            int    `json:"bugs"`
	Results         int    `json:"results"`
	Casefails       int    `json:"caseFails"`
	Stepnumber      int    `json:"stepNumber"`
	Statusname      string `json:"statusName"`
}

type ListProductsTestCasesMsg struct {
	Page      int                 `json:"page"`
	Total     int                 `json:"total"`
	Limit     int                 `json:"limit"`
	Testcases []TestcasesListBody `json:"testcases"`
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

type TestCasesCreateMsg struct {
	TestCasesMeta
	Lib             int             `json:"lib"`
	Path            int             `json:"path"`
	Precondition    string          `json:"precondition"`
	Auto            string          `json:"auto"`
	Frame           string          `json:"frame"`
	Stage           string          `json:"stage"`
	Howrun          string          `json:"howRun"`
	Scriptedby      string          `json:"scriptedBy"`
	Scripteddate    any             `json:"scriptedDate"`
	Scriptstatus    string          `json:"scriptStatus"`
	Scriptlocation  string          `json:"scriptLocation"`
	Frequency       string          `json:"frequency"`
	Order           int             `json:"order"`
	Reviewedby      string          `json:"reviewedBy"`
	Revieweddate    any             `json:"reviewedDate"`
	Version         int             `json:"version"`
	Linkcase        string          `json:"linkCase"`
	Frombug         int             `json:"fromBug"`
	Fromcaseid      int             `json:"fromCaseID"`
	Fromcaseversion int             `json:"fromCaseVersion"`
	Lastrunner      string          `json:"lastRunner"`
	Lastrundate     any             `json:"lastRunDate"`
	Lastrunresult   string          `json:"lastRunResult"`
	Tobugs          []interface{}   `json:"toBugs"`
	Steps           []TestCasesStep `json:"steps"`
	Files           []interface{}   `json:"files"`
	Currentversion  int             `json:"currentVersion"`
}

type TestCasesStep struct {
	ID      int    `json:"id"`
	Parent  int    `json:"parent"`
	Case    int    `json:"case"`
	Version int    `json:"version"`
	Type    string `json:"type"`
	Desc    string `json:"desc"`
	Expect  string `json:"expect"`
}

type TestCasesGetMsg struct {
	TestCasesMeta
	Plan               int           `json:"plan"`
	Task               int           `json:"task"`
	Totask             int           `json:"toTask"`
	Tostory            int           `json:"toStory"`
	Severity           int           `json:"severity"`
	Os                 string        `json:"os"`
	Browser            string        `json:"browser"`
	Hardware           string        `json:"hardware"`
	Found              string        `json:"found"`
	Steps              string        `json:"steps"`
	Confirmed          int           `json:"confirmed"`
	Activatedcount     int           `json:"activatedCount"`
	Activateddate      time.Time     `json:"activatedDate"`
	Feedbackby         string        `json:"feedbackBy"`
	Notifyemail        string        `json:"notifyEmail"`
	Mailto             []interface{} `json:"mailto"`
	Openedbuild        string        `json:"openedBuild"`
	Assignedto         UserMeta      `json:"assignedTo"`
	Assigneddate       time.Time     `json:"assignedDate"`
	Deadline           any           `json:"deadline"`
	Resolvedby         UserMeta      `json:"resolvedBy"`
	Resolution         string        `json:"resolution"`
	Resolvedbuild      string        `json:"resolvedBuild"`
	Resolveddate       any           `json:"resolvedDate"`
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
	Executionname      string        `json:"executionName"`
	Storystatus        string        `json:"storyStatus"`
	Lateststoryversion int           `json:"latestStoryVersion"`
	Taskname           any           `json:"taskName"`
	Planname           any           `json:"planName"`
	Projectname        string        `json:"projectName"`
	Tocases            []any         `json:"toCases"`
	Files              []any         `json:"files"`
}

type TestCasesResult struct {
	ID          int                   `json:"id"`
	Run         int                   `json:"run"`
	Case        int                   `json:"case"`
	Version     int                   `json:"version"`
	Job         int                   `json:"job"`
	Compile     int                   `json:"compile"`
	CaseResult  string                `json:"caseResult"` // 用例结果 pass
	StepResults []TestCasesStepResult `json:"stepResults"`
	ZTFResult   string                `json:"ZTFResult"`
	Node        int                   `json:"node"`
	Lastrunner  string                `json:"lastRunner"`
	Date        string                `json:"date"`
	Duration    int                   `json:"duration"`
	Xml         any                   `json:"xml"`
	Deploy      int                   `json:"deploy"`
	Build       any                   `json:"build"`
	Task        int                   `json:"task"`
	NodeName    string                `json:"nodeName"`
	Files       []any                 `json:"files"`
}

type TestCasesStepResult struct {
	ID      int    `json:"id"`
	Parent  int    `json:"parent"`
	Case    int    `json:"case"`
	Version int    `json:"version"`
	Type    string `json:"type"` // 步骤类型 step
	Desc    string `json:"desc"`
	Expect  string `json:"expect"`
	Result  string `json:"result"`
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Real    string `json:"real"`
	Files   []any  `json:"files"`
}

type TestCasesResultResp struct {
	Results []TestCasesResult `json:"results"`
}

func (s *TestCasesService) ListByProducts(id int, parms ...any) (*ListProductsTestCasesMsg, *req.Response, error) {
	var et ListProductsTestCasesMsg
	req := s.client.client.R().
		SetHeader("Token", s.client.token)
	if len(parms) > 0 {
		req = req.SetQueryParams(parms[0].(map[string]string))
	}
	resp, err := req.
		SetSuccessResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/testcases", id)))
	return &et, resp, err
}

// Create 创建用例
func (s *TestCasesService) Create(id int, build TestCasesCreateMeta) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/testcases", id)))
	return &u, resp, err
}

// DeleteByID 删除用例
func (s *TestCasesService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改用例
func (s *TestCasesService) UpdateByID(id int, tc TestCasesCreateMeta) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&tc).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}

// GetByID 获取用例详情
func (s *TestCasesService) GetByID(id int) (*TestCasesCreateMsg, *req.Response, error) {
	var u TestCasesCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/testcases/%d", id)))
	return &u, resp, err
}

// GetByID 获取用例详情
func (s *TestCasesService) GetResultByID(id any) (*TestCasesResultResp, *req.Response, error) {
	var u TestCasesResultResp
	// 判断id的类型，如果是string，就是用例id，如果是int，就是结果id
	if _, ok := id.(string); ok {
		if strings.HasPrefix(id.(string), "case_") {
			id = strings.TrimPrefix(id.(string), "case_")
		}
	}
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/testcases/%v/results", id)))
	return &u, resp, err
}
