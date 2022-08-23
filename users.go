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

type UsersService struct {
	client *Client
}

type User struct {
	Profile Profile `json:"profile"`
}

type UserMeta struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Avatar   string `json:"avatar"`
	Realname string `json:"realname"`
}

type Profile struct {
	UserMeta
	Company       int         `json:"company"`
	Type          string      `json:"type"`
	Dept          int         `json:"dept"`
	Role          Role        `json:"role"`
	Pinyin        string      `json:"pinyin"`
	Nickname      string      `json:"nickname"`
	Commiter      string      `json:"commiter"`
	Birthday      interface{} `json:"birthday"`
	Gender        string      `json:"gender"`
	Email         string      `json:"email"`
	Skype         string      `json:"skype"`
	Qq            string      `json:"qq"`
	Mobile        string      `json:"mobile"`
	Phone         string      `json:"phone"`
	Weixin        string      `json:"weixin"`
	Dingding      string      `json:"dingding"`
	Slack         string      `json:"slack"`
	Whatsapp      string      `json:"whatsapp"`
	Address       string      `json:"address"`
	Zipcode       string      `json:"zipcode"`
	Nature        string      `json:"nature"`
	Analysis      string      `json:"analysis"`
	Strategy      string      `json:"strategy"`
	Join          interface{} `json:"join"`
	Visits        int         `json:"visits"`
	Visions       string      `json:"visions"`
	IP            string      `json:"ip"`
	Last          time.Time   `json:"last"`
	Fails         int         `json:"fails"`
	Locked        interface{} `json:"locked"`
	Feedback      string      `json:"feedback"`
	Ranzhi        string      `json:"ranzhi"`
	Ldap          string      `json:"ldap"`
	Score         int         `json:"score"`
	ScoreLevel    int         `json:"scoreLevel"`
	Deleted       string      `json:"deleted"`
	ClientStatus  string      `json:"clientStatus"`
	ClientLang    string      `json:"clientLang"`
	Admin         bool        `json:"admin"`
	SuperReviewer bool        `json:"superReviewer"`
}

type Role struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (s *UsersService) Self() (*User, *req.Response, error) {
	var u User
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL("/user"))
	return &u, resp, err
}

func (s *UsersService) GetByID(id int64) (*User, *req.Response, error) {
	var u User
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/users/%d", id)))
	return &u, resp, err
}
