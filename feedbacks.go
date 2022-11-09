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

type FeedBacksService struct {
	client *Client
}

type ListFeedBacksMsg struct {
	Page      int `json:"page"`
	Total     int `json:"total"`
	Limit     int `json:"limit"`
	Feedbacks []struct {
		ID            int           `json:"id"`
		Product       int           `json:"product"`
		Module        int           `json:"module"`
		Title         string        `json:"title"`
		Type          string        `json:"type"`
		Solution      string        `json:"solution"`
		Desc          string        `json:"desc"`
		Status        string        `json:"status"`
		Substatus     string        `json:"subStatus"`
		Public        string        `json:"public"`
		Notify        string        `json:"notify"`
		Notifyemail   string        `json:"notifyEmail"`
		Likes         string        `json:"likes"`
		Result        int           `json:"result"`
		Faq           int           `json:"faq"`
		Openedby      UserMeta      `json:"openedBy"`
		Openeddate    time.Time     `json:"openedDate"`
		Reviewedby    interface{}   `json:"reviewedBy"`
		Revieweddate  interface{}   `json:"reviewedDate"`
		Processedby   interface{}   `json:"processedBy"`
		Processeddate interface{}   `json:"processedDate"`
		Closedby      interface{}   `json:"closedBy"`
		Closeddate    interface{}   `json:"closedDate"`
		Closedreason  string        `json:"closedReason"`
		Editedby      UserMeta      `json:"editedBy"`
		Editeddate    time.Time     `json:"editedDate"`
		Assignedto    interface{}   `json:"assignedTo"`
		Assigneddate  string        `json:"assignedDate"`
		Feedbackby    string        `json:"feedbackBy"`
		Mailto        []interface{} `json:"mailto"`
		Deleted       bool          `json:"deleted"`
		Dept          int           `json:"dept"`
	} `json:"feedbacks"`
}

type FeedBacksGetMsg struct {
	ID            int           `json:"id"`
	Product       int           `json:"product"`
	Module        int           `json:"module"`
	Title         string        `json:"title"`
	Type          string        `json:"type"`
	Solution      string        `json:"solution"`
	Desc          string        `json:"desc"`
	Status        string        `json:"status"`
	Substatus     string        `json:"subStatus"`
	Public        string        `json:"public"`
	Notify        string        `json:"notify"`
	Notifyemail   string        `json:"notifyEmail"`
	Likes         string        `json:"likes"`
	Result        int           `json:"result"`
	Faq           int           `json:"faq"`
	Openedby      UserMeta      `json:"openedBy"`
	Openeddate    time.Time     `json:"openedDate"`
	Reviewedby    string        `json:"reviewedBy"`
	Revieweddate  string        `json:"reviewedDate"`
	Processedby   string        `json:"processedBy"`
	Processeddate string        `json:"processedDate"`
	Closedby      interface{}   `json:"closedBy"`
	Closeddate    interface{}   `json:"closedDate"`
	Closedreason  string        `json:"closedReason"`
	Editedby      string        `json:"editedBy"`
	Editeddate    string        `json:"editedDate"`
	Assignedto    interface{}   `json:"assignedTo"`
	Assigneddate  time.Time     `json:"assignedDate"`
	Feedbackby    string        `json:"feedbackBy"`
	Mailto        []interface{} `json:"mailto"`
	Deleted       bool          `json:"deleted"`
	Likescount    int           `json:"likesCount"`
	Files         []interface{} `json:"files"`
	Likeicon      string        `json:"likeIcon"`
	Publicstatus  string        `json:"publicStatus"`
	Productname   string        `json:"productName"`
	Modulename    string        `json:"moduleName"`
	Resulttype    string        `json:"resultType"`
	Actions       []struct {
		ID         int           `json:"id"`
		Objecttype string        `json:"objectType"`
		Objectid   int           `json:"objectID"`
		Product    string        `json:"product"`
		Project    int           `json:"project"`
		Execution  int           `json:"execution"`
		Actor      string        `json:"actor"`
		Action     string        `json:"action"`
		Date       string        `json:"date"`
		Comment    string        `json:"comment"`
		Extra      string        `json:"extra"`
		Read       string        `json:"read"`
		Vision     string        `json:"vision"`
		Efforted   int           `json:"efforted"`
		History    []interface{} `json:"history"`
		Desc       string        `json:"desc"`
	} `json:"actions"`
}

type FeedBacksCreateMeta struct {
	Product     int    `json:"product"`
	Module      int    `json:"module,omitempty"`
	Public      int    `json:"public,omitempty"`
	Notify      int    `json:"notify,omitempty"`
	Title       string `json:"title"`
	Type        string `json:"type,omitempty"`
	Desc        string `json:"Desc,omitempty"`
	NotifyEmail string `json:"notifyEmail,omitempty"`
	FeedbackBy  string `json:"feedbackBy,omitempty"`
}

type FeedBacksAssign struct {
	Assignedto string `json:"assignedTo"`
	Comment    string `json:"comment,omitempty"`
	Mailto     string `json:"mailto,omitempty"`
}

type FeedBacksClose struct {
	ClosedReason string `json:"closedReason"`
	Comment      string `json:"comment,omitempty"`
}

func (s *FeedBacksService) List(solution, orderBy, page, limit string) (*ListFeedBacksMsg, *req.Response, error) {
	var et ListFeedBacksMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"solution": solution,
			"orderBy":  orderBy,
			"page":     page,
			"limit":    limit,
		}).
		SetResult(&et).
		Get(s.client.RequestURL("/feedbacks"))
	return &et, resp, err
}

// Create 创建反馈
func (s *FeedBacksService) Create(fb FeedBacksCreateMeta) (*FeedBacksGetMsg, *req.Response, error) {
	var u FeedBacksGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&fb).
		SetResult(&u).
		Post(s.client.RequestURL("/feedbacks"))
	return &u, resp, err
}

// DeleteByID 删除反馈
func (s *FeedBacksService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/feedbacks/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改反馈
func (s *FeedBacksService) UpdateByID(id int, exec FeedBacksCreateMeta) (*FeedBacksGetMsg, *req.Response, error) {
	var u FeedBacksGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/feedbacks/%d", id)))
	return &u, resp, err
}

// GetByID 获取反馈详情
func (s *FeedBacksService) GetByID(id int) (*FeedBacksGetMsg, *req.Response, error) {
	var u FeedBacksGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/feedbacks/%d", id)))
	return &u, resp, err
}

// Assign 指派反馈
func (s *FeedBacksService) Assign(id int, assign FeedBacksAssign) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&assign).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/feedbacks/%d/assign", id)))
	return &u, resp, err
}

// CloseByID 关闭反馈
func (s *FeedBacksService) CloseByID(id int, close FeedBacksClose) (*BugGetMsg, *req.Response, error) {
	var u BugGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&close).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/feedbacks/%d/close", id)))
	return &u, resp, err
}
