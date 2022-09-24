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
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/go-zentao"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"admin",
		"jaege1ugh4ooYip7",
		zentao.WithBaseURL("https://zentao-ysicing.cloud.okteto.net"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		log.Fatal(err)
	}
	pds, _, err := zt.Products.List()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Products count: %v", len(pds.Products))
	for _, pd := range pds.Products {
		log.Printf("products id: %v", pd.ID)
		pdps, _, err := zt.ProductPlans.List(pd.ID)
		if err != nil {
			panic(err)
		}
		for _, pdp := range pdps.Plans {
			log.Printf("products plans id: %v", pdp.ID)
			spew.Dump(pdp)
		}
	}
}
