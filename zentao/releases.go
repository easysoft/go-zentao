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

type ReleasesService struct {
	client *Client
}

type Releases struct {
	ID          int         `json:"id"`
	Project     int         `json:"project"`
	Product     int         `json:"product"`
	Branch      int         `json:"branch"`
	Build       int         `json:"build"`
	Name        string      `json:"name"`
	Marker      string      `json:"marker"`
	Date        string      `json:"date"`
	Stories     string      `json:"stories"`
	Bugs        string      `json:"bugs"`
	Leftbugs    string      `json:"leftBugs"`
	Desc        string      `json:"desc"`
	Mailto      string      `json:"mailto"`
	Notify      interface{} `json:"notify"`
	Status      string      `json:"status"`
	Substatus   string      `json:"subStatus"`
	Deleted     bool        `json:"deleted"`
	Productname string      `json:"productName"`
	Buildid     int         `json:"buildID"`
	Buildname   string      `json:"buildName"`
	Projectname string      `json:"projectName"`
}

type ReleasesMsg struct {
	Total    int        `json:"total"`
	Releases []Releases `json:"releases"`
}

// ProjectsList 获取项目发布列表
func (s *ReleasesService) ProjectsList(id int) (*ReleasesMsg, *req.Response, error) {
	var u ReleasesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/releases", id)))
	return &u, resp, err
}

// ProductsList 获取产品发布列表
func (s *ReleasesService) ProductsList(id int) (*ReleasesMsg, *req.Response, error) {
	var u ReleasesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/releases", id)))
	return &u, resp, err
}
