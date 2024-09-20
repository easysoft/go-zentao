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
	Source     StoriesSource   `json:"source,omitempty"`
	SourceNote string          `json:"sourceNote,omitempty"` // 来源备注
	Pri        int             `json:"pri,omitempty"`
	Category   StoriesCategory `json:"category,omitempty"`
	Estimate   float64         `json:"estimate,omitempty"` // 预估工时
	Keywords   string          `json:"keywords,omitempty"`
	Parent     any             `json:"parent,omitempty"` // 可能是数组, 产品计划时是数组
	Reviewer   []string        `json:"reviewer,omitempty"`
}

type StoriesBody struct {
	StoriesExtMeta
	ID                  int           `json:"id"`
	Product             int           `json:"product"`
	Project             int           `json:"project,omitempty"`
	Branch              int           `json:"branch"`
	Module              int           `json:"module"`
	Plan                string        `json:"plan"`
	Frombug             int           `json:"fromBug"`
	Title               string        `json:"title"`
	Type                string        `json:"type"`
	Status              StoriesStatus `json:"status"`
	Substatus           string        `json:"subStatus"`
	Color               string        `json:"color"`
	Stage               StoriesStage  `json:"stage"`
	Stagedby            any           `json:"stagedBy"`
	Mailto              any           `json:"mailto"`   // 概率数组
	Openedby            any           `json:"openedBy"` // 产品计划是string, 其他可能是UserMeta
	Openeddate          string        `json:"openedDate"`
	Assignedto          any           `json:"assignedTo"`
	Assigneddate        string        `json:"assignedDate"`
	Lasteditedby        any           `json:"lastEditedBy"`
	Lastediteddate      string        `json:"lastEditedDate"`
	Reviewedby          any           `json:"reviewedBy"`
	Revieweddate        string        `json:"reviewedDate"`
	Closedby            any           `json:"closedBy"`
	Closeddate          string        `json:"closedDate"`
	Closedreason        string        `json:"closedReason"`
	Tobug               int           `json:"toBug"`
	Childstories        string        `json:"childStories"`
	Linkstories         string        `json:"linkStories"`
	Duplicatestory      int           `json:"duplicateStory"`
	Version             int           `json:"version"`
	Urchanged           string        `json:"URChanged"`
	Deleted             any           `json:"deleted"` // 需求返回时是bool, 产品计划是string
	Plantitle           string        `json:"planTitle,omitempty"`
	NotReview           []string      `json:"notReview,omitempty"`
	NeedSummaryEstimate bool          `json:"needSummaryEstimate,omitempty"`
	ProductStatus       string        `json:"productStatus,omitempty"`
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
	Children   []interface{} `json:"children,omitempty"`
}

type StoriesClose struct {
	Closedreason   StoriesCloseReason `json:"closedReason,omitempty"`
	Duplicatestory int                `json:"duplicateStory,omitempty"`
	Comment        string             `json:"comment,omitempty"`
}

type StoriesActive struct {
	AssignedTo string `json:"assignedTo,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type StoriesEstimate struct {
}

type StoriesReview struct {
	ReviewedDate string              `json:"reviewedDate,omitempty"`
	Result       StoriesReviewResult `json:"result,omitempty"`
	ClosedReason StoriesCloseReason  `json:"closedReason,omitempty"`
	Pri          int                 `json:"pri,omitempty"`
	Estimate     float64             `json:"estimate"`
	Comment      string              `json:"comment,omitempty"`
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

// CloseByID 关闭需求
func (s *StoriesService) CloseByID(id int, story StoriesClose) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/close", id)))
	return &u, resp, err
}

// ActiveByID 激活需求
func (s *StoriesService) ActiveByID(id int, story StoriesActive) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/active", id)))
	return &u, resp, err
}

// AssignByID 指派需求
func (s *StoriesService) AssignByID(id int, story StoriesActive) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/assign", id)))
	return &u, resp, err
}

// GetEstimateByID 获取工时, 开源版不支持
func (s *StoriesService) GetEstimateByID(id int) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/stories/%d/estimate", id)))
	return &u, resp, err
}

// UpdateEstimateByID 更新工时
func (s *StoriesService) UpdateEstimateByID(id int, story StoriesEstimate) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/estimate", id)))
	return &u, resp, err
}

// ChildByID 子需求
func (s *StoriesService) ChildByID(id int, story StoriesMeta) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/child", id)))
	return &u, resp, err
}

// RecallByID 撤回评审
func (s *StoriesService) RecallByID(id int) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/stories/%d/recall", id)))
	return &u, resp, err
}

// ReviewByID 审核需求
func (s *StoriesService) ReviewByID(id int, story StoriesReview) (*StoriesMsg, *req.Response, error) {
	var u StoriesMsg
	story.ClosedReason = CloseReasonDone
	if story.Estimate <= 0 {
		story.Estimate = 1.0
	}
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/stories/%d/review", id)))
	return &u, resp, err
}
