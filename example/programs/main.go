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
		zentao.WithBaseURL("http://127.0.0.1"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
	)
	if err != nil {
		log.Fatal(err)
	}
	pl, _, err := zt.Programs.List("")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("program count: %v", len(pl.Programs))
	cp, _, err := zt.Programs.Create(zentao.ProgramsMeta{
		Name:   fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		Parent: 0,
		Desc:   fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		Begin:  time.Now().Format("2006-01-02"),
		End:    time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("program id: %v", cp.ID)
	_, _, err = zt.Programs.DeleteByID(cp.ID)
	if err != nil {
		log.Fatal(err)
	}
	zt.Programs.GetByID(cp.ID)
}
