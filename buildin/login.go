//
//  Copyright 2024, zentao
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

package buildin

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
)

type LoginService struct {
	client *Client
}

type BuiltINData struct {
	Status string `json:"status"`
	Data   string `json:"data"`
	Md5    string `json:"md5"`
}

type SessionData struct {
	Title       string      `json:"title,omitempty"`
	SessionName string      `json:"sessionName,omitempty"`
	SessionID   string      `json:"sessionID"`
	Rand        int         `json:"rand,omitempty"`
	Pager       interface{} `json:"pager,omitempty"`
}

type LoginRespData struct {
	Status string        `json:"status"`
	User   LoginUserData `json:"user"`
}

type LoginUserData struct {
	ID             string          `json:"id"`
	Company        string          `json:"company"`
	Type           string          `json:"type"`
	Dept           string          `json:"dept"`
	Account        string          `json:"account"`
	Role           string          `json:"role"`
	Realname       string          `json:"realname"`
	Pinyin         string          `json:"pinyin"`
	Nickname       string          `json:"nickname"`
	Commiter       string          `json:"commiter"`
	Avatar         string          `json:"avatar"`
	Birthday       string          `json:"birthday"`
	Gender         string          `json:"gender"`
	Email          string          `json:"email"`
	Skype          string          `json:"skype"`
	Qq             string          `json:"qq"`
	Mobile         string          `json:"mobile"`
	Phone          string          `json:"phone"`
	Weixin         string          `json:"weixin"`
	Dingding       string          `json:"dingding"`
	Slack          string          `json:"slack"`
	Whatsapp       string          `json:"whatsapp"`
	Address        string          `json:"address"`
	Zipcode        string          `json:"zipcode"`
	Nature         string          `json:"nature"`
	Analysis       string          `json:"analysis"`
	Strategy       string          `json:"strategy"`
	Join           string          `json:"join"`
	Visits         string          `json:"visits"`
	Visions        string          `json:"visions"`
	IP             string          `json:"ip"`
	Last           string          `json:"last"`
	Fails          string          `json:"fails"`
	Locked         string          `json:"locked"`
	Feedback       string          `json:"feedback"`
	Ranzhi         string          `json:"ranzhi"`
	Ldap           string          `json:"ldap"`
	Score          string          `json:"score"`
	ScoreLevel     string          `json:"scoreLevel"`
	ResetToken     string          `json:"resetToken"`
	ClientStatus   string          `json:"clientStatus"`
	ClientLang     string          `json:"clientLang"`
	LastTime       string          `json:"lastTime"`
	Admin          bool            `json:"admin"`
	ModifyPassword bool            `json:"modifyPassword"`
	Rights         LoginUserRights `json:"rights"`
	Groups         struct {
		Num13 string `json:"13"`
	} `json:"groups"`
	View  LoginUserView `json:"view"`
	Token string        `json:"token"`
}

type LoginUserRights struct {
	Rights struct {
		Program struct {
			Activate                bool `json:"activate"`
			Batchunlinkstakeholders bool `json:"batchunlinkstakeholders"`
			Browse                  bool `json:"browse"`
			Close                   bool `json:"close"`
			Create                  bool `json:"create"`
			Createstakeholder       bool `json:"createstakeholder"`
			Delete                  bool `json:"delete"`
			Edit                    bool `json:"edit"`
			Kanban                  bool `json:"kanban"`
			Product                 bool `json:"product"`
			Project                 bool `json:"project"`
			Stakeholder             bool `json:"stakeholder"`
			Start                   bool `json:"start"`
			Suspend                 bool `json:"suspend"`
			Unbindwhitelist         bool `json:"unbindwhitelist"`
			Unlinkstakeholder       bool `json:"unlinkstakeholder"`
			Updateorder             bool `json:"updateorder"`
			View                    bool `json:"view"`
		} `json:"program"`
		Projectplan struct {
			Browse bool `json:"browse"`
		} `json:"projectplan"`
		Traincourse struct {
			Browse         bool `json:"browse"`
			Practice       bool `json:"practice"`
			Practicebrowse bool `json:"practicebrowse"`
			Practiceview   bool `json:"practiceview"`
			Updatepractice bool `json:"updatepractice"`
			Viewchapter    bool `json:"viewchapter"`
			Viewcourse     bool `json:"viewcourse"`
		} `json:"traincourse"`
		Index struct {
			Index int `json:"index"`
		} `json:"index"`
		My struct {
			Index int `json:"index"`
		} `json:"my"`
	} `json:"rights"`
	Acls       []any  `json:"acls"`
	Projects   string `json:"projects"`
	Programs   string `json:"programs"`
	Products   string `json:"products"`
	Executions string `json:"executions"`
}

type LoginUserView struct {
	Account  string `json:"account"`
	Programs string `json:"programs"`
	Products string `json:"products"`
	Projects string `json:"projects"`
	Sprints  string `json:"sprints"`
}

func (s *LoginService) GetSessionID() (*string, *req.Response, error) {
	var data BuiltINData
	var resp *req.Response
	var err error
	if s.client.zentaomode == ZentaoGetMode {
		resp, err = s.client.client.R().
			SetSuccessResult(&data).
			Get(s.client.RequestURL("/index.php?m=api&f=getSessionID&t=json"))
	} else {
		resp, err = s.client.client.R().
			SetSuccessResult(&data).
			Get(s.client.RequestURL("/api-getsessionid.json"))
	}
	if err != nil {
		return nil, resp, err
	}

	if data.Status != "success" {
		return nil, resp, fmt.Errorf("get session id failed status not success")
	}
	var result SessionData
	err = json.Unmarshal([]byte(data.Data), &result)
	if err != nil {
		return nil, resp, fmt.Errorf("unmarshal session data failed: %w", err)
	}
	return &result.SessionID, resp, err
}

func (s *LoginService) Login() (bool, *LoginRespData, *req.Response, error) {
	var data LoginRespData
	var resp *req.Response
	var err error
	if s.client.zentaomode == ZentaoGetMode {
		resp, err = s.client.client.R().
			SetSuccessResult(&data).
			Get(s.client.RequestURLFmt("/index.php?m=user&f=login&t=json&account=%s&password=%s&zentaosid=%s", s.client.username, s.client.password, s.client.zentaosid))
	} else {
		resp, err = s.client.client.R().
			SetSuccessResult(&data).
			Get(s.client.RequestURLFmt("/user-login.json?account=%s&password=%s&zentaosid=%s", s.client.username, s.client.password, s.client.zentaosid))
	}
	if err != nil {
		return false, &data, resp, err
	}
	if data.Status != "success" {
		return false, &data, resp, fmt.Errorf("get session id failed status not success")
	}
	return true, &data, resp, nil
}
