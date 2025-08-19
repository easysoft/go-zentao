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
	Branch     int     `json:"branch,omitempty"` // 创建需求时, 该字段不写
	Module     int     `json:"module,omitempty"`
	Os         string  `json:"os,omitempty"`
	Browser    string  `json:"browser,omitempty"`
	Task       int     `json:"task,omitempty"`
	Story      int     `json:"story,omitempty"`
	Deadline   string  `json:"deadline,omitempty"`
	Title      string  `json:"title"`
	Project    int     `json:"project"`
	Execution  int     `json:"execution,omitempty"`
	Severity   int     `json:"severity"`
	Pri        int     `json:"pri"`
	Steps      string  `json:"steps"`
	Type       BugType `json:"type"`
	AssignedTo any     `json:"assignedTo"`
	FeedbackBy any     `json:"feedbackBy,omitempty"` // 仅bug接口
}

type BugBody struct {
	BugMeta
	ID             int            `json:"id"`
	Product        int            `json:"product"`
	Plan           int            `json:"plan"`
	StoryVersion   int            `json:"storyVersion"`
	PriOrder       string         `json:"priOrder"`
	SeverityOrder  int            `json:"severityOrder"`
	Totask         int            `json:"toTask"`
	Tostory        int            `json:"toStory"`
	Keywords       string         `json:"keywords"`
	Hardware       string         `json:"hardware"`
	Found          string         `json:"found"`
	Status         BugStatus      `json:"status,omitempty"`
	SubStatus      string         `json:"subStatus"`
	Color          string         `json:"color"`
	Confirmed      int            `json:"confirmed"`
	ActivatedCount int            `json:"activatedCount"`
	ActivatedDate  any            `json:"activatedDate"`
	Entry          string         `json:"entry"`
	Lines          string         `json:"lines"`
	V1             string         `json:"v1"`
	V2             string         `json:"v2"`
	DuplicateBug   int            `json:"duplicateBug"`
	LinkBug        string         `json:"linkBug"`
	Case           int            `json:"case"`
	CaseVersion    int            `json:"caseVersion"`
	Feedback       int            `json:"feedback"`
	Result         int            `json:"result"`
	Repo           int            `json:"repo"`
	MR             int            `json:"mr"`
	RepoType       string         `json:"repoType"`
	IssueKey       string         `json:"issueKey"`
	TestTask       int            `json:"testtask"`
	Deleted        any            `json:"deleted"`               // 字符串orbool
	NotifyEmail    string         `json:"notifyEmail,omitempty"` // 仅bug接口
	Mailto         any            `json:"mailto"`
	OpenedBy       any            `json:"openedBy"`
	OpenedDate     string         `json:"openedDate"`
	OpenedBuild    any            `json:"openedBuild"`
	AssignedDate   any            `json:"assignedDate"`
	ResolvedBy     any            `json:"resolvedBy"`
	Resolution     BugCloseReason `json:"resolution"`
	ResolvedBuild  string         `json:"resolvedBuild"`
	ResolvedDate   any            `json:"resolvedDate"`
	ClosedBy       any            `json:"closedBy"`
	ClosedDate     string         `json:"closedDate"`
	LasteditedBy   any            `json:"lastEditedBy"`
	LasteditedDate string         `json:"lastEditedDate"`
	NeedConfirm    bool           `json:"needconfirm,omitempty"` // 仅列表返回
	RelatedBug     string         `json:"relatedBug"`
	StatusName     string         `json:"statusName"`
	ProductStatus  string         `json:"productStatus"`
}

type BugCreateMeta struct {
	BugMeta
	AllBuilds   string   `json:"allBuilds,omitempty"` // 默认为all,所有版本
	Openedbuild []string `json:"openedBuild"`
}

type BugUpdateMeta struct {
	Title       string         `json:"title"`
	Project     int            `json:"project,omitempty"`
	Execution   int            `json:"execution,omitempty"`
	Openedbuild any            `json:"openedBuild,omitempty"`
	AssignedTo  any            `json:"assignedTo,omitempty"`
	Pri         int            `json:"pri,omitempty"`
	Severity    int            `json:"severity,omitempty"`
	Type        BugType        `json:"type,omitempty"`
	Story       int            `json:"story,omitempty"`
	ResolvedBy  any            `json:"resolvedBy,omitempty"`
	ClosedBy    any            `json:"closedBy,omitempty"`
	Resolution  BugCloseReason `json:"resolution,omitempty"`
	Product     int            `json:"product,omitempty"`
	Plan        int            `json:"plan,omitempty"`
	Task        int            `json:"task,omitempty"`
	Module      int            `json:"module,omitempty"`
	Steps       string         `json:"steps,omitempty"`
	Mailto      any            `json:"os,omitempty"`
	Keywords    any            `json:"keywords,omitempty"`
}

type BugGetMsg struct {
	BugBody
	ExecutionName      string `json:"executionName"`
	StoryTitle         string `json:"storyTitle"`
	StoryStatus        string `json:"storyStatus"`
	LatestStoryVersion any    `json:"latestStoryVersion"` // 未关联需求时string, 关联需求时int
	TaskName           any    `json:"taskName"`
	PlanName           any    `json:"planName"`
	ProjectName        string `json:"projectName"`
	ToCases            any    `json:"toCases"`
	Files              any    `json:"files"`
	LinkMRTitles       any    `json:"linkMRTitles"`
	ModuleTitle        string `json:"moduleTitle"`
	Actions            any    `json:"actions"`
	PreAndNext         any    `json:"preAndNext"`
}

type AssignBug struct {
	AssignedTo string `json:"assignedTo"`
	Mailto     any    `json:"mailto,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type ConfirmBug struct {
	AssignBug
	Pri      int       `json:"pri"`
	Type     BugType   `json:"type"`
	Status   BugStatus `json:"status"`
	Deadline string    `json:"deadline,omitempty"`
}

type ResolveBug struct {
	Resolution    BugCloseReason `json:"resolution"`
	ResolvedBuild string         `json:"resolvedBuild"`
	ResolvedDate  string         `json:"resolvedDate,omitempty"`
	DuplicateBug  any            `json:"duplicateBug,omitempty"`
	AssignedTo    string         `json:"assignedTo"`
	Comment       string         `json:"comment"`
}

type ActiveBug struct {
	AssignedTo  string `json:"assignedTo,omitempty"`
	Comment     string `json:"comment,omitempty"`
	OpenedBuild any    `json:"openedBuild"`
}

type CommentBug struct {
	Comment string `json:"comment,omitempty"`
}

func (s *BugsService) ListByProducts(id int64, op ListOptions) (*ListProductsBugsMsg, *req.Response, error) {
	var et ListProductsBugsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryString(op.toURLValues()).
		SetSuccessResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/bugs", id)))
	return &et, resp, err
}

func (s *BugsService) ListByProjects(id int64, op ListOptions) (*ListProductsBugsMsg, *req.Response, error) {
	var et ListProductsBugsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryString(op.toURLValues()).
		SetSuccessResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/bugs", id)))
	return &et, resp, err
}

func (s *BugsService) ListByExecutions(id int64, op ListOptions) (*ListProductsBugsMsg, *req.Response, error) {
	var et ListProductsBugsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryString(op.toURLValues()).
		SetSuccessResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/bugs", id)))
	return &et, resp, err
}

// Create 创建Bug
func (s *BugsService) Create(id int, build BugCreateMeta) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/bugs", id)))
	return &u, resp, err
}

// DeleteByID 删除Bug
func (s *BugsService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改Bug
func (s *BugsService) UpdateByID(id int, exec BugUpdateMeta) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}

// GetByID 获取Bug详情
func (s *BugsService) GetByID(id int) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/bugs/%d", id)))
	return &u, resp, err
}

// CloseByID 关闭Bug(需要已经resolve)
func (s *BugsService) CloseByID(id int, comment CommentBug) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&comment).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/bugs/%d/close", id)))
	return &u, resp, err
}

// AssignByID 指派Bug
func (s *BugsService) AssignByID(id int, assign AssignBug) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&assign).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/bugs/%d/assign", id)))
	return &u, resp, err
}

// ConfirmByID 确认Bug
func (s *BugsService) ConfirmByID(id int, confirm ConfirmBug) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&confirm).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/bugs/%d/confirm", id)))
	return &u, resp, err
}

// ResolveByID 解决Bug
func (s *BugsService) ResolveByID(id int, resolve ResolveBug) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&resolve).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/bugs/%d/resolve", id)))
	return &u, resp, err
}

// ActiveByID 激活Bug
func (s *BugsService) ActiveByID(id int, active ActiveBug) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&active).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/bugs/%d/active", id)))
	return &u, resp, err
}

// EstimateByID 评估Bug
func (s *BugsService) EstimateByID(id int) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/bugs/%d/estimate", id)))
	return &u, resp, err
}
