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
	"fmt"
	"log"
	"time"

	"github.com/easysoft/go-zentao/v21/zentao"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"admin",
		"jaege1ugh4ooYip7",
		zentao.WithBaseURL("http://127.0.0.1"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		panic(err)
	}
	p1, _, err := zt.Projects.List("100", "1")
	if err != nil {
		panic(err)
	}
	log.Printf("projects count: %v", len(p1.Projects))
	p2, _, err := zt.Projects.Create(zentao.ProjectsCreateMeta{
		Name:     fmt.Sprintf("gosdk_p1_%s", time.Now().Format("20060102150405")),
		Code:     fmt.Sprintf("gosdk_p1_%s", time.Now().Format("20060102150405")),
		Begin:    time.Now().Format("2006-01-02"),
		End:      time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
		Products: []int{1},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("created project: %v", p2.ID)
	p3, _, err := zt.Projects.GetByID(p2.ID)
	if err != nil {
		panic(err)
	}
	log.Printf("get project: %v", p3.ID)
	p4, _, err := zt.Projects.UpdateByID(p2.ID, zentao.ProjectsUpdateMeta{
		Name: fmt.Sprintf("gosdk_p4_%s", time.Now().Format("20060102150405")),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("updated project: %v", p4.ID)
	_, _, err = zt.Projects.DeleteByID(p2.ID)
	if err != nil {
		panic(err)
	}
}
