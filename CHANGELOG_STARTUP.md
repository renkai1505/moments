# 启动脚本更新日志

## 🐛 修复的问题

### 1. 端口配置不一致
**问题**: Docker配置使用端口3000，但后端默认端口是37892，导致前端无法正确代理到后端。

**修复**:
- ✅ 更新 `docker-compose.yml` 使用端口37892
- ✅ 更新 `docker-start.sh` 使用端口37892  
- ✅ 更新 `docker-update.sh` 使用端口37892
- ✅ 更新 README.md 中的端口说明

### 2. 缺少便捷的启动脚本
**问题**: 用户需要手动启动前端和后端，操作繁琐。

**修复**:
- ✅ 创建 `start.sh` (Linux/macOS)
- ✅ 创建 `start.bat` (Windows)
- ✅ 支持开发模式和Docker模式
- ✅ 支持单独启动前端或后端
- ✅ 支持停止服务和清理构建文件

## 🚀 新增功能

### 1. 一键启动脚本
- **开发模式**: 同时启动前端和后端开发服务器
- **Docker模式**: 使用Docker容器运行应用
- **单独启动**: 可选择只启动前端或后端
- **服务管理**: 支持停止服务和清理构建文件

### 2. 智能检测
- 自动检查依赖是否安装 (Node.js, pnpm, Go, Docker)
- 自动检查端口是否被占用
- 自动安装项目依赖
- 自动创建必要的数据目录

### 3. 日志管理
- 开发模式自动生成日志文件
- 支持查看前端和后端日志
- 进程管理，支持优雅停止

### 4. 跨平台支持
- Linux/macOS: `start.sh`
- Windows: `start.bat`
- 统一的命令接口

## 📁 新增文件

```
.
├── start.sh              # Linux/macOS 启动脚本
├── start.bat             # Windows 启动脚本  
├── STARTUP_GUIDE.md      # 详细使用指南
├── .env.example          # 环境配置模板
└── CHANGELOG_STARTUP.md  # 本更新日志
```

## 🔧 修改的文件

### 配置文件
- `docker-compose.yml` - 修复端口配置
- `docker-start.sh` - 修复端口配置
- `docker-update.sh` - 修复端口配置
- `README.md` - 更新端口说明和添加启动指南

## 📋 使用说明

### 快速开始
```bash
# 开发模式
./start.sh dev

# Docker模式  
./start.sh docker

# 查看帮助
./start.sh help
```

### 访问地址
- **开发模式**: 前端 http://localhost:3000, 后端 http://localhost:37892
- **Docker模式**: http://localhost:37892

## 🎯 改进效果

1. **简化部署**: 一键启动，无需手动配置
2. **修复端口问题**: 前后端端口配置统一
3. **提升开发体验**: 自动依赖检查和安装
4. **增强可维护性**: 统一的启动和管理方式
5. **跨平台支持**: 支持Linux、macOS和Windows

## 🔄 向后兼容

- 原有的 `make` 命令仍然可用
- 原有的Docker命令仍然可用
- 新增的启动脚本是对现有功能的增强，不会影响现有使用方式