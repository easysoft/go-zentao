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

	"github.com/easysoft/go-zentao/v21/buildin"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// 创建客户端（需要配置真实的URL、用户名和密码）
	zt, err := buildin.NewBasicAuthClient(
		"admin",
		"jaege1ugh4ooYip7",
		buildin.WithBaseURL("http://127.0.0.1"),
		buildin.WithoutProxy(),
	)
	if err != nil {
		panic(err)
	}

	// 获取用户的所有任务、Bug和故事
	fmt.Println("=== 获取用户的所有工作项 ===")
	allItems, _, err := zt.GetUserTasksBugsStories()
	if err != nil {
		panic(err)
	}

	// 显示用户基本信息
	fmt.Printf("用户: %s (%s)\n", allItems.Profile.Realname, allItems.Profile.Account)
	fmt.Printf("部门ID: %d\n", allItems.Profile.Dept)
	fmt.Printf("邮箱: %s\n", allItems.Profile.Email)
	fmt.Println()

	// 显示统计信息
	taskCount, bugCount, storyCount := allItems.GetItemCounts()
	fmt.Printf("任务总数: %d\n", taskCount)
	fmt.Printf("Bug总数: %d\n", bugCount)
	fmt.Printf("故事总数: %d\n", storyCount)
	fmt.Println()

	// 显示活跃工作项
	fmt.Println("=== 活跃工作项 ===")
	activeTasks, activeBugs, activeStories := allItems.GetActiveItems()
	fmt.Printf("进行中的任务: %d\n", len(activeTasks))
	for _, task := range activeTasks {
		fmt.Printf("  - [%d] %s (优先级: %d, 进度: %d%%)\n",
			task.ID, task.Name, task.Pri, task.Progress)
	}

	fmt.Printf("活跃的Bug: %d\n", len(activeBugs))
	for _, bug := range activeBugs {
		fmt.Printf("  - [%d] %s (严重程度: %d, 优先级: %d)\n",
			bug.ID, bug.Title, bug.Severity, bug.Pri)
	}

	fmt.Printf("活跃的故事: %d\n", len(activeStories))
	for _, story := range activeStories {
		fmt.Printf("  - [%d] %s (阶段: %s)\n",
			story.ID, story.Title, story.Stage)
	}
	fmt.Println()

	// 仅获取任务
	fmt.Println("=== 仅获取任务列表 ===")
	tasks, _, err := zt.GetUserTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("任务总数: %d\n", tasks.Total)

	// 按状态过滤任务
	doingTasks := tasks.FilterTasksByStatus("doing")
	waitTasks := tasks.FilterTasksByStatus("wait")
	doneTasks := tasks.FilterTasksByStatus("done")

	fmt.Printf("进行中: %d, 等待中: %d, 已完成: %d\n",
		len(doingTasks), len(waitTasks), len(doneTasks))

	// 仅获取Bug
	fmt.Println("\n=== 仅获取Bug列表 ===")
	bugs, _, err := zt.GetUserBugs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bug总数: %d\n", bugs.Total)

	// 按状态过滤Bug
	activeOnlyBugs := bugs.FilterBugsByStatus("active")
	resolvedBugs := bugs.FilterBugsByStatus("resolved")
	closedBugs := bugs.FilterBugsByStatus("closed")

	fmt.Printf("活跃: %d, 已解决: %d, 已关闭: %d\n",
		len(activeOnlyBugs), len(resolvedBugs), len(closedBugs))

	// 仅获取故事
	fmt.Println("\n=== 仅获取故事列表 ===")
	stories, _, err := zt.GetUserStories()
	if err != nil {
		panic(err)
	}
	fmt.Printf("故事总数: %d\n", stories.Total)

	// 按状态过滤故事
	activeOnlyStories := stories.FilterStoriesByStatus("active")
	closedOnlyStories := stories.FilterStoriesByStatus("closed")

	fmt.Printf("活跃: %d, 已关闭: %d\n",
		len(activeOnlyStories), len(closedOnlyStories))

	// 获取今天完成的工作项
	fmt.Println("\n=== 今天完成的工作项 ===")
	completedTasks, resolvedTodayBugs, closedTodayStories := allItems.GetCompletedToday()
	fmt.Printf("今天完成的任务: %d\n", len(completedTasks))
	fmt.Printf("今天解决的Bug: %d\n", len(resolvedTodayBugs))
	fmt.Printf("今天关闭的故事: %d\n", len(closedTodayStories))

	// 详细输出（可选，用于调试）
	if false { // 设为true查看详细数据结构
		fmt.Println("\n=== 详细数据结构 ===")
		spew.Dump(allItems)
	}

	fmt.Println("\n示例程序执行完成！")
}
