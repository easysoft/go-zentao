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

// UserType 用户类型
type UserType string

var (
	InSideUser  UserType = "inside"  // 内部用户
	OutSideUser UserType = "outside" // 外部用户
)

type UserGender string

var (
	ManGender   UserGender = "m" // 男
	WomanGender UserGender = "f" // 女
)

// CustomResp 通用Resp
type CustomResp struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
