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

package main

import (
	"log"

	buildinClient "github.com/easysoft/go-zentao/v21/buildin"
	"github.com/easysoft/go-zentao/v21/zentao"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	loadSessionID()
	// basicAuthExample()
	// userExample()
}

// This example shows how to create a client with username and password.
func basicAuthExample() {
	zt, err := zentao.NewBasicAuthClient(
		"demo",
		"quickon4You",
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
	)
	if err != nil {
		log.Fatal(err)
	}
	u, _, err := zt.Users.Self()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account: %s", u.Profile.Account)
}

func userExample() {
	zt, err := zentao.NewBasicAuthClient(
		"demo",
		"quickon4You",
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
	)
	if err != nil {
		log.Fatal(err)
	}
	u, _, err := zt.Users.GetByID(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account: %s", u.Account)
}

func loadSessionID() {
	zt, err := buildinClient.NewBasicAuthClient(
		"demo",
		"quickon4You",
		buildinClient.WithDevMode(),
		buildinClient.WithDumpAll(),
	)
	if err != nil {
		log.Fatal(err)
	}
	u, _, err := zt.Login.GetSessionID()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("session id: %s", *u)
	loginStatus, loginUser, _, err := zt.Login.Login()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("login status: %v", loginStatus)
	spew.Dump(loginUser)
}
