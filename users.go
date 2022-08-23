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

	"github.com/imroc/req/v3"
)

type UsersService struct {
	client *Client
}

type UserMeta struct {
	ID       int         `json:"id"`
	Dept     int         `json:"dept"`
	Account  string      `json:"account"`
	Realname string      `json:"realname"`
	Role     interface{} `json:"role"`
	Email    string      `json:"email"`
}

type UserCreateMeta struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Realname string `json:"realname,omitempty"`
	Visions  string `json:"visions,omitempty"`
}

type SelfMsg struct {
	Profile UserProfile `json:"profile"`
}

type UserProfile struct {
	UserMeta
	UserSocial
	Admin bool `json:"admin"`
}

type UserGetMsg struct {
	UserMeta
	UserSocial
}

type UserCreateMsg struct {
	UserMeta
	Pinyin string `json:"pinyin"`
	UserSocial
	Visions      string `json:"visions"`
	Feedback     string `json:"feedback"`
	Ldap         string `json:"ldap"`
	Clientstatus string `json:"clientStatus"`
	Clientlang   string `json:"clientLang"`
}

type UserSocial struct {
	Score      int         `json:"score"`
	Scorelevel int         `json:"scoreLevel"`
	Deleted    interface{} `json:"deleted"`
	Ranzhi     string      `json:"ranzhi"`
	Last       interface{} `json:"last"`
	Fails      int         `json:"fails"`
	Locked     interface{} `json:"locked"`
	Join       interface{} `json:"join"`
	Visits     int         `json:"visits"`
	IP         string      `json:"ip"`
	Company    int         `json:"company"`
	Type       string      `json:"type"`
	Nickname   string      `json:"nickname"`
	Commiter   string      `json:"commiter"`
	Avatar     string      `json:"avatar"`
	Birthday   interface{} `json:"birthday"`
	Address    string      `json:"address"`
	Zipcode    string      `json:"zipcode"`
	Nature     string      `json:"nature"`
	Analysis   string      `json:"analysis"`
	Strategy   string      `json:"strategy"`
	Gender     string      `json:"gender"`
	Skype      string      `json:"skype"`
	Qq         string      `json:"qq"`
	Mobile     string      `json:"mobile"`
	Phone      string      `json:"phone"`
	Weixin     string      `json:"weixin"`
	Dingding   string      `json:"dingding"`
	Slack      string      `json:"slack"`
	Whatsapp   string      `json:"whatsapp"`
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

type CustomResp struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Self 获取我的个人信息
func (s *UsersService) Self() (*SelfMsg, *req.Response, error) {
	var u SelfMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL("/user"))
	return &u, resp, err
}

// GetByID 获取用户信息
func (s *UsersService) GetByID(id int) (*UserGetMsg, *req.Response, error) {
	var u UserGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

// DeleteByID 删除用户
func (s *UsersService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

// Create 创建用户
func (s *UsersService) Create(user UserCreateMeta) (*UserCreateMsg, *req.Response, error) {
	var u UserCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetResult(&u).
		Post(s.client.RequestURL("/users"))
	return &u, resp, err
}

// UpdateByID 更新用户
func (s *UsersService) UpdateByID(id int, user UserUpdateMeta) (*UserGetMsg, *req.Response, error) {
	var u UserGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&user).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}

// List 获取用户列表
func (s *UsersService) List(limit, page string) (*UserList, *req.Response, error) {
	var u UserList
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetQueryParams(map[string]string{
			"page":  page,
			"limit": limit,
		}).
		SetResult(&u).
		Get(s.client.RequestURL("/users"))
	return &u, resp, err
}
