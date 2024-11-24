# blog

## 简介

本项目是基于 [bluebell-plus](https://github.com/mao888/bluebell-plus) 进行的功能改进与扩展，旨在进一步优化项目的性能和用户体验。本项目包括后端和前端两个部分，分别实现数据处理和用户界面展示。

---

## 项目结构

### 后端结构

后端基于 **Go** 语言开发，目录结构如下：
```bash
├─bin
├─conf
├─controller
├─dao
│  ├─api
│  ├─mysql
│  └─redis
├─docs
├─logger
├─logic
├─middlewares
├─models
├─pkg
│  ├─jwt
│  ├─smms
│  └─snowflake
├─routers
├─settings
├─static
│  ├─css
│  ├─fonts
│  ├─img
│  └─js
└─templates

```

主要功能包括：
- 数据库管理（MySQL, Redis）
- RESTful API 实现
- 用户认证与授权
- 日志记录与性能优化

### 前端结构

前端基于 **Vue.js** 框架，目录结构如下：
```bash
├─public
│  └─static
└─src
    ├─assets
    │  ├─css
    │  ├─font
    │  └─images
    ├─components
    ├─router
    ├─service
    ├─store
    ├─utils
    └─views
        └─components
```

主要功能包括：
- 界面交互与动态展示
- API 数据绑定与状态管理
- 响应式设计与多设备兼容

---

## 部署步骤

### 环境准备

1. 安装 **Docker** 和 **Docker Compose**。
2. 确保本地安装了 **Go 1.19+** 和 **Node.js 16+**。

### 后端部署

1. 克隆代码仓库并进入项目目录：
```bash
git clone https://github.com/Gezelligheid1010/blog
cd <项目目录>
```
2. 启动后端服务
```bash
docker-compose up -d
```

### 前端部署
1. 进入前端目录：
```bash
cd <项目目录>/frontend
```
2. 安装依赖并启动服务：
```bash
npm install
npm run serve
```
3. 打开浏览器访问 http://localhost:8080。

### 主要改动
### 1. 评论功能改造
- **原有方式**：评论数据通过外部 API 调用获取。
- **改进方案**：  
  - 将评论信息存储到本地 MySQL 数据库，设计了全新的数据库表结构。
  - 开发了后端 API，实现用户增加评论的功能。
  - 改写前端评论组件，提升了评论加载的稳定性和响应速度。

### 2. 社区模块替换
- **原有模块**：展示 GitHub 热门项目的功能。
- **改进方案**：  
  - 删除 GitHub 热门项目模块。
  - 新增了社区项目模块，用于展示用户社区内的精选项目。
  - 重新设计了模块的前端 UI，并调整了相关 API 逻辑，确保功能简洁易用。


### 许可证

本项目基于 [MIT License]() 开源。感谢原项目的贡献！
