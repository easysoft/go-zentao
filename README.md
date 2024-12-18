# go-zentao

[![Go Report Card](https://goreportcard.com/badge/github.com/easysoft/go-zentao)](https://goreportcard.com/report/github.com/easysoft/go-zentao)
![GitHub](https://img.shields.io/github/license/easysoft/go-zentao?style=flat-square)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/easysoft/go-zentao?filename=go.mod&style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/easysoft/go-zentao?style=flat-square)
![zentao](https://goproxy.cn/stats/github.com/easysoft/go-zentao/badges/download-count.svg)

Zentao API client enabling Go programs to interact with Zentao in a simple and uniform way

## NOTE

基于ZenTao最新开源版本[`21.2`](https://github.com/easysoft/zentaopms/releases/tag/zentaopms_21.2_20241205)

## API

### 支持Restful

- [x] Token
  - [x] 获取Token
  - [x] Ping
- [x] 用户(User)
  - [x] 获取用户信息
  - [x] 获取我的个人信息
  - [x] 删除用户
  - [x] 创建用户
  - [x] 修改用户信息
  - [x] 获取用户列表
- [x] 项目集(Programs)
  - [x] 获取项目集列表
  - [x] 修改项目集
  - [x] 获取项目集详情
  - [x] 删除项目集
  - [x] 创建项目集
- [x] 产品(Products)
  - [x] 获取产品列表
  - [x] 修改产品
  - [x] 获取产品详情
  - [x] 删除产品
  - [x] 创建产品
- [x] 产品计划(ProductsPlans)
  - [x] 创建计划
  - [x] 创建子计划
  - [x] 获取产品计划列表
  - [x] 获取计划详情
  - [x] 修改计划
  - [x] 删除计划
  - [x] 产品计划关联需求
  - [x] 产品计划取消关联需求
  - [x] 产品计划关联Bug
  - [x] 产品计划取消关联Bug
  - [ ] 操作计划(开启、关闭、完成)
- [x] 发布(Releases)
  - [x] 获取项目发布列表
  - [x] 获取产品发布列表
- [x] 需求(Stories)
  - [x] 获取项目需求列表
  - [x] 获取产品需求列表
  - [x] 获取执行需求列表
  - [x] 变更需求
  - [x] 创建需求
  - [x] 获取需求详情
  - [x] 删除需求
  - [x] 修改需求其他字段
  - [x] 激活需求
  - [x] 关闭需求
  - [x] 指派需求
  - [ ] 获取工时(开源版不支持)
  - [ ] 预估工时(开源版不支持)
  - [ ] 子需求
  - [x] 撤回评审
  - [x] 评审需求
- [x] 项目(Projects)
  - [x] 创建项目
  - [x] 获取项目列表
  - [x] 获取项目详情
  - [x] 修改项目
  - [x] 删除项目
- [x] 执行(Executions)
  - [x] 获取项目的执行列表
  - [x] 创建执行
  - [x] 修改执行
  - [x] 查看执行详情
  - [x] 删除执行
- [x] 任务(Tasks)
  - [x] 获取执行任务列表
  - [x] 获取任务详情
  - [x] 创建任务
  - [x] 删除任务
  - [x] 修改任务
- [x] Bug
  - [x] 获取产品Bug列表
  - [x] 获取项目Bug列表
  - [x] 获取执行Bug列表
  - [x] 获取Bug详情
  - [x] 创建Bug
  - [x] 删除Bug
  - [x] 修改Bug
  - [x] 关闭Bug
  - [x] 指派Bug
  - [x] 确认Bug
  - [x] 解决Bug
  - [x] 激活Bug
  - [ ] 评估Bug(开源版不支持)
- [x] 版本(Builds)
  - [x] 获取项目版本列表
  - [x] 获取执行版本详情
  - [x] 修改版本
  - [x] 创建版本
  - [x] 获取版本详情
  - [x] 删除版本
- [x] 用例(TestCases)
  - [x] 获取产品用例列表
  - [x] 获取用例详情
  - [x] 创建用例
  - [x] 删除用例
  - [x] 修改用例
- [x] 测试单(TestTasks)
  - [x] 获取测试单列表
  - [x] 获取项目的测试单
  - [x] 获取测试单详情
  - [x] 创建测试单
  - [x] 删除测试单
  - [x] 修改测试单
- [x] 反馈(Feedbacks)
  - [x] 获取反馈列表
  - [x] 获取反馈详情
  - [x] 创建反馈
  - [x] 删除反馈
  - [x] 修改反馈
  - [x] 关闭反馈
  - [x] 指派反馈
- [x] 模块(Modules)
  - [x] 获取模块列表
- [ ] 其他(Misc)
  - [x] 获取系统配置参数
  - [x] 获取禅道版本

### 支持PathINFO/Get内置页面接口

> 有限支持, 是对Restful API的拓展

## 使用

```go
import "github.com/easysoft/go-zentao/v20/zentao"
```

## 测试账号

> 推荐本地部署

```bash
#部署方式:
docker compose -f hack/docker-compose.yml up -d
```


