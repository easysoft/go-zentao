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

	"github.com/easysoft/go-zentao/v20/zentao"
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
		log.Fatal(err)
	}
	pds, _, err := zt.Products.List()
	if err != nil {
		panic(err)
	}
	id := 0
	if len(pds.Products) == 0 {
		pd, _, err := zt.Products.Create(zentao.ProductsMeta{
			Name: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
			Code: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		})
		if err != nil {
			panic(err)
		}
		id = pd.ID
	} else {
		id = pds.Products[0].ID
	}
	pjs, _, err := zt.Projects.List("100", "1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Projects count: %v", len(pjs.Projects))
	pj, _, err := zt.Projects.Create(zentao.ProjectsCreateMeta{
		Name:     fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		Code:     fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		Begin:    time.Now().Format("2006-01-02"),
		End:      time.Now().Add(72 * time.Hour).Format("2006-01-02"),
		Products: []int{id},
		Parent:   0,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Create Projects: %v", pj.ID)
}
