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

package main

import (
	"log"

	"github.com/easysoft/go-zentao/v20/zentao"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"ysicing",
		"Zt@1234",
		zentao.WithBaseURL("http://192.168.94.17:8080"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		log.Fatal(err)
	}
	p1tc, _, err := zt.TestCases.ListByProducts(1)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(p1tc)
	p2tc, _, err := zt.TestCases.ListByProducts(1, map[string]string{"module": "1"})
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(p2tc)
	ts, _, err := zt.TestCases.GetResultByID("4")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(ts)
	ts1, _, err := zt.TestCases.GetResultByID("case_4")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(ts1)
}
