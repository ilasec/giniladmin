# giniladmin 架构与配置设计方案

本设计文档为 giniladmin 项目提出一套现代化、高可维护性、高可扩展性的 Go 架构与最佳配置文件组织实践，实现主框架解耦、领域驱动模块化和多应用协同配置。

---

## 一、总体项目结构设计

```
giniladmin/
├── cmd/                # 各应用/服务启动入口，多应用支持
│   ├── admin/
│   │   └── main.go
│   └── ...             # 其他应用，如 tools、worker 等
│
├── internal/           # 领域驱动核心代码，仅对本工程暴露
│   ├── user/           # 示例业务领域模块，按下述规范拆分
│   │   ├── handler.go  # API路由、参数校验、响应
│   │   ├── service.go  # 业务逻辑
│   │   ├── repo.go     # 数据访问
│   │   ├── model.go    # 域模型/DTO/VO
│   │   └── doc.go      # 文档注释（自动生成 swagger）
│   ├── file/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repo.go
│   │   ├── model.go
│   │   └── doc.go
│   └── ...             # 其他领域模块，按同一规范扩展
│
├── internal/platform/  # 平台能力(如core/middleware/event/engine)
│   ├── middleware/     # 通用中间件
│   ├── event/          # 事件或引擎能力
│   ├── utils/          # 通用工具
│   ├── conf/           # 配置加载与统一入口
│   ├── errorcode/      # 错误码标准
│   └── ...
│
├── pkg/                # 可被外部import的通用包
│
├── api/                # OpenAPI/Swagger文档、gRPC等
│
├── configs/            # 配置文件（主配置+多应用分配置）
│
├── docs/               # 项目文档与设计说明
│
├── web/                # 前端代码
│
├── Makefile
├── go.mod
└── README.md
```

---

## 二、领域驱动模块（internal/mymodule）规范

每个业务领域应自洽组织，推荐如下：

```
internal/mymodule/
├── handler.go   # 路由及 API handler
├── service.go   # 领域业务逻辑
├── repo.go      # 数据访问、外部依赖
├── model.go     # 类型与结构体定义
└── doc.go       # 模块注释（swagger友好）
```

**优势：**
- 每个领域自成闭环，责任清晰，可测试性强。
- handler/service 层间依赖通过接口实现。
- 外部依赖（数据库、三方服务等）全部下沉至 repo 层。

---

## 三、配置文件层次与加载

### 1. 组织结构

```
configs/
├── config.toml      # 主/通用配置（数据库、MQ、日志等平台公共）
├── admin.toml       # admin 应用专属配置（端口、特有参数等）
├── fileops.toml     # fileops 服务专属配置
├── ...              # 其他应用自定义配置
```

### 2. 配置内容示例

**config.toml**
```toml
[server]
http_port = 8000
log_level = "info"

[database]
dsn = "user:pass@tcp(127.0.0.1:3306)/giniladmin"
max_idle_conns = 10
max_open_conns = 100
```

**admin.toml**
```toml
[admin]
dashboard_title = "giniladmin Admin Panel"
max_sessions = 50
jwt_secret = "your_secret_token"

[server]
http_port = 9000  # 应用覆盖全局配置
```

### 3. 加载与合并方案

- 启动时优先加载全局 `config.toml`，再加载 `${app}.toml`，自动合并。
- 实现可用 viper/universal-config，支持多级env/覆盖，向下传递。
- 配置加载逻辑集中于 `internal/platform/conf/config.go`，对外提供统一接口。

**伪代码：**
```go
import "github.com/spf13/viper"

func LoadConfig(app string) *viper.Viper {
    v := viper.New()
    v.SetConfigFile("configs/config.toml")
    _ = v.ReadInConfig()
    if app != "" {
        v.SetConfigFile("configs/" + app + ".toml")
        _ = v.MergeInConfig()
    }
    return v
}
```

### 4. 管理建议与优点

- 各 app 可全自定义/扩展参数，互不干扰。
- 支持 prod/dev/test 多环境，安全易运维。
- 配置与代码解耦，支持动态热更/一键换环境。

---

## 四、整体优势

- 清晰分层，平台与业务能力解耦。
- 支持多业务/多应用并行开发，领域自洽易维护。
- 配置自上而下清晰传递，适合大型团队和长期演进。
- 易于前后端分离开发、DevOps 演进与微服务升级。

---

如需 starter 代码或具体模块设计细节，请参看 docs/ 目录其他文档或联系维护者协助。
