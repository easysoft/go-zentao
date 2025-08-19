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
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type UserService struct {
	client *Client
}

type UserTaskBugStoryResponse struct {
	Profile UserProfile `json:"profile"`
	Task    TaskList    `json:"task"`
	Bug     BugList     `json:"bug"`
	Story   StoryList   `json:"story"`
}

type TaskList struct {
	Total int              `json:"total"`
	Tasks []UserTaskDetail `json:"tasks"`
}

type BugList struct {
	Total int             `json:"total"`
	Bugs  []UserBugDetail `json:"bugs"`
}

type StoryList struct {
	Total   int               `json:"total"`
	Stories []UserStoryDetail `json:"stories"`
}

type UserProfile struct {
	ID            int      `json:"id"`
	Company       int      `json:"company"`
	Dept          int      `json:"dept"`
	Account       string   `json:"account"`
	Type          string   `json:"type"`
	Realname      string   `json:"realname"`
	Superior      string   `json:"superior"`
	Pinyin        string   `json:"pinyin"`
	Role          any      `json:"role"`
	Post          int      `json:"post"`
	InitialPost   int      `json:"initialPost"`
	AuthTrain     string   `json:"authTrain"`
	Nickname      string   `json:"nickname"`
	Commiter      string   `json:"commiter"`
	Avatar        string   `json:"avatar"`
	Birthday      any      `json:"birthday"`
	Gender        string   `json:"gender"`
	Email         string   `json:"email"`
	Skype         string   `json:"skype"`
	QQ            string   `json:"qq"`
	Yahoo         string   `json:"yahoo"`
	GTalk         string   `json:"gtalk"`
	Wangwang      string   `json:"wangwang"`
	Mobile        string   `json:"mobile"`
	Phone         string   `json:"phone"`
	Weixin        string   `json:"weixin"`
	Dingding      string   `json:"dingding"`
	Slack         string   `json:"slack"`
	Whatsapp      string   `json:"whatsapp"`
	Address       string   `json:"address"`
	Zipcode       string   `json:"zipcode"`
	Nature        string   `json:"nature"`
	Analysis      string   `json:"analysis"`
	Strategy      string   `json:"strategy"`
	Join          any      `json:"join"`
	Visits        int      `json:"visits"`
	Visions       string   `json:"visions"`
	IP            string   `json:"ip"`
	Last          string   `json:"last"`
	Fails         int      `json:"fails"`
	Locked        any      `json:"locked"`
	Feedback      string   `json:"feedback"`
	Ranzhi        string   `json:"ranzhi"`
	Ldap          string   `json:"ldap"`
	Score         int      `json:"score"`
	ScoreLevel    int      `json:"scoreLevel"`
	ResetToken    string   `json:"resetToken"`
	Deleted       string   `json:"deleted"`
	Status        string   `json:"status"`
	ClientLang    string   `json:"clientLang"`
	ClientStatus  string   `json:"clientStatus"`
	Admin         bool     `json:"admin"`
	SuperReviewer bool     `json:"superReviewer"`
	View          UserView `json:"view"`
}

type UserView struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Programs string `json:"programs"`
	Sprints  string `json:"sprints"`
	Products string `json:"products"`
	Projects string `json:"projects"`
}

type UserTaskDetail struct {
	ID                 int         `json:"id"`
	Project            int         `json:"project"`
	Parent             int         `json:"parent"`
	IsParent           int         `json:"isParent"`
	IsTpl              int         `json:"isTpl"`
	Path               string      `json:"path"`
	Execution          int         `json:"execution"`
	Module             int         `json:"module"`
	Design             int         `json:"design"`
	Story              int         `json:"story"`
	StoryVersion       int         `json:"storyVersion"`
	DesignVersion      int         `json:"designVersion"`
	FromBug            int         `json:"fromBug"`
	FromIssue          int         `json:"fromIssue"`
	Docs               string      `json:"docs"`
	DocVersions        string      `json:"docVersions"`
	Feedback           int         `json:"feedback"`
	Name               string      `json:"name"`
	Type               string      `json:"type"`
	Mode               string      `json:"mode"`
	Pri                int         `json:"pri"`
	Estimate           int         `json:"estimate"`
	Consumed           int         `json:"consumed"`
	Left               int         `json:"left"`
	Deadline           string      `json:"deadline"`
	Status             string      `json:"status"`
	SubStatus          string      `json:"subStatus"`
	Color              string      `json:"color"`
	Mailto             string      `json:"mailto"`
	Keywords           string      `json:"keywords"`
	Desc               string      `json:"desc"`
	Version            int         `json:"version"`
	OpenedBy           string      `json:"openedBy"`
	OpenedDate         string      `json:"openedDate"`
	AssignedTo         string      `json:"assignedTo"`
	AssignedDate       string      `json:"assignedDate"`
	EstStarted         string      `json:"estStarted"`
	RealStarted        any         `json:"realStarted"`
	EstimateStartDate  any         `json:"estimateStartDate"`
	ActualStartDate    any         `json:"actualStartDate"`
	FinishedBy         string      `json:"finishedBy"`
	FinishedDate       any         `json:"finishedDate"`
	FinishedList       string      `json:"finishedList"`
	CanceledBy         string      `json:"canceledBy"`
	CanceledDate       any         `json:"canceledDate"`
	ClosedBy           string      `json:"closedBy"`
	ClosedDate         any         `json:"closedDate"`
	RealDuration       int         `json:"realDuration"`
	PlanDuration       int         `json:"planDuration"`
	ClosedReason       string      `json:"closedReason"`
	LastEditedBy       string      `json:"lastEditedBy"`
	LastEditedDate     string      `json:"lastEditedDate"`
	ActivatedDate      any         `json:"activatedDate"`
	Order              int         `json:"order"`
	Repo               int         `json:"repo"`
	MR                 int         `json:"mr"`
	Entry              string      `json:"entry"`
	Lines              string      `json:"lines"`
	V1                 string      `json:"v1"`
	V2                 string      `json:"v2"`
	Deleted            string      `json:"deleted"`
	Vision             string      `json:"vision"`
	ReplacetypeBy      string      `json:"replacetypeBy"`
	ReplacetypeDate    any         `json:"replacetypeDate"`
	ExecutionID        int         `json:"executionID"`
	ExecutionName      string      `json:"executionName"`
	ProjectName        string      `json:"projectName"`
	ExecutionMultiple  string      `json:"executionMultiple"`
	ExecutionType      string      `json:"executionType"`
	StoryID            interface{} `json:"storyID"`
	StoryTitle         string      `json:"storyTitle"`
	StoryStatus        string      `json:"storyStatus"`
	LatestStoryVersion interface{} `json:"latestStoryVersion"`
	PriOrder           int         `json:"priOrder"`
	NeedConfirm        bool        `json:"needConfirm"`
	AssignedToRealName string      `json:"assignedToRealName"`
	Progress           int         `json:"progress"`
	RawParent          int         `json:"rawParent"`
	DesignName         string      `json:"designName"`
	EstimateLabel      string      `json:"estimateLabel"`
	ConsumedLabel      string      `json:"consumedLabel"`
	LeftLabel          string      `json:"leftLabel"`
	CanBeChanged       bool        `json:"canBeChanged"`
	IsChild            bool        `json:"isChild"`
	ParentName         string      `json:"parentName"`
}

type UserBugDetail struct {
	ID             int    `json:"id"`
	Project        int    `json:"project"`
	Product        int    `json:"product"`
	Injection      string `json:"injection"`
	Identify       string `json:"identify"`
	Branch         int    `json:"branch"`
	Module         int    `json:"module"`
	Execution      int    `json:"execution"`
	Plan           int    `json:"plan"`
	Story          int    `json:"story"`
	StoryVersion   int    `json:"storyVersion"`
	Task           int    `json:"task"`
	ToTask         int    `json:"toTask"`
	ToStory        int    `json:"toStory"`
	Title          string `json:"title"`
	Keywords       string `json:"keywords"`
	Severity       int    `json:"severity"`
	Pri            int    `json:"pri"`
	Type           string `json:"type"`
	OS             string `json:"os"`
	Browser        string `json:"browser"`
	Hardware       string `json:"hardware"`
	Found          string `json:"found"`
	Steps          string `json:"steps"`
	Status         string `json:"status"`
	SubStatus      string `json:"subStatus"`
	Color          string `json:"color"`
	Confirmed      int    `json:"confirmed"`
	ActivatedCount int    `json:"activatedCount"`
	ActivatedDate  any    `json:"activatedDate"`
	FeedbackBy     string `json:"feedbackBy"`
	NotifyEmail    string `json:"notifyEmail"`
	Mailto         string `json:"mailto"`
	Reason         string `json:"reason"`
	WritedBy       string `json:"writedBy"`
	InjectCommit   string `json:"injectCommit"`
	OpenedBy       string `json:"openedBy"`
	OpenedDate     string `json:"openedDate"`
	OpenedBuild    string `json:"openedBuild"`
	AssignedTo     string `json:"assignedTo"`
	AssignedDate   string `json:"assignedDate"`
	Deadline       string `json:"deadline"`
	ResolvedBy     string `json:"resolvedBy"`
	Resolution     string `json:"resolution"`
	ResolvedBuild  string `json:"resolvedBuild"`
	ResolvedDate   string `json:"resolvedDate"`
	ClosedBy       string `json:"closedBy"`
	ClosedDate     string `json:"closedDate"`
	DuplicateBug   int    `json:"duplicateBug"`
	RelatedBug     string `json:"relatedBug"`
	Case           int    `json:"case"`
	CaseVersion    int    `json:"caseVersion"`
	Feedback       int    `json:"feedback"`
	Result         int    `json:"result"`
	TestTask       int    `json:"testtask"`
	Repo           int    `json:"repo"`
	MR             int    `json:"mr"`
	Entry          string `json:"entry"`
	Lines          string `json:"lines"`
	V1             string `json:"v1"`
	V2             string `json:"v2"`
	RepoType       string `json:"repoType"`
	IssueKey       string `json:"issueKey"`
	LastEditedBy   string `json:"lastEditedBy"`
	LastEditedDate string `json:"lastEditedDate"`
	Deleted        string `json:"deleted"`
	PriOrder       int    `json:"priOrder"`
	SeverityOrder  int    `json:"severityOrder"`
	ProductName    string `json:"productName"`
	Shadow         int    `json:"shadow"`
	ConfirmeObject []any  `json:"confirmeObject"`
	URs            string `json:"URs"`
	StatusName     string `json:"statusName"`
}

type UserStoryDetail struct {
	ID                  int         `json:"id"`
	Product             int         `json:"product"`
	Project             interface{} `json:"project"`
	Branch              int         `json:"branch"`
	Module              int         `json:"module"`
	Plan                string      `json:"plan"`
	Source              string      `json:"source"`
	SourceNote          string      `json:"sourceNote"`
	FromBug             int         `json:"fromBug"`
	Title               string      `json:"title"`
	Type                string      `json:"type"`
	Pri                 int         `json:"pri"`
	Estimate            any         `json:"estimate"`
	Status              string      `json:"status"`
	SubStatus           string      `json:"subStatus"`
	Color               string      `json:"color"`
	Stage               string      `json:"stage"`
	StagedBy            any         `json:"stagedBy"`
	Mailto              any         `json:"mailto"`
	Keywords            string      `json:"keywords"`
	OpenedBy            any         `json:"openedBy"`
	OpenedDate          string      `json:"openedDate"`
	AssignedTo          any         `json:"assignedTo"`
	AssignedDate        string      `json:"assignedDate"`
	LastEditedBy        any         `json:"lastEditedBy"`
	LastEditedDate      string      `json:"lastEditedDate"`
	ReviewedBy          any         `json:"reviewedBy"`
	ReviewedDate        string      `json:"reviewedDate"`
	ClosedBy            any         `json:"closedBy"`
	ClosedDate          string      `json:"closedDate"`
	ClosedReason        string      `json:"closedReason"`
	ToBug               int         `json:"toBug"`
	ChildStories        string      `json:"childStories"`
	LinkStories         string      `json:"linkStories"`
	DuplicateStory      int         `json:"duplicateStory"`
	Version             int         `json:"version"`
	Parent              any         `json:"parent"`
	URChanged           string      `json:"URChanged"`
	Deleted             any         `json:"deleted"`
	PlanTitle           string      `json:"planTitle"`
	Category            string      `json:"category"`
	Reviewer            any         `json:"reviewer"`
	NotReview           any         `json:"notReview"`
	NeedSummaryEstimate bool        `json:"needSummaryEstimate"`
	ProductStatus       string      `json:"productStatus"`
}

type UserFields struct {
	Task  bool `json:"task,omitempty"`
	Bug   bool `json:"bug,omitempty"`
	Story bool `json:"story,omitempty"`
}

// GetUserTasksBugsStories 获取用户的任务、Bug、故事列表
// GET /user?fields=task,bug,story
func (c *Client) GetUserTasksBugsStories(fields ...string) (*UserTaskBugStoryResponse, *req.Response, error) {
	var result UserTaskBugStoryResponse

	// 构建查询参数
	queryParams := make(map[string]string)
	if len(fields) > 0 {
		queryParams["fields"] = strings.Join(fields, ",")
	} else {
		// 默认获取所有字段
		queryParams["fields"] = "task,bug,story"
	}

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetCookies(c.buildAllCookies()...).
		SetQueryParams(queryParams).
		SetSuccessResult(&result).
		Get(c.buildAPIURL("/user"))

	return &result, resp, err
}

// GetUserTasks 仅获取用户的任务列表
func (c *Client) GetUserTasks() (*TaskList, *req.Response, error) {
	result, resp, err := c.GetUserTasksBugsStories("task")
	if err != nil {
		return nil, resp, err
	}
	return &result.Task, resp, nil
}

// GetUserBugs 仅获取用户的Bug列表
func (c *Client) GetUserBugs() (*BugList, *req.Response, error) {
	result, resp, err := c.GetUserTasksBugsStories("bug")
	if err != nil {
		return nil, resp, err
	}
	return &result.Bug, resp, nil
}

// GetUserStories 仅获取用户的故事列表
func (c *Client) GetUserStories() (*StoryList, *req.Response, error) {
	result, resp, err := c.GetUserTasksBugsStories("story")
	if err != nil {
		return nil, resp, err
	}
	return &result.Story, resp, nil
}

// buildAllCookies 构建所有需要的cookies，包括zentaosid，确保无重复
func (c *Client) buildAllCookies() []*http.Cookie {
	cookies := []*http.Cookie{
		{
			Name:  "lang",
			Value: "zh-cn",
		},
		{
			Name:  "device",
			Value: "desktop",
		},
		{
			Name:  "theme",
			Value: "default",
		},
		{
			Name:  "vision",
			Value: "rnd",
		},
		{
			Name:  "tab",
			Value: "execution",
		},
	}

	// 只在有zentaosid时添加，避免重复
	if c.zentaosid != "" {
		cookies = append(cookies, &http.Cookie{
			Name:  "zentaosid",
			Value: c.zentaosid,
		})
	}

	return cookies
}

// buildAPIURL 构建API URL
func (c *Client) buildAPIURL(path string) string {
	return fmt.Sprintf("%s/%s%s", c.baseURL.String(), apiVersionPath, path)
}

// GetUserProfile 获取用户基本信息
func (c *Client) GetUserProfile() (*UserProfile, *req.Response, error) {
	result, resp, err := c.GetUserTasksBugsStories()
	if err != nil {
		return nil, resp, err
	}
	return &result.Profile, resp, nil
}

// FilterTasksByStatus 根据状态过滤任务
func (tl *TaskList) FilterTasksByStatus(status string) []UserTaskDetail {
	var filtered []UserTaskDetail
	for _, task := range tl.Tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// FilterBugsByStatus 根据状态过滤Bug
func (bl *BugList) FilterBugsByStatus(status string) []UserBugDetail {
	var filtered []UserBugDetail
	for _, bug := range bl.Bugs {
		if bug.Status == status {
			filtered = append(filtered, bug)
		}
	}
	return filtered
}

// FilterStoriesByStatus 根据状态过滤故事
func (sl *StoryList) FilterStoriesByStatus(status string) []UserStoryDetail {
	var filtered []UserStoryDetail
	for _, story := range sl.Stories {
		if story.Status == status {
			filtered = append(filtered, story)
		}
	}
	return filtered
}

// GetActiveItems 获取活跃的工作项
func (resp *UserTaskBugStoryResponse) GetActiveItems() (activeTasks []UserTaskDetail, activeBugs []UserBugDetail, activeStories []UserStoryDetail) {
	activeTasks = resp.Task.FilterTasksByStatus("doing")
	activeBugs = resp.Bug.FilterBugsByStatus("active")
	activeStories = resp.Story.FilterStoriesByStatus("active")
	return
}

// GetItemCounts 获取各类型工作项的数量统计
func (resp *UserTaskBugStoryResponse) GetItemCounts() (taskCount, bugCount, storyCount int) {
	return resp.Task.Total, resp.Bug.Total, resp.Story.Total
}

// GetCompletedToday 获取今天完成的工作项
func (resp *UserTaskBugStoryResponse) GetCompletedToday() (completedTasks []UserTaskDetail, resolvedBugs []UserBugDetail, closedStories []UserStoryDetail) {
	today := time.Now().Format("2006-01-02")

	// 过滤今天完成的任务
	for _, task := range resp.Task.Tasks {
		if task.Status == "done" && strings.HasPrefix(task.FinishedDate.(string), today) {
			completedTasks = append(completedTasks, task)
		}
	}

	// 过滤今天解决的Bug
	for _, bug := range resp.Bug.Bugs {
		if bug.Status == "resolved" && strings.HasPrefix(bug.ResolvedDate, today) {
			resolvedBugs = append(resolvedBugs, bug)
		}
	}

	// 过滤今天关闭的故事
	for _, story := range resp.Story.Stories {
		if story.Status == "closed" && strings.HasPrefix(story.ClosedDate, today) {
			closedStories = append(closedStories, story)
		}
	}

	return
}
