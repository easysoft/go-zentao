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

package buildin

import (
	"testing"
)

func TestUserService(t *testing.T) {
	// 这个测试需要真实的禅道服务器和凭据，所以在CI环境中跳过
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// 这里可以添加集成测试，但需要真实的禅道环境
}

func TestUserTaskBugStoryResponse_GetItemCounts(t *testing.T) {
	// 测试数据结构的方法
	resp := &UserTaskBugStoryResponse{
		Task: TaskList{
			Total: 5,
			Tasks: []UserTaskDetail{},
		},
		Bug: BugList{
			Total: 3,
			Bugs:  []UserBugDetail{},
		},
		Story: StoryList{
			Total:   2,
			Stories: []UserStoryDetail{},
		},
	}

	taskCount, bugCount, storyCount := resp.GetItemCounts()

	if taskCount != 5 {
		t.Errorf("Expected task count 5, got %d", taskCount)
	}
	if bugCount != 3 {
		t.Errorf("Expected bug count 3, got %d", bugCount)
	}
	if storyCount != 2 {
		t.Errorf("Expected story count 2, got %d", storyCount)
	}
}

func TestTaskList_FilterTasksByStatus(t *testing.T) {
	taskList := &TaskList{
		Total: 3,
		Tasks: []UserTaskDetail{
			{ID: 1, Name: "Task 1", Status: "doing"},
			{ID: 2, Name: "Task 2", Status: "wait"},
			{ID: 3, Name: "Task 3", Status: "doing"},
		},
	}

	doingTasks := taskList.FilterTasksByStatus("doing")
	waitTasks := taskList.FilterTasksByStatus("wait")
	doneTasks := taskList.FilterTasksByStatus("done")

	if len(doingTasks) != 2 {
		t.Errorf("Expected 2 doing tasks, got %d", len(doingTasks))
	}
	if len(waitTasks) != 1 {
		t.Errorf("Expected 1 wait task, got %d", len(waitTasks))
	}
	if len(doneTasks) != 0 {
		t.Errorf("Expected 0 done tasks, got %d", len(doneTasks))
	}

	// 验证过滤结果的正确性
	if doingTasks[0].ID != 1 || doingTasks[1].ID != 3 {
		t.Error("Filtered doing tasks do not match expected IDs")
	}
	if waitTasks[0].ID != 2 {
		t.Error("Filtered wait task does not match expected ID")
	}
}

func TestBugList_FilterBugsByStatus(t *testing.T) {
	bugList := &BugList{
		Total: 3,
		Bugs: []UserBugDetail{
			{ID: 1, Title: "Bug 1", Status: "active"},
			{ID: 2, Title: "Bug 2", Status: "resolved"},
			{ID: 3, Title: "Bug 3", Status: "active"},
		},
	}

	activeBugs := bugList.FilterBugsByStatus("active")
	resolvedBugs := bugList.FilterBugsByStatus("resolved")
	closedBugs := bugList.FilterBugsByStatus("closed")

	if len(activeBugs) != 2 {
		t.Errorf("Expected 2 active bugs, got %d", len(activeBugs))
	}
	if len(resolvedBugs) != 1 {
		t.Errorf("Expected 1 resolved bug, got %d", len(resolvedBugs))
	}
	if len(closedBugs) != 0 {
		t.Errorf("Expected 0 closed bugs, got %d", len(closedBugs))
	}
}

func TestStoryList_FilterStoriesByStatus(t *testing.T) {
	storyList := &StoryList{
		Total: 3,
		Stories: []UserStoryDetail{
			{ID: 1, Title: "Story 1", Status: "active"},
			{ID: 2, Title: "Story 2", Status: "closed"},
			{ID: 3, Title: "Story 3", Status: "active"},
		},
	}

	activeStories := storyList.FilterStoriesByStatus("active")
	closedStories := storyList.FilterStoriesByStatus("closed")
	draftStories := storyList.FilterStoriesByStatus("draft")

	if len(activeStories) != 2 {
		t.Errorf("Expected 2 active stories, got %d", len(activeStories))
	}
	if len(closedStories) != 1 {
		t.Errorf("Expected 1 closed story, got %d", len(closedStories))
	}
	if len(draftStories) != 0 {
		t.Errorf("Expected 0 draft stories, got %d", len(draftStories))
	}
}

func TestUserTaskBugStoryResponse_GetActiveItems(t *testing.T) {
	resp := &UserTaskBugStoryResponse{
		Task: TaskList{
			Total: 2,
			Tasks: []UserTaskDetail{
				{ID: 1, Name: "Task 1", Status: "doing"},
				{ID: 2, Name: "Task 2", Status: "wait"},
			},
		},
		Bug: BugList{
			Total: 2,
			Bugs: []UserBugDetail{
				{ID: 1, Title: "Bug 1", Status: "active"},
				{ID: 2, Title: "Bug 2", Status: "resolved"},
			},
		},
		Story: StoryList{
			Total: 2,
			Stories: []UserStoryDetail{
				{ID: 1, Title: "Story 1", Status: "active"},
				{ID: 2, Title: "Story 2", Status: "closed"},
			},
		},
	}

	activeTasks, activeBugs, activeStories := resp.GetActiveItems()

	if len(activeTasks) != 1 {
		t.Errorf("Expected 1 active task, got %d", len(activeTasks))
	}
	if len(activeBugs) != 1 {
		t.Errorf("Expected 1 active bug, got %d", len(activeBugs))
	}
	if len(activeStories) != 1 {
		t.Errorf("Expected 1 active story, got %d", len(activeStories))
	}

	// 验证返回的活跃项目是正确的
	if activeTasks[0].Status != "doing" {
		t.Error("Active task should have 'doing' status")
	}
	if activeBugs[0].Status != "active" {
		t.Error("Active bug should have 'active' status")
	}
	if activeStories[0].Status != "active" {
		t.Error("Active story should have 'active' status")
	}
}
