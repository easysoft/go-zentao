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

type ProjectExecutions struct {
	Page       int         `json:"page"`
	Total      int         `json:"total"`
	Limit      int         `json:"limit"`
	Executions []Execution `json:"executions"`
}

type Execution struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Project    int       `json:"project"`
	Code       string    `json:"code"`
	Type       string    `json:"type"`
	Parent     int       `json:"parent"`
	Begin      string    `json:"begin"`
	End        string    `json:"end"`
	Status     string    `json:"status"`
	OpenedBy   string    `json:"openedBy"`
	OpenedDate time.Time `json:"openedDate"`
	Progress   int       `json:"progress"`
}

type ExecutionsService struct {
	client *Client
}

func (s *ExecutionsService) ListByProject(id int64) (*ProjectExecutions, *req.Response, error) {
	var pe ProjectExecutions
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&pe).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/executions", id)))
	return &pe, resp, err
}
