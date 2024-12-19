//
//  Copyright 2024, easysoft
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
	p1, _, err := zt.Bugs.ListByProducts(1, zentao.ListOptions{
		Page:  1,
		Limit: 1,
		// Status: "all", // 列出所有
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 bugs count: %v", len(p1.Bugs))
	p2, _, err := zt.Bugs.ListByProjects(2, zentao.ListOptions{Page: 1, Limit: 1})
	if err != nil {
		panic(err)
	}
	log.Printf("project 1 bugs count: %v", len(p2.Bugs))
	p3, _, err := zt.Bugs.ListByExecutions(1, zentao.ListOptions{Page: 1, Limit: 1})
	if err != nil {
		panic(err)
	}
	log.Printf("execution 1 bugs count: %v", len(p3.Bugs))
	p4, _, err := zt.Bugs.GetByID(p3.Bugs[0].ID)
	if err != nil {
		panic(err)
	}
	log.Printf("bug 1: %v", p4.Title)
	p5, _, err := zt.Bugs.Create(1, zentao.BugCreateMeta{
		BugMeta: zentao.BugMeta{
			Title:      fmt.Sprintf("go sdk bug %v", p1.Total+1),
			Module:     1,
			Type:       zentao.CodeErrorBugType,
			Severity:   1,
			Pri:        1,
			Project:    2,
			Story:      1,
			Execution:  3,
			Task:       6,
			Steps:      "<p>[步骤]</p><p></p><p>[结果]</p><p></p><p>[期望]</p>",
			AssignedTo: "productManager",
			FeedbackBy: "zhangsan",
		},
		AllBuilds:   "on",
		Openedbuild: []string{"trunk"}, // 主干
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 create bug: %v", p5.Title)
	p51, _, err := zt.Bugs.UpdateByID(p5.ID, zentao.BugUpdateMeta{
		Title: fmt.Sprintf("fake go sdk bug %v", p5.ID),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 update bug: %v", p51.Title)
	p6, _, err := zt.Bugs.AssignByID(p5.ID, zentao.AssignBug{
		AssignedTo: "admin",
		Comment:    "go sdk assign bug",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 assign bug: %v", p6.Title)
	p7, _, err := zt.Bugs.ConfirmByID(p5.ID, zentao.ConfirmBug{
		AssignBug: zentao.AssignBug{
			AssignedTo: "admin",
			Comment:    "go sdk confirm bug",
		},
		Pri:    1,
		Type:   zentao.CodeErrorBugType,
		Status: zentao.ActiveBugStatus,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 confirm bug: %v", p7.Title)
	p8, _, err := zt.Bugs.ResolveByID(p5.ID, zentao.ResolveBug{
		Resolution: zentao.BugExternal,
		Comment:    "go sdk resolve bug",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 resolve bug: %v", p8.Title)
	p9, _, err := zt.Bugs.ActiveByID(p5.ID, zentao.ActiveBug{
		Comment:     "go sdk active bug",
		AssignedTo:  "admin",
		OpenedBuild: []string{"trunk", "2"},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 active bug: %v", p9.Title)
	_, _, err = zt.Bugs.ResolveByID(p5.ID, zentao.ResolveBug{
		Resolution: zentao.BugExternal,
		Comment:    "go sdk resolve2 bug",
	})
	if err != nil {
		panic(err)
	}
	p10, _, err := zt.Bugs.CloseByID(p5.ID, zentao.CommentBug{
		Comment: "go sdk close bug",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 close bug: %v", p10.Title)
	// p11, _, err := zt.Bugs.ListByProducts(1, zentao.ListOptions{
	// 	Page:   1,
	// 	Limit:  100,
	// 	Status: "all", // 列出所有
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("product 1 all bugs count: %v", len(p11.Bugs))
	// p12, _, err := zt.Bugs.ListByProducts(1, zentao.ListOptions{
	// 	Page:   1,
	// 	Limit:  100,
	// 	Status: "active", // 列出激活的
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("product 1 bugs count: %v", len(p12.Bugs))
}
