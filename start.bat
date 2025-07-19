@echo off
setlocal enabledelayedexpansion

REM Moments 一键启动脚本 (Windows版本)
REM 支持开发模式和Docker模式

set "MODE=%~1"
if "%MODE%"=="" set "MODE=dev"

echo [INFO] 启动模式: %MODE%

if "%MODE%"=="dev" goto :dev
if "%MODE%"=="docker" goto :docker
if "%MODE%"=="frontend" goto :frontend
if "%MODE%"=="backend" goto :backend
if "%MODE%"=="stop" goto :stop
if "%MODE%"=="clean" goto :clean
if "%MODE%"=="help" goto :help
goto :help

:dev
echo [INFO] 启动开发模式...
call :check_command node
call :check_command pnpm
call :check_command go

echo [INFO] 安装前端依赖...
cd front
call pnpm install
cd ..

echo [INFO] 安装后端依赖...
cd backend
go mod download
cd ..

echo [INFO] 启动后端服务器 (端口: 37892)...
cd backend
start /B go run . > ..\backend.log 2>&1
cd ..

echo [INFO] 等待后端启动...
timeout /t 3 /nobreak >nul

echo [INFO] 启动前端开发服务器 (端口: 3000)...
cd front
start /B pnpm run dev > ..\frontend.log 2>&1
cd ..

echo [SUCCESS] 开发模式启动完成!
echo [INFO] 前端地址: http://localhost:3000
echo [INFO] 后端地址: http://localhost:37892
echo [INFO] API文档: http://localhost:37892/swagger/
echo [INFO] 日志文件: frontend.log, backend.log
echo [INFO] 使用 'start.bat stop' 停止服务
goto :end

:docker
echo [INFO] 启动Docker模式...
call :check_command docker

echo [INFO] 创建数据目录...
if not exist moments mkdir moments

echo [INFO] 停止并删除现有容器...
docker stop moments 2>nul
docker rm moments 2>nul

echo [INFO] 启动Docker容器...
docker run --name moments -e JWT_KEY=cfqYVP6CZm9mSqLVGlmL -e PORT=37892 -d -v %cd%\moments:/app/data -p 37892:37892 kingwrcy/moments:latest

echo [SUCCESS] Docker模式启动完成!
echo [INFO] 应用地址: http://localhost:37892
echo [INFO] 数据目录: %cd%\moments
echo [INFO] 使用 'start.bat stop' 停止服务
goto :end

:frontend
echo [INFO] 启动前端开发服务器...
call :check_command node
call :check_command pnpm

cd front
call pnpm install
call pnpm run dev
cd ..
goto :end

:backend
echo [INFO] 启动后端开发服务器...
call :check_command go

cd backend
go mod download
go run .
cd ..
goto :end

:stop
echo [INFO] 停止所有服务...

REM 停止Docker容器
docker stop moments 2>nul
docker rm moments 2>nul
echo [INFO] 已停止Docker容器

echo [SUCCESS] 所有服务已停止
goto :end

:clean
echo [INFO] 清理构建文件...

echo [INFO] 清理前端构建文件...
cd front
if exist .output rmdir /s /q .output
if exist dist rmdir /s /q dist
if exist node_modules rmdir /s /q node_modules
cd ..

echo [INFO] 清理后端构建文件...
cd backend
if exist dist rmdir /s /q dist
if exist public rmdir /s /q public
go clean
cd ..

echo [INFO] 删除日志文件...
if exist frontend.log del frontend.log
if exist backend.log del backend.log

echo [SUCCESS] 构建文件已清理
goto :end

:help
echo 用法: start.bat [选项]
echo.
echo 选项:
echo   dev         启动开发模式 (前端 + 后端)
echo   docker      启动Docker模式
echo   frontend    仅启动前端开发服务器
echo   backend     仅启动后端开发服务器
echo   stop        停止所有服务
echo   clean       清理构建文件
echo   help        显示此帮助信息
echo.
echo 示例:
echo   start.bat dev       # 启动开发模式
echo   start.bat docker    # 启动Docker模式
echo   start.bat stop      # 停止所有服务
goto :end

:check_command
where %1 >nul 2>&1
if errorlevel 1 (
    echo [ERROR] %1 未安装，请先安装 %1
    exit /b 1
)
exit /b 0

:end