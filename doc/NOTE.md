## web项目Demo
- 前后端+数据库+事务
  - https://blog.csdn.net/wsjzzcbq/article/details/127268924
  - https://gitee.com/wsjzzcbq/go-web

## 项目结构
### 参考文献
- web应用项目架构参考：https://blog.csdn.net/ygq13572549874/article/details/145559849
```text
app/
│── cmd/               # 入口文件
│   ├── main.go        # 主程序入口
│
│── internal/          # 内部应用逻辑（不对外暴露）
│   ├── app/           # 业务应用
│   │   ├── handlers/  # HTTP 处理函数
│   │   ├── services/  # 业务逻辑层
│   │   ├── repositories/ # 数据访问层
│   │   ├── models/    # 数据模型
│   │   ├── middleware/ # 中间件
│   │   └── routes/    # 路由管理
│   │
│   ├── config/        # 配置相关
│   ├── database/      # 数据库初始化、迁移
│   ├── logger/        # 日志组件
│   ├── utils/         # 工具函数
│   └── pkg/           # 内部可复用模块
│
│── api/               # API 相关定义
│   ├── openapi/       # OpenAPI 规范（Swagger等）
│   ├── protobuf/      # gRPC Protobuf 定义
│   └── graphql/       # GraphQL Schema 定义
│
│── migrations/        # 数据库迁移文件
│── scripts/           # 启动、构建、部署脚本
│── test/              # 测试代码（集成测试等）
│── web/               # 前端代码（如果有）
│── docs/              # 项目文档
│── .env               # 环境变量文件
│── go.mod             # Go 依赖管理文件
│── go.sum             # 依赖校验文件
│── Makefile           # 常用命令封装
│── README.md          # 说明文档
```
- 通用项目架构参考：https://juejin.cn/post/7209553403810791482
```text
project
├── CHANGELOG
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── README.md
├── api
│   └── openapi
│       └── openapi.yaml
├── assets
├── build
├── cmd
│   ├── ctl
│   │   └── main.go
│   ├── server
│   │   └── main.go
│   └── task
│       └── main.go
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── init
├── internal
│   ├── app
│   │   ├── ctl
│   │   ├── server
│   │   └── task
│   └── pkg
├── pkg
├── scripts
├── test
├── third_party
├── tools
└── web
```
- 另一个web项目架构
```text
├── cmd
│   ├── server
│   │   ├── main.go
│   │   ├── server_test.go
│   └── client
│       ├── main.go
│       └── client_test.go
├── pkg
│   ├── http
│   │   ├── server.go
│   │   ├── server_test.go
│   │   ├── client.go
│   │   └── client_test.go
│   ├── database
│   │   ├── connection.go
│   │   └── connection_test.go
│   └── utils
│       ├── converter.go
│       └── converter_test.go
└── internal
    ├── config
    │   ├── config.go
    │   └── config_test.go
    └── model
        ├── user.go
        └── user_test.go
```
- go官方文档：https://go.dev/doc/modules/layout#server-project

## 变量命名
### 参考文献
- https://zhuanlan.zhihu.com/p/216001587