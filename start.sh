#!/bin/bash

# 启动后端
echo "启动后端服务..."
cd backend
./../dist/moments &
BACKEND_PID=$!

# 等待后端启动
sleep 3

# 启动前端
echo "启动前端服务..."
cd ../front
pnpm run dev &
FRONTEND_PID=$!

echo "应用已启动！"
echo "后端 PID: $BACKEND_PID"
echo "前端 PID: $FRONTEND_PID"
echo "访问地址: http://localhost:3000"

# 等待用户中断
trap "echo '正在停止服务...'; kill $BACKEND_PID $FRONTEND_PID; exit" INT
wait