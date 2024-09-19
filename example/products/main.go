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
		log.Fatal(err)
	}
	pl, _, err := zt.Products.List()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Products count: %v", len(pl.Products))
	cp, _, err := zt.Products.Create(zentao.ProductsMeta{
		Name: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
		Code: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("program id: %v", cp.ID)
	getmsg, _, err := zt.Products.GetByID(cp.ID)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(getmsg)
	// _, _, err = zt.Products.DeleteByID(cp.ID)
	// _, _, err = zt.ProductPlans.Create(cp.ID, zentao.ProductPlanMeta{
	// 	Title: fmt.Sprintf("abc%d%d", time.Now().Minute(), time.Now().Second()),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zt.ProductPlans.List(cp.ID)
	// zt.ProductPlans.GetByID(5)
}
