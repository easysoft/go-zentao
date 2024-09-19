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

type ReleasesCommon struct {
	ID           int    `json:"id"`
	Project      string `json:"project"`
	Product      int    `json:"product"`
	Branch       string `json:"branch"`
	Shadow       int    `json:"shadow"`
	Name         string `json:"name"`
	Marker       string `json:"marker"`
	Date         string `json:"date"`
	ReleasedDate string `json:"releasedDate"`
	Stories      string `json:"stories"`
	Bugs         string `json:"bugs"`
	Leftbugs     string `json:"leftBugs"`
	Desc         string `json:"desc"`
	Mailto       any    `json:"mailto"`
	Notify       any    `json:"notify"`
	Status       string `json:"status"`
	Substatus    string `json:"subStatus"`
	CreatedBy    any    `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	Deleted      bool   `json:"deleted"`
	ProductName  string `json:"productName"`
	ProductType  string `json:"productType"`
}

type BuildInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Project     int    `json:"project"`
	Execution   int    `json:"execution"`
	ProjectName string `json:"projectName,omitempty"`
	Link        string `json:"link,omitempty"`
	Product     int    `json:"product,omitempty"`
	Branch      string `json:"branch,omitempty"`
	ScmPath     string `json:"scmPath,omitempty"`
	FilePath    string `json:"filePath,omitempty"`
	BranchName  string `json:"branchName,omitempty"`
}

type ProjectRelease struct {
	ReleasesCommon
	BranchName string               `json:"branchName,omitempty"`
	BuildInfos map[string]BuildInfo `json:"buildInfos,omitempty"`
}

type ProjectReleasesMsg struct {
	Total    int              `json:"total"`
	Limit    int              `json:"limit"`
	Page     int              `json:"page"`
	Releases []ProjectRelease `json:"releases"`
}

type ProductRelease struct {
	ReleasesCommon
	BranchName  string      `json:"branchName,omitempty"`
	ProjectName string      `json:"projectName,omitempty"`
	Actions     []string    `json:"actions,omitempty"`
	Rowspan     int         `json:"rowspan,omitempty"`
	Build       BuildInfo   `json:"build,omitempty"`
	Builds      []BuildInfo `json:"builds,omitempty"`
}

type ProductReleasesMsg struct {
	Total    int              `json:"total"`
	Limit    int              `json:"limit"`
	Page     int              `json:"page"`
	Releases []ProductRelease `json:"releases"`
}

// ProjectsList 获取项目发布列表
func (s *ReleasesService) ProjectsList(id int) (*ProjectReleasesMsg, *req.Response, error) {
	var u ProjectReleasesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/projects/%d/releases", id)))
	return &u, resp, err
}

// ProductsList 获取产品发布列表
func (s *ReleasesService) ProductsList(id int) (*ProductReleasesMsg, *req.Response, error) {
	var u ProductReleasesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/releases", id)))
	return &u, resp, err
}
