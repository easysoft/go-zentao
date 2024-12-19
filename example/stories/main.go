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
	p1, _, err := zt.Stories.ProductsList(1)
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 stories count: %v", len(p1.Stories))
	p2, _, err := zt.Stories.ProjectsList(2)
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 stories count: %v", len(p2.Stories))
	createMsg := zentao.StoriesCreateMeta{
		Product: 1,
	}
	createMsg.Title = "test story"
	createMsg.Pri = 1
	createMsg.Category = zentao.CategoryFeature
	createMsg.Spec = "test spec"
	// 可选
	createMsg.Verify = "test verify"
	p3, _, err := zt.Stories.Create(createMsg)
	if err != nil {
		panic(err)
	}
	log.Printf("created story id: %v", p3.ID)
	p4, _, err := zt.Stories.UpdateByID(p3.ID, zentao.StoriesMeta{
		Title: "test story updated",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("updated story id: %v", p4.ID)
	uptateMsg := zentao.StoriesUpdateFieldMeta{}
	uptateMsg.Category = zentao.CategoryOther
	uptateMsg.Reviewer = []string{"dev1"}
	p5, _, err := zt.Stories.UpdateFieldByID(p4.ID, uptateMsg)
	if err != nil {
		panic(err)
	}
	log.Printf("updated story filed id: %v", p5.ID)
	p6, _, err := zt.Stories.GetByID(p4.ID)
	if err != nil {
		panic(err)
	}
	spew.Dump(p6)
	p7, _, err := zt.Stories.CloseByID(p6.ID, zentao.StoriesClose{
		Closedreason: zentao.CloseReasonBydesign,
		Comment:      "test close",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("close story id: %v", p7.ID)
	p8, _, err := zt.Stories.ActiveByID(p7.ID, zentao.StoriesActive{
		Comment:    "test active",
		AssignedTo: "dev1",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("active story id: %v", p8.ID)
	p9, _, err := zt.Stories.AssignByID(p8.ID, zentao.StoriesActive{
		Comment:    "test assign",
		AssignedTo: "dev2",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("assign story id: %v", p9.ID)
	// p10, _, err := zt.Stories.GetEstimateByID(p9.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("assign story id: %v", p10.ID)
	// p10, _, err := zt.Stories.UpdateEstimateByID(p9.ID, zentao.StoriesEstimate{})
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("assign story id: %v", p10.ID)
	_, _, err = zt.Stories.RecallByID(p9.ID)
	if err != nil {
		panic(err)
	}
	log.Printf("recall story id: %v", p9.ID)
	p11, _, err := zt.Stories.ReviewByID(p9.ID, zentao.StoriesReview{
		Result:   zentao.ReviewResultPass,
		Estimate: 1.1,
		Comment:  "test review",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("review story id: %v", p11.ID)
	_, _, err = zt.Stories.DeleteByID(p9.ID)
	if err != nil {
		panic(err)
	}
}
