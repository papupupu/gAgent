# gAgent

一个基于Go语言开发的Web应用。

## 项目结构

```
.
├── api/            # API接口定义
├── cmd/            # 主程序入口
├── configs/        # 配置文件
├── internal/       # 内部包
│   ├── handlers/   # HTTP处理器
│   ├── middleware/ # 中间件
│   ├── models/     # 数据模型
│   ├── repository/ # 数据访问层
│   ├── server/     # 服务器配置
│   └── service/    # 业务逻辑层
├── pkg/           # 可重用的包
└── app.yaml       # 应用配置文件
```

## 开发环境

- Go 1.22+

## 运行

```bash
go run cmd/server/main.go
```