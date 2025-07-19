# Moments 一键启动指南

本项目提供了多种启动方式，支持开发模式和Docker模式。

## 🚀 快速开始

### 开发模式 (推荐)
```bash
# Linux/macOS
./start.sh dev

# Windows
start.bat dev
```

### Docker模式
```bash
# Linux/macOS
./start.sh docker

# Windows
start.bat docker
```

## 📋 启动选项

| 选项 | 说明 | 适用场景 |
|------|------|----------|
| `dev` | 启动开发模式 (前端 + 后端) | 开发调试 |
| `docker` | 启动Docker模式 | 生产部署 |
| `frontend` | 仅启动前端开发服务器 | 前端开发 |
| `backend` | 仅启动后端开发服务器 | 后端开发 |
| `stop` | 停止所有服务 | 停止服务 |
| `clean` | 清理构建文件 | 清理缓存 |
| `help` | 显示帮助信息 | 查看帮助 |

## 🔧 环境要求

### 开发模式
- **Node.js** (v16+)
- **pnpm** (v8+)
- **Go** (v1.23+)

### Docker模式
- **Docker** (v20+)

## 📁 项目结构

```
.
├── front/                 # 前端项目 (Nuxt.js)
├── backend/              # 后端项目 (Go)
├── start.sh              # Linux/macOS 启动脚本
├── start.bat             # Windows 启动脚本
├── docker-compose.yml    # Docker Compose 配置
├── docker-start.sh       # Docker 启动脚本
└── moments/              # 数据目录 (自动创建)
```

## 🌐 访问地址

### 开发模式
- **前端**: http://localhost:3000
- **后端**: http://localhost:37892
- **API文档**: http://localhost:37892/swagger/

### Docker模式
- **应用**: http://localhost:37892
- **API文档**: http://localhost:37892/swagger/

## 🔍 故障排除

### 端口冲突
如果遇到端口被占用的问题：

```bash
# 查看端口占用
lsof -i :3000    # 前端端口
lsof -i :37892   # 后端端口

# 停止占用端口的进程
kill -9 <PID>
```

### 依赖问题
确保安装了所需的依赖：

```bash
# 检查 Node.js
node --version

# 检查 pnpm
pnpm --version

# 检查 Go
go version

# 检查 Docker
docker --version
```

### 权限问题
```bash
# 给脚本执行权限
chmod +x start.sh
```

## 📝 日志文件

启动后会在项目根目录生成以下日志文件：
- `frontend.log` - 前端日志
- `backend.log` - 后端日志

## 🛠️ 开发指南

### 前端开发
```bash
# 进入前端目录
cd front

# 安装依赖
pnpm install

# 启动开发服务器
pnpm run dev
```

### 后端开发
```bash
# 进入后端目录
cd backend

# 安装依赖
go mod download

# 启动开发服务器
go run .
```

### 构建项目
```bash
# 使用 Makefile
make build

# 或手动构建
make frontend
make backend
```

## 🔧 配置说明

### 环境变量
- `PORT` - 后端端口 (默认: 37892)
- `JWT_KEY` - JWT密钥
- `DB` - 数据库配置
- `CORS_ORIGIN` - CORS配置
- `UPLOAD_DIR` - 上传目录
- `LOG_LEVEL` - 日志级别

### Docker配置
Docker模式使用以下配置：
- 端口映射: `37892:37892`
- 数据卷: `./moments:/app/data`
- 环境变量: `JWT_KEY`, `PORT`

## 🚨 注意事项

1. **端口配置**: 确保前端代理配置与后端端口一致
2. **数据持久化**: Docker模式数据保存在 `./moments` 目录
3. **日志查看**: 开发模式日志保存在根目录的 `.log` 文件中
4. **进程管理**: 使用 `stop` 命令正确停止服务

## 📞 支持

如果遇到问题，请检查：
1. 依赖是否正确安装
2. 端口是否被占用
3. 日志文件中的错误信息
4. 网络连接是否正常

## 🔄 更新

### 更新Docker镜像
```bash
./docker-update.sh
```

### 清理并重新构建
```bash
# 清理构建文件
./start.sh clean

# 重新启动
./start.sh dev
```