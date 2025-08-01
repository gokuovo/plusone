# PlusOne Gin API 项目

这是一个基于 Go 和 Gin 框架构建的现代化、结构清晰的 RESTful API 项目。它不仅展示了一个典型的分层 Web 应用架构，还集成了一系列生产级的最佳实践，特别是在**可观测性 (Observability)** 和**健壮性 (Robustness)** 方面。

## ✨ 核心特性

- **分层架构**: 采用经典的 `Controller -> Service -> Repository` 模式，职责分离，易于维护。
- **现代化日志系统**:
    - **结构化日志 (slog)**: 基于 Go 官方推荐的 `slog` 库，所有日志输出为 JSON 格式，便于机器解析和分析。
    - **请求追踪 (Trace ID)**: 通过中间件为每个请求自动注入唯一的 `Trace ID`，轻松串联起单次请求的完整生命周期日志。
    - **自动化请求日志**: 自动记录每个 API 请求的详细信息，包括状态码、耗时、客户端 IP 等。
- **健壮的后端实践**:
    - **事务管理**: 在关键业务（如用户注册）中实现了数据库事务，保证数据一致性。
    - **全链路上下文 (`context.Context`)**: 在所有层级中传递上下文，为超时控制和取消信号传播提供了原生支持。
    - **优雅的资源清理**: 使用 `defer` 确保在程序退出时，数据库连接和日志文件句柄被安全关闭。
- **清晰的 API 设计**:
    - **数据传输层 (DTO)**: 设立独立的 `dto` 包，使用 `Input`/`Output` 后缀明确定义 API 的数据结构，确保 API 边界清晰且安全。
    - **动态 API 文档**: 集成 `Swagger`，通过代码注解自动生成并提供交互式 API 文档。
- **高可配置性**:
    - **双数据库支持**: 通过 GORM 支持 `MySQL` 和 `SQLite`，可在配置中轻松切换。
    - **环境配置**: 通过 `.env` 文件管理所有配置，如数据库、端口、日志级别等。
- **安全**: 使用 `golang-jwt/jwt/v5` 实现安全的无状态 JWT 认证。

## 🛠️ 技术栈

- **Web 框架**: Gin
- **日志**: slog (Go 1.21+)
- **数据库 ORM**: GORM
- **API 文档**: swaggo/swag
- **唯一 ID**: google/uuid

## 🚀 快速开始

### 1. 环境要求
- Go 1.21 或更高版本
- `swag` 命令行工具

### 2. 安装与运行
```bash
# 克隆仓库
git clone https://github.com/gokuovo/plusone.git
cd plusone

# 安装依赖
go mod tidy

# 安装 swag 工具 (如果尚未安装)
go install github.com/swaggo/swag/cmd/swag@latest

# 生成 API 文档
swag init

# 运行应用 (推荐创建 .env 文件，否则将使用默认配置)
go run main.go
```

服务器将在 `.env` 文件中配置的端口上启动（默认为 `http://localhost:8080`）。

### 3. 配置说明
在项目根目录下创建一个 `.env` 文件来覆盖默认配置：

```dotenv
# 服务器端口
SERVER_PORT=8080

# JWT 密钥
JWT_SECRET=your_super_secret_key

# 数据库类型 (可选项: "mysql" 或 "sqlite")
DB_TYPE=sqlite

# 数据库连接源
# 对于 SQLite, 这是一个文件名
DB_SOURCE=oneplusone.db
# 对于 MySQL, 这是一个 DSN 字符串
# DB_SOURCE=user:password@tcp(127.0.0.1:3306)/oneplusone?charset=utf8mb4&parseTime=True&loc=Local

# GORM 日志级别 (可选项: "silent", "error", "warn", "info")
DB_LOG_LEVEL=info
```

## 📚 API 文档

项目启动后，即可在浏览器中访问交互式 API 文档：

- **Swagger UI 地址**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

> **注意**: 每当您修改了代码中的 API 注解后，都需要重新运行 `swag init` 命令来更新文档。

## 🏗️ 项目结构

```
.
├── config/         # 配置加载
├── controllers/    # 控制器 (HTTP 请求处理)
├── di/             # 依赖注入容器
├── docs/           # 由 swag 生成的 Swagger 文档
├── dto/            # 数据传输对象 (API 输入/输出结构)
├── log/            # 日志文件存放目录
├── middlewares/    # 中间件 (认证, 日志)
├── models/         # 数据库模型 (GORM)
├── repositories/   # 数据仓库 (数据库操作)
├── response/       # 统一的 API 响应封装
├── routes/         # 路由定义
├── services/       # 业务逻辑服务
├── utils/          # 通用工具 (数据库, JWT, slog 日志等)
├── .env.example    # 环境变量配置示例
├── go.mod
├── go.sum
├── main.go         # 程序主入口
└── README.md
```
