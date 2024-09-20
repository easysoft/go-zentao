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

package zentao

// UserType 用户类型
type UserType string

var (
	InSideUser  UserType = "inside"  // 内部用户
	OutSideUser UserType = "outside" // 外部用户
)

type UserGender string

var (
	ManGender   UserGender = "m" // 男
	WomanGender UserGender = "f" // 女
)

type ACL string // 访问控制

var (
	ACLPrivate ACL = "private" // 私有
	ACLOpen    ACL = "open"    // 公开
)

type StoriesSource string // 需求来源

var (
	SourceCustomer StoriesSource = "customer" // 客户
	SourceMarket   StoriesSource = "market"   // 市场
	SourceProduct  StoriesSource = "po"       // 产品
	SourceUser     StoriesSource = "user"     // 用户
)

type StoriesStage string // 需求状态

var (
	StageWait       StoriesStage = "wait"       // 未开始
	StagePlanned    StoriesStage = "planned"    // 已计划
	StageProjected  StoriesStage = "projected"  // 已立项
	StageDeveloping StoriesStage = "developing" // 开发中
	StageDeveloped  StoriesStage = "developed"  // 研发完毕
	StageTesting    StoriesStage = "testing"    // 测试中
	StageTested     StoriesStage = "tested"     // 测试完毕
	StageVerified   StoriesStage = "verified"   // 已验收
	StageReleased   StoriesStage = "released"   // 已发布
	StageClosed     StoriesStage = "closed"     // 已关闭
)

type StoriesCategory string // 需求分类

var (
	CategoryFeature     StoriesCategory = "feature"     // 功能
	CategoryInterface   StoriesCategory = "interface"   // 接口
	CategoryPerformance StoriesCategory = "performance" // 性能
	CategorySafe        StoriesCategory = "safe"        // 安全
	CategoryExperience  StoriesCategory = "experience"  // 体验
	CategoryImprove     StoriesCategory = "improve"     // 改进
	CategoryOther       StoriesCategory = "other"       // 其他
)

type StoriesStatus string // 需求状态

var (
	StatusDraft   StoriesStatus = "draft"   // 草稿
	StatusActive  StoriesStatus = "active"  // 激活
	StatusClosed  StoriesStatus = "closed"  // 关闭
	StatusChanged StoriesStatus = "changed" // 已变更
)

type StoriesCloseReason string // 需求关闭原因

var (
	CloseReasonDone      StoriesCloseReason = "done"      // 完成
	CloseReasonDuplicate StoriesCloseReason = "duplicate" // 重复
	CloseReasonPostponed StoriesCloseReason = "postponed" // 延期
	CloseReasonwillnotdo StoriesCloseReason = "willnotdo" // 不做
	CloseReasonCancel    StoriesCloseReason = "cancel"    // 取消
	CloseReasonBydesign  StoriesCloseReason = "bydesign"  // 设计如此
)

type StoriesReviewResult string // 需求评审结果

var (
	ReviewResultPass    StoriesReviewResult = "pass"    // 通过
	ReviewResultRevert  StoriesReviewResult = "revert"  // 退回 撤销变更
	ReviewResultClarify StoriesReviewResult = "clarify" // 有待明确
)

type ProjectAuth string // 项目权限

var (
	ProjectExtend ProjectAuth = "extend" // 继承
	ProjectReset  ProjectAuth = "reset"  // 重新定义
)

type ExecutionLifeTime string // 生命周期

var (
	LifeTimeShort ExecutionLifeTime = "short" // 短期
	LifeTimeLong  ExecutionLifeTime = "long"  // 长期
	LifeTimeOps   ExecutionLifeTime = "ops"   // 运维
	LifeTimeNull  ExecutionLifeTime = ""      // 未定义
)

// CustomResp 通用Resp
type CustomResp struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
