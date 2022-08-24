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

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ysicing/go-zentao"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"admin",
		"jaege1ugh4ooYip7",
		zentao.WithBaseURL("http://172.77.77.12"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		log.Fatal(err)
	}
	u, _, err := zt.Users.Self()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account: %s", u.Profile.Account)
	cu, _, err := zt.Users.Create(zentao.UserCreateMeta{
		Account:  fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()), // 不超过30位且字母、数字或下划线
		Password: "demo11111111x.x",
		Realname: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
	})
	if err != nil {
		log.Fatal(err)
	}
	if gu, _, err := zt.Users.GetByID(cu.ID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("account: %v, name: %s", gu.ID, gu.Realname)
	}
	time.Sleep(time.Second * 2)
	zt.Users.UpdateByID(cu.ID, zentao.UserUpdateMeta{Realname: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second())})
	if gu, _, err := zt.Users.GetByID(cu.ID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("account: %v, name: %s", gu.ID, gu.Realname)
	}
	zt.Users.DeleteByID(cu.ID)
	ul, _, err := zt.Users.List("10", "1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user %v", ul.Total)
}
