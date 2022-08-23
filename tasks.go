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

type ExecutionTasks struct {
	Page  int    `json:"page"`
	Total int    `json:"total"`
	Limit int    `json:"limit"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID                 int           `json:"id"`
	Project            int           `json:"project"`
	Parent             int           `json:"parent"`
	Execution          int           `json:"execution"`
	Module             int           `json:"module"`
	Design             int           `json:"design"`
	Story              int           `json:"story"`
	StoryVersion       int           `json:"storyVersion"`
	DesignVersion      int           `json:"designVersion"`
	FromBug            int           `json:"fromBug"`
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	Pri                int           `json:"pri"`
	Estimate           int           `json:"estimate"`
	Consumed           int           `json:"consumed"`
	Left               int           `json:"left"`
	Deadline           interface{}   `json:"deadline"`
	Status             string        `json:"status"`
	SubStatus          string        `json:"subStatus"`
	Color              string        `json:"color"`
	Mailto             []interface{} `json:"mailto"`
	Desc               string        `json:"desc"`
	Version            int           `json:"version"`
	OpenedBy           UserMeta      `json:"openedBy"`
	OpenedDate         time.Time     `json:"openedDate"`
	AssignedTo         UserMeta      `json:"assignedTo"`
	AssignedDate       time.Time     `json:"assignedDate"`
	EstStarted         string        `json:"estStarted"`
	RealStarted        time.Time     `json:"realStarted"`
	FinishedBy         UserMeta      `json:"finishedBy"`
	FinishedDate       time.Time     `json:"finishedDate"`
	FinishedList       string        `json:"finishedList"`
	CanceledBy         interface{}   `json:"canceledBy"`
	CanceledDate       interface{}   `json:"canceledDate"`
	ClosedBy           interface{}   `json:"closedBy"`
	ClosedDate         interface{}   `json:"closedDate"`
	PlanDuration       int           `json:"planDuration"`
	RealDuration       int           `json:"realDuration"`
	ClosedReason       string        `json:"closedReason"`
	LastEditedBy       UserMeta      `json:"lastEditedBy"`
	LastEditedDate     time.Time     `json:"lastEditedDate"`
	ActivatedDate      string        `json:"activatedDate"`
	Deleted            bool          `json:"deleted"`
	StoryID            interface{}   `json:"storyID"`
	StoryTitle         interface{}   `json:"storyTitle"`
	Product            interface{}   `json:"product"`
	Branch             interface{}   `json:"branch"`
	LatestStoryVersion interface{}   `json:"latestStoryVersion"`
	StoryStatus        interface{}   `json:"storyStatus"`
	AssignedToRealName string        `json:"assignedToRealName"`
	NeedConfirm        bool          `json:"needConfirm"`
	Progress           int           `json:"progress"`
}

type TasksService struct {
	client *Client
}

func (s *TasksService) ListByExecution(id int64) (*ExecutionTasks, *req.Response, error) {
	var et ExecutionTasks
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&et).
		Get(s.client.RequestURL(fmt.Sprintf("/executions/%d/tasks", id)))
	return &et, resp, err
}
