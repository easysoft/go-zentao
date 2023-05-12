# go-zentao

[![Go Report Card](https://goreportcard.com/badge/github.com/easysoft/go-zentao)](https://goreportcard.com/report/github.com/easysoft/go-zentao)
![GitHub](https://img.shields.io/github/license/easysoft/go-zentao?style=flat-square)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/easysoft/go-zentao?filename=go.mod&style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/easysoft/go-zentao?style=flat-square)
![zentao](https://goproxy.cn/stats/github.com/easysoft/go-zentao/badges/download-count.svg)

Zentao API client enabling Go programs to interact with Zentao in a simple and uniform way

## NOTE

基于ZenTao最新开源版本 [`18.3`](https://github.com/quicklyon/zentao-docker), 即将适配最新版本`20.dev`

## 支持

- [x] Token
- [x] 用户(User)
  - [x] 获取我的个人信息(Token)
  - [x] 删除用户
  - [x] 创建用户
  - [x] 修改用户信息
  - [x] 获取用户列表
  - [x] 获取用户信息
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
  - [x] 获取产品计划列表
  - [x] 获取计划详情
  - [x] 修改计划
  - [x] 删除计划
  - [x] 产品计划关联需求
  - [x] 产品计划取消关联需求
  - [x] 产品计划关联Bug
  - [x] 产品计划取消关联Bug
- [x] 发布
  - [x] 获取项目发布列表
  - [x] 获取产品发布列表
- [x] 需求(Stories)
  - [x] 获取项目需求列表
  - [x] 变更需求
  - [x] 获取产品需求列表
  - [x] 创建需求
  - [x] 获取执行需求列表
  - [x] 获取需求详情
  - [x] 删除需求
  - [x] 修改需求其他字段
- [x] 项目(Projects)
  - [x] 创建项目
  - [x] 获取项目列表
  - [x] 获取项目详情
  - [x] 修改项目
  - [x] 删除项目
- [x] 版本(Builds)
  - [x] 获取项目版本列表
  - [x] 获取执行版本详情
  - [x] 修改版本
  - [x] 创建版本
  - [x] 获取版本详情
  - [x] 删除版本
- [x] 执行(Executions)
  - [x] 获取项目的执行列表
  - [x] 创建执行
  - [x] 修改执行
  - [x] 查看执行详情
  - [x] 删除执行
- [x] 任务
  - [x] 获取执行任务列表
  - [x] 获取任务详情
  - [x] 创建任务
  - [x] 删除任务
  - [x] 修改任务
- [x] Bug
  - [x] 获取产品Bug列表
  - [x] 获取Bug详情
  - [x] 创建Bug
  - [x] 删除Bug
  - [x] 修改Bug
- [x] 用例
  - [x] 获取产品用例列表
  - [x] 获取用例详情
  - [x] 创建用例
  - [x] 删除用例
  - [x] 修改用例
- [x] 测试单
  - [x] 获取测试单列表
  - [x] 获取项目的测试单
  - [x] 获取测试单详情
  - [x] 创建测试单
  - [x] 删除测试单
  - [x] 修改测试单
- [x] 反馈
  - [x] 获取反馈列表
  - [x] 获取反馈详情
  - [x] 创建反馈
  - [x] 删除反馈
  - [x] 修改反馈
  - [x] 关闭反馈
  - [x] 指派反馈

## 测试账号

> 推荐本地部署

```bash
# 部署方式: okteto up -f hack/okteto.yml --deploy
地址: https://zentao-easysoft.cloud.okteto.net
账号: admin/jaege1ugh4ooYip7
```

## TODO

- [ ] 优化代码 & 添加单元测试
- [ ] 支持二次开发API

### 已知问题

> 将在`20`版本优化

- [ ] 部分字段接口返回有时候是结构体有时候是字符串
