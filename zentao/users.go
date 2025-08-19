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

type UsersService struct {
	client *Client
}

type UserMeta struct {
	ID       int         `json:"id"`
	Dept     int         `json:"dept,omitempty"`
	Account  string      `json:"account"`
	Realname string      `json:"realname"`
	Role     interface{} `json:"role,omitempty"` // Role这个字段比较特殊
	Email    string      `json:"email,omitempty"`
	Avatar   string      `json:"avatar,omitempty"`
}

type UserCreateMeta struct {
	Account  string     `json:"account"`
	Password string     `json:"password"`
	Realname string     `json:"realname,omitempty"`
	Visions  string     `json:"visions,omitempty"` // 研发综合界面 rnd, 运营管理界面 lite) 默认是rnd
	Email    string     `json:"email,omitempty"`
	Gender   UserGender `json:"gender"`
	Type     UserType   `json:"type,omitempty"`  // 用户类型, 为空默认是内部用户
	Role     string     `json:"role,omitempty"`  // 职位
	Group    string     `json:"group,omitempty"` // 权限分组, 示例3,4
	// 内部用户参数
	Dept int `json:"dept,omitempty"`
}

type SelfMsg struct {
	Profile UserProfile `json:"profile"`
}

type UserProfile struct {
	UserMeta
	UserSocial
	Admin        bool   `json:"admin,omitempty"`
	Pinyin       string `json:"pinyin,omitempty"`
	Visions      string `json:"visions,omitempty"`
	Feedback     string `json:"feedback,omitempty"`
	Ldap         string `json:"ldap,omitempty"`
	Clientstatus string `json:"clientStatus,omitempty"`
	Clientlang   string `json:"clientLang,omitempty"`
}

type UserSocial struct {
	Score      int        `json:"score"`
	Scorelevel int        `json:"scoreLevel"`
	Deleted    string     `json:"deleted"`
	Ranzhi     string     `json:"ranzhi"`
	Last       string     `json:"last"`
	Fails      int        `json:"fails"`
	Locked     string     `json:"locked"`
	Join       string     `json:"join"`
	Visits     int        `json:"visits"`
	IP         string     `json:"ip"`
	Company    int        `json:"company"`
	Type       UserType   `json:"type"`
	Nickname   string     `json:"nickname"`
	Commiter   string     `json:"commiter"`
	Birthday   string     `json:"birthday"`
	Address    string     `json:"address"`
	Zipcode    string     `json:"zipcode"`
	Nature     string     `json:"nature"`
	Analysis   string     `json:"analysis"`
	Strategy   string     `json:"strategy"`
	Gender     UserGender `json:"gender"`
	Skype      string     `json:"skype"`
	Qq         string     `json:"qq"`
	Mobile     string     `json:"mobile"`
	Phone      string     `json:"phone"`
	Weixin     string     `json:"weixin"`
	Dingding   string     `json:"dingding"`
	Slack      string     `json:"slack"`
	Whatsapp   string     `json:"whatsapp"`
}

type UserUpdateMeta struct {
	Dept     int    `json:"dept,omitempty"`
	Role     string `json:"role,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Realname string `json:"realname,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type UserList struct {
	Page  int        `json:"page"`
	Total int        `json:"total"`
	Limit int        `json:"limit"`
	Users []UserMeta `json:"users"`
}

type DepartmentsMsg []DepartmentMeta

type DepartmentMeta struct {
	DepartmentCustomMeta
	ManagerName string           `json:"managerName"`
	URL         string           `json:"url"`
	Key         string           `json:"key"`
	Text        string           `json:"text"`
	Items       []DepartmentMeta `json:"items,omitempty"` // 子部门
}

type DepartmentCustomMeta struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Parent   int    `json:"parent"`
	Path     string `json:"path"`
	Grade    int    `json:"grade"`
	Order    int    `json:"order"`
	Position string `json:"position"`
	Function string `json:"function"`
	Manager  string `json:"manager"`
}

type GroupMeta struct {
	ID        int    `json:"id"`
	Project   int    `json:"project"`
	Vision    string `json:"vision"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Desc      string `json:"desc"`
	ACL       any    `json:"acl"`
	Developer string `json:"developer"`
	Accounts  any    `json:"accounts,omitempty"` // 用户列表 map[string]string
	Privs     any    `json:"privs"`
}

type GroupsMsg []GroupMeta

// Self 获取我的个人信息
// GET /user
func (s *UsersService) Self() (*SelfMsg, *req.Response, error) {
	var u SelfMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL("/user"))
	return &u, resp, err
}

// GetByID 获取用户信息
func (s *UsersService) GetByID(id int) (*UserProfile, *req.Response, error) {
	var u UserProfile
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

// DeleteByID 删除用户
func (s *UsersService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

// Create 创建用户
func (s *UsersService) Create(user UserCreateMeta) (*UserProfile, *req.Response, error) {
	var u UserProfile
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetSuccessResult(&u).
		Post(s.client.RequestURL("/users"))
	return &u, resp, err
}

// UpdateByID 更新用户
func (s *UsersService) UpdateByID(id int, user UserUpdateMeta) (*UserProfile, *req.Response, error) {
	var u UserProfile
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetSuccessResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

func (s *UsersService) List(limit, page string) (*UserList, *req.Response, error) {
	var u UserList
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"page":  page,
			"limit": limit,
		}).
		SetSuccessResult(&u).
		Get(s.client.RequestURL("/users"))
	return &u, resp, err
}

func (s *UsersService) ListAllDepartments() (*DepartmentsMsg, *req.Response, error) {
	var result DepartmentsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&result).
		Get(s.client.RequestURL("/departments"))
	return &result, resp, err
}

func (s *UsersService) GetDepartmentByID(deptID int) (*DepartmentCustomMeta, *req.Response, error) {
	var result DepartmentCustomMeta
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&result).
		Get(s.client.RequestURL(fmt.Sprintf("/departments/%d", deptID)))
	return &result, resp, err
}

func (s *UsersService) ListAllGroups() (*GroupsMsg, *req.Response, error) {
	var result GroupsMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetSuccessResult(&result).
		Get(s.client.RequestURL("/groups"))
	return &result, resp, err
}
