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

type ExecutionTasks struct {
	Page  int    `json:"page"`
	Total int    `json:"total"`
	Limit int    `json:"limit"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	TaskBody
	StoryVersion       int         `json:"storyVersion"`
	DesignVersion      int         `json:"designVersion"`
	FromBug            int         `json:"fromBug"`
	SubStatus          string      `json:"subStatus"`
	OpenedBy           UserMeta    `json:"openedBy"`
	OpenedDate         time.Time   `json:"openedDate"`
	AssignedTo         UserMeta    `json:"assignedTo"`
	AssignedDate       time.Time   `json:"assignedDate"`
	EstStarted         string      `json:"estStarted"`
	RealStarted        time.Time   `json:"realStarted"`
	FinishedBy         UserMeta    `json:"finishedBy"`
	FinishedDate       time.Time   `json:"finishedDate"`
	FinishedList       string      `json:"finishedList"`
	CanceledBy         interface{} `json:"canceledBy"`
	CanceledDate       interface{} `json:"canceledDate"`
	ClosedBy           interface{} `json:"closedBy"`
	ClosedDate         interface{} `json:"closedDate"`
	PlanDuration       int         `json:"planDuration"`
	RealDuration       int         `json:"realDuration"`
	ClosedReason       string      `json:"closedReason"`
	LastEditedBy       UserMeta    `json:"lastEditedBy"`
	LastEditedDate     time.Time   `json:"lastEditedDate"`
	ActivatedDate      string      `json:"activatedDate"`
	StoryID            interface{} `json:"storyID"`
	StoryTitle         interface{} `json:"storyTitle"`
	Product            interface{} `json:"product"`
	Branch             interface{} `json:"branch"`
	LatestStoryVersion interface{} `json:"latestStoryVersion"`
	StoryStatus        interface{} `json:"storyStatus"`
	AssignedToRealName string      `json:"assignedToRealName"`
	NeedConfirm        bool        `json:"needConfirm"`
}

type TaskCreateMeta struct {
	Module     int      `json:"module,omitempty"`
	Story      int      `json:"story,omitempty"`
	FromBug    int      `json:"fromBug,omitempty"`
	Pri        int      `json:"pri,omitempty"`
	Estimate   float64  `json:"estimate,omitempty"`
	Name       string   `json:"name"`
	Assignedto []string `json:"assignedTo"`
	Type       string   `json:"type"`
	Eststarted string   `json:"estStarted"`
	Deadline   string   `json:"deadline"`
}

type TaskBody struct {
	ID        int         `json:"id"`
	Project   int         `json:"project"`
	Parent    int         `json:"parent"`
	Execution int         `json:"execution"`
	Module    int         `json:"module"`
	Design    int         `json:"design"`
	Story     int         `json:"story"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Pri       int         `json:"pri"`
	Estimate  int         `json:"estimate"`
	Consumed  int         `json:"consumed"`
	Left      int         `json:"left"`
	Deadline  string      `json:"deadline"`
	Status    string      `json:"status"`
	Color     string      `json:"color"`
	Mailto    interface{} `json:"mailto"`
	Desc      string      `json:"desc"`
	Version   int         `json:"version"`
	Deleted   bool        `json:"deleted"`
	Progress  int         `json:"progress"`
}

type TaskCreateMsg struct {
	Task
	Children []interface{} `json:"children,omitempty"`
	Team     []interface{} `json:"team,omitempty"`
	Files    []interface{} `json:"files,omitempty"`
}

type TaskGetMsg struct {
	TaskCreateMsg
	Storyspec     string        `json:"storySpec"`
	Storyverify   string        `json:"storyVerify"`
	Storyfiles    []interface{} `json:"storyFiles"`
	Executionname string        `json:"executionName"`
	Moduletitle   string        `json:"moduleTitle"`
	Actions       []ActionMeta  `json:"actions"`
	Preandnext    struct {
		Pre  string `json:"pre"`
		Next string `json:"next"`
	} `json:"preAndNext"`
}

type TeamMeta struct {
	UserMeta
	Root     int    `json:"root"`
	Type     string `json:"type"`
	Limited  string `json:"limited"`
	Join     string `json:"join"`
	Days     int    `json:"days"`
	Hours    int    `json:"hours"`
	Estimate int    `json:"estimate"`
	Consumed int    `json:"consumed"`
	Left     int    `json:"left"`
	Order    int    `json:"order"`
	Progress int    `json:"progress"`
}

type ActionMeta struct {
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
	History    []interface{} `json:"history"`
	Desc       string        `json:"desc"`
}

type TasksService struct {
	client *Client
}

func (s *TasksService) ListByExecution(id int64) (*ExecutionTasks, *req.Response, error) {
	var et ExecutionTasks
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/tasks", id)))
	return &et, resp, err
}

// Create 创建任务
func (s *TasksService) Create(id int, build TaskCreateMeta) (*TaskCreateMsg, *req.Response, error) {
	var u TaskCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&build).
		SetSuccessResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/executions/%d/tasks", id)))
	return &u, resp, err
}

// DeleteByID 删除任务
func (s *TasksService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/tasks/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改任务
func (s *TasksService) UpdateByID(id int, exec TaskCreateMeta) (*TaskCreateMsg, *req.Response, error) {
	var u TaskCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&exec).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/tasks/%d", id)))
	return &u, resp, err
}

// GetByID 获取任务详情
func (s *TasksService) GetByID(id int) (*TaskGetMsg, *req.Response, error) {
	var u TaskGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/tasks/%d", id)))
	return &u, resp, err
}
