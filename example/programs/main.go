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

	"github.com/davecgh/go-spew/spew"
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
	pgs, _, err := zt.Programs.List("")
	if err != nil {
		panic(err)
	}
	log.Printf("list all programs, count: %v", len(pgs.Programs))
	p1, _, err := zt.Programs.Create(zentao.ProgramsMeta{
		Name:   fmt.Sprintf("gosdk_%d%d", time.Now().Minute(), time.Now().Second()),
		Parent: 0,
		Desc:   fmt.Sprintf("gosdk_%d%d", time.Now().Minute(), time.Now().Second()),
		Begin:  time.Now().Format("2006-01-02"),
		End:    time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("program id: %v", p1.ID)
	p2, _, err := zt.Programs.GetByID(p1.ID)
	if err != nil {
		panic(err)
	}
	spew.Dump(p2)
	p3, _, err := zt.Programs.UpdateByID(p1.ID, zentao.ProgramsMeta{
		Name: fmt.Sprintf("gosdk_%d%d", time.Now().Minute(), time.Now().Second()),
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(p3)
	_, _, err = zt.Programs.DeleteByID(p1.ID)
	if err != nil {
		panic(err)
	}
}
