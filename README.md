# Go Backend Starter

一个使用 Go 标准库 `net/http` 构建的后端入门项目，包含清晰的目录结构、路由、中间件、配置读取、模型和简单 REST API 示例。

## 项目结构

```text
.
├── main.go
├── go.mod
├── config/
│   └── config.go
├── handlers/
│   ├── ping.go
│   ├── response.go
│   └── users.go
├── models/
│   └── user.go
├── routes/
│   ├── middleware.go
│   └── routes.go
└── README.md
```

## 运行项目

默认端口是 `8080`：

```bash
go run .
```

也可以通过环境变量指定端口：

```bash
PORT=3000 go run .
```

Windows PowerShell：

```powershell
$env:PORT = "3000"
go run .
```

## API 示例

### 健康检查

```bash
curl http://localhost:8080/ping
```

响应：

```json
{"message":"pong"}
```

### 获取用户列表

```bash
curl http://localhost:8080/users
```

### 添加用户

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Charlie","email":"charlie@example.com"}'
```

Windows PowerShell：

```powershell
Invoke-RestMethod -Method Post `
  -Uri http://localhost:8080/users `
  -ContentType "application/json" `
  -Body '{"name":"Charlie","email":"charlie@example.com"}'
```

## 后续扩展建议

- 将 `handlers.UserHandler` 中的内存存储替换为 repository/service 层。
- 在 `config` 包中增加数据库连接字符串、日志级别等配置。
- 为 handlers 添加单元测试。
- 增加 graceful shutdown、请求 ID、CORS 或鉴权中间件。
