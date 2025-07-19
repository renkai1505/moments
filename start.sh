#!/bin/bash

# Moments 一键启动脚本
# 支持开发模式和Docker模式

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查命令是否存在
check_command() {
    if ! command -v $1 &> /dev/null; then
        print_error "$1 未安装，请先安装 $1"
        exit 1
    fi
}

# 检查端口是否被占用
check_port() {
    if lsof -Pi :$1 -sTCP:LISTEN -t >/dev/null ; then
        print_warning "端口 $1 已被占用"
        return 1
    fi
    return 0
}

# 显示帮助信息
show_help() {
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  dev         启动开发模式 (前端 + 后端)"
    echo "  docker      启动Docker模式"
    echo "  frontend    仅启动前端开发服务器"
    echo "  backend     仅启动后端开发服务器"
    echo "  stop        停止所有服务"
    echo "  clean       清理构建文件"
    echo "  help        显示此帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 dev       # 启动开发模式"
    echo "  $0 docker    # 启动Docker模式"
    echo "  $0 stop      # 停止所有服务"
}

# 启动开发模式
start_dev() {
    print_info "启动开发模式..."
    
    # 检查依赖
    check_command "node"
    check_command "pnpm"
    check_command "go"
    
    # 检查端口
    check_port 3000 || print_warning "前端端口 3000 可能被占用"
    check_port 37892 || print_warning "后端端口 37892 可能被占用"
    
    # 安装依赖
    print_info "安装前端依赖..."
    cd front
    pnpm install
    cd ..
    
    print_info "安装后端依赖..."
    cd backend
    go mod download
    cd ..
    
    # 启动后端
    print_info "启动后端服务器 (端口: 37892)..."
    cd backend
    nohup go run . > ../backend.log 2>&1 &
    BACKEND_PID=$!
    echo $BACKEND_PID > ../backend.pid
    cd ..
    
    # 等待后端启动
    print_info "等待后端启动..."
    sleep 3
    
    # 启动前端
    print_info "启动前端开发服务器 (端口: 3000)..."
    cd front
    nohup pnpm run dev > ../frontend.log 2>&1 &
    FRONTEND_PID=$!
    echo $FRONTEND_PID > ../frontend.pid
    cd ..
    
    print_success "开发模式启动完成!"
    print_info "前端地址: http://localhost:3000"
    print_info "后端地址: http://localhost:37892"
    print_info "API文档: http://localhost:37892/swagger/"
    print_info "日志文件: frontend.log, backend.log"
    print_info "使用 '$0 stop' 停止服务"
}

# 启动Docker模式
start_docker() {
    print_info "启动Docker模式..."
    
    check_command "docker"
    
    # 检查端口
    check_port 3000 || print_warning "端口 3000 可能被占用"
    
    # 创建数据目录
    mkdir -p moments
    
    # 停止并删除现有容器
    docker stop moments 2>/dev/null || true
    docker rm moments 2>/dev/null || true
    
    # 启动容器
    print_info "启动Docker容器..."
    docker run --name moments \
        -e JWT_KEY=cfqYVP6CZm9mSqLVGlmL \
        -e PORT=3000 \
        -d \
        -v $(pwd)/moments:/app/data \
        -p 3000:3000 \
        kingwrcy/moments:latest
    
    print_success "Docker模式启动完成!"
    print_info "应用地址: http://localhost:3000"
    print_info "数据目录: $(pwd)/moments"
    print_info "使用 '$0 stop' 停止服务"
}

# 仅启动前端
start_frontend() {
    print_info "启动前端开发服务器..."
    
    check_command "node"
    check_command "pnpm"
    check_port 3000 || print_warning "端口 3000 可能被占用"
    
    cd front
    pnpm install
    pnpm run dev
}

# 仅启动后端
start_backend() {
    print_info "启动后端开发服务器..."
    
    check_command "go"
    check_port 37892 || print_warning "端口 37892 可能被占用"
    
    cd backend
    go mod download
    go run .
}

# 停止所有服务
stop_services() {
    print_info "停止所有服务..."
    
    # 停止开发模式进程
    if [ -f backend.pid ]; then
        BACKEND_PID=$(cat backend.pid)
        if kill -0 $BACKEND_PID 2>/dev/null; then
            kill $BACKEND_PID
            print_info "已停止后端进程 (PID: $BACKEND_PID)"
        fi
        rm -f backend.pid
    fi
    
    if [ -f frontend.pid ]; then
        FRONTEND_PID=$(cat frontend.pid)
        if kill -0 $FRONTEND_PID 2>/dev/null; then
            kill $FRONTEND_PID
            print_info "已停止前端进程 (PID: $FRONTEND_PID)"
        fi
        rm -f frontend.pid
    fi
    
    # 停止Docker容器
    if docker ps -q -f name=moments | grep -q .; then
        docker stop moments
        docker rm moments
        print_info "已停止Docker容器"
    fi
    
    print_success "所有服务已停止"
}

# 清理构建文件
clean_build() {
    print_info "清理构建文件..."
    
    # 清理前端构建文件
    cd front
    rm -rf .output dist node_modules
    cd ..
    
    # 清理后端构建文件
    cd backend
    rm -rf dist public
    go clean
    cd ..
    
    # 删除日志和PID文件
    rm -f frontend.log backend.log frontend.pid backend.pid
    
    print_success "构建文件已清理"
}

# 主函数
main() {
    case "${1:-dev}" in
        "dev")
            start_dev
            ;;
        "docker")
            start_docker
            ;;
        "frontend")
            start_frontend
            ;;
        "backend")
            start_backend
            ;;
        "stop")
            stop_services
            ;;
        "clean")
            clean_build
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        *)
            print_error "未知选项: $1"
            show_help
            exit 1
            ;;
    esac
}

# 捕获中断信号
trap 'print_info "正在停止服务..."; stop_services; exit 0' INT TERM

# 运行主函数
main "$@"