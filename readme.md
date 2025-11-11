# giniladmin

## 项目简介

giniladmin 是一个高度模块化、便于扩展的 Go 语言后端与前端全栈管理系统。它适用于需要多工具集成或服务模块拆分的中大型工程。后端采用“应用-组件”分层思路，每个业务应用可独立挂载多个 service 组件，方便功能解耦和二次开发。

## 架构设计与目录结构

### 目录结构
```
giniladmin/
├── cmd/              # 多应用/命令入口（主入口+多子命令）
│   ├── main.go
│   ├── options/      # 命令全局参数
│   └── subcmds/      # 支持扩展的多工具/服务命令目录
├── config/           # 各环境配置
├── constants/        # 全局常量
├── document/         # 项目内技术/接口文档
├── internal/         # 主业务逻辑实现与核心对象
├── pkg/              # 公用库和基础包
├── script/           # 运维及开发脚本
├── web/              # 前端子项目（如 frontend/ 可用于单页应用等）
│   └── frontend/
├── Makefile
├── go.mod
├── go.sum
├── version
└── .gitignore
```

### 设计理念

- **多应用/工具支持**：`cmd/` 结构允许一个工程同时支持多个命令行入口（如网络工具、文件工具、自定义服务等），每个子目录/子命令可独立开发和注册。
- **服务与组件模块化**：每个“应用”下可挂载若干 service（服务组件），每个组件实现独立 API、功能与模块职责，便于划分团队协作和单元测试。
- **前后端解耦**：后端为 RESTful API，前端（如 web/frontend）完全解耦，适应微服务及现代 SPA 工程需求。

## 快速开始

1. 克隆仓库

   ```sh
   git clone https://github.com/ilasec/giniladmin.git
   ```

2. 启动命令行/后端服务

   ```sh
   cd giniladmin
   make build   # 编译
   make run     # 启动默认应用（main.go 可扩展支持多入口）
   ```

   - 使用多应用/子命令功能详见 `cmd/subcmds/` 下相关代码及文档。

3. 启动前端

   ```sh
   cd web/frontend
   npm install
   npm run dev
   ```

## 贡献与定制

- 欢迎以“新命令模块”或“service 组件”的方式贡献代码。
- 文档与最佳实践见 `document/` 目录。

## 许可证

MIT

