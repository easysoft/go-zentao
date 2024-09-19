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
		panic(err)
	}
	pds, _, err := zt.ProductPlans.List(1)
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 plan count: %v", len(pds.Plans))
	p1, _, err := zt.ProductPlans.Create(1, zentao.ProductPlanMeta{
		Title: "gosdk-" + time.Now().String(),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 create productplan id: %v", p1.ID)
	p2, _, err := zt.ProductPlans.GetByID(p1.ID)
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 get productplan id: %v", p2.ID)
	spew.Dump(p2)
	p3, _, err := zt.ProductPlans.UpdateByID(p2.ID, zentao.ProductPlanMeta{
		Title: "gosdk-" + time.Now().String(),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 update productplan id: %v", p3.ID)
	spew.Dump(p3)
	p4, _, err := zt.ProductPlans.LinkStories(p3.ID, zentao.PlansStoriesIDs{
		Stories: []int{1, 2, 3},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 link stories to productplan, stories: %v", len(p4.Stories))
	p5, _, err := zt.ProductPlans.UnLinkStories(p3.ID, zentao.PlansStoriesIDs{
		Stories: []int{1, 2},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 unlink stories to productplan, exist stories: %v", len(p5.Stories))
	p6, _, err := zt.ProductPlans.LinkBugs(p3.ID, zentao.PlansBugIDs{
		Bugs: []int{1, 2, 3},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 link bugs to productplan, bugs: %v", len(p6.Bugs))
	p7, _, err := zt.ProductPlans.UnLinkBugs(p3.ID, zentao.PlansBugIDs{
		Bugs: []int{1},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 link bugs to productplan, exist bugs: %v", len(p7.Bugs))
	p8, _, err := zt.ProductPlans.CreateSub(1, p3.ID, zentao.ProductPlanMeta{
		Title: "sub-gosdk-" + time.Now().String(),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 create sub productplan id: %v", p8.ID)
	_, _, err = zt.ProductPlans.DeleteByID(p8.ID)
	if err != nil {
		panic(err)
	}
}
