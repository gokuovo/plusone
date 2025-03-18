# OnePlusOne Gin项目

一个基于Gin框架的RESTful API项目，具有清晰的层级架构，集成了JWT鉴权和日志中间件，支持MySQL和SQLite数据库。

## 项目特点

- 基于Gin框架的RESTful API
- 清晰的分层架构（控制器、服务、仓库模式）
- 使用GORM进行数据库操作
- 同时支持MySQL和SQLite数据库，可自由切换
- JWT鉴权中间件
- 自定义日志中间件
- 完整的用户认证功能（注册、登录、获取用户信息）

## 项目结构

```
├── config/             # 配置相关
├── controllers/        # 控制器层，处理HTTP请求
├── middlewares/        # 中间件（鉴权、日志）
├── models/             # 数据模型
├── repositories/       # 数据访问层
├── services/           # 业务服务层
├── utils/              # 工具函数
├── routes/             # 路由管理
├── .env                # 环境配置
├── main.go             # 主程序入口
└── README.md           # 项目文档
```

## 快速开始

### 环境要求

- Go 1.16或更高版本

### 安装与运行

1. 克隆仓库
```bash
git clone https://github.com/yourusername/oneplusone.git
cd oneplusone
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置环境变量
编辑`.env`文件，设置数据库连接等参数

4. 运行应用
```bash
go run main.go
```

服务器将在`http://localhost:8080`上启动

## API接口

### 用户相关接口

- **注册用户**
  - `POST /api/register`
  - 请求体: `{"username": "test", "password": "password", "email": "test@example.com", "nickname": "测试用户"}`

- **用户登录**
  - `POST /api/login`
  - 请求体: `{"username": "test", "password": "password"}`
  - 返回JWT令牌

- **获取用户信息**
  - `GET /api/user/info`
  - 请求头: `Authorization: Bearer [token]`

## 数据库切换

在`.env`文件中修改`DB_TYPE`和`DB_SOURCE`参数以切换数据库:

```
# SQLite配置
DB_TYPE=sqlite
DB_SOURCE=oneplusone.db

# MySQL配置
DB_TYPE=mysql
DB_SOURCE=user:password@tcp(localhost:3306)/oneplusone?charset=utf8mb4&parseTime=True&loc=Local
```

## 开发和贡献

欢迎贡献代码或提出问题！
