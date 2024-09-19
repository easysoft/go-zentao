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
		panic(err)
	}
	pl, _, err := zt.Products.List()
	if err != nil {
		panic(err)
	}
	log.Printf("Products count: %v", len(pl.Products))
	p1, _, err := zt.Products.Create(zentao.ProductsMeta{
		Name: "gosdk_2",
		Code: "gosdk_2",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product id: %v", p1.ID)
	p2, _, err := zt.Products.GetByID(p1.ID)
	if err != nil {
		panic(err)
	}
	spew.Dump(p2)
	p3, _, err := zt.Products.UpdateByID(p1.ID, zentao.ProductsMeta{
		Name: fmt.Sprintf("gosdk_22"),
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(p3)
	_, _, err = zt.Products.DeleteByID(p1.ID)
	if err != nil {
		panic(err)
	}
}
